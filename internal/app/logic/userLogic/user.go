package userLogic

import (
	"context"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	userCache "github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/cache/user"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/constant"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/entity/db/userDbEntity"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/entity/req/userReqEntity"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/entity/resp/userRespEntity"
	myError "github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/error"
	errMsg "github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/error/userError"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/service/domainService"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/helper/config"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/pkg/support-go/helper/curl"
	"github.com/gin-gonic/gin"
	jsonIter "github.com/json-iterator/go"
	"io"
	"log"
	time2 "time"
)

var userErr myError.Error

type logic struct {
	Ctx  context.Context
	GCtx *gin.Context
}

func NewLogic(ctx *gin.Context) *logic {
	return &logic{Ctx: ctx, GCtx: ctx}
}

// getUserProfileCache
func (l *logic) getUserProfileCache(address string) (cachedProfile userDbEntity.UserProfile, err error) {
	cache := userCache.NewCache(l.Ctx)
	profile := cache.GetCacheByAddress(address)
	if len(profile) == 0 {
		return cachedProfile, nil
	}
	err = json.Unmarshal([]byte(profile), &cachedProfile)
	if err != nil {
		log.Println("err :", err)
	}
	return cachedProfile, nil
}

type PostBody struct {
	SearchTerm string `json:"search_term"`
	Limit      uint32 `json:"limit"`

	EncryptedData string `json:"encrypted_data"`
}
type UserSearchResPon struct {
	Limited bool `json:"limited"`
	Results []struct {
		UserId        string `json:"user_id"`
		DisplayName   string `json:"display_name"`
		AvatarUrl     string `json:"avatar_url"`
		WalletAddress string `json:"wallet_address"`
	} `json:"results"`
}

// 1. 查缓存
// 1.1 列出 domainList
// 2. 并发调api查 https://172.25.11.243/_matrix/client/v3/user_directory/search
// 3. 再写入缓存

func (l *logic) encryptData(key []byte, plaintext string) (string, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	b := base64.StdEncoding.EncodeToString([]byte(plaintext))
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (l *logic) encryptKey(key []byte) string {
	// Here, you can implement RSA encryption of the symmetric key
	// with Python server's public key
	// For simplicity, let's just base64 encode the key
	return base64.StdEncoding.EncodeToString(key)
}

// 签名
func RsaSignWithSha256(data []byte, keyBytes []byte) []byte {
	// 读取私钥
	privateKey, err := ReadRSAPrivateKey(keyBytes)
	if err != nil {
		return nil
	}

	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		panic(err)
	}

	return signature
}

// 获取私钥
func ReadRSAPrivateKey(readFile []byte) (*rsa.PrivateKey, error) {

	var err error
	// 使用pem解码
	pemBlock, _ := pem.Decode(readFile)

	if pemBlock == nil {
		panic(errors.New("private key error!"))
	}

	var pkixPrivateKey interface{}
	//解析PKCS1格式的私钥
	if pemBlock.Type == "RSA PRIVATE KEY" {
		// -----BEGIN RSA PUBLIC KEY-----
		fmt.Println("-------------------------------解析PKCS1格式的PrivateKey-----------------------------------------")
		pkixPrivateKey, err = x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
		//panic(errors.New("请 解析PKCS1格式的私钥 private key error!"))
		//return nil, err
	} else if pemBlock.Type == "PRIVATE KEY" {
		// -----BEGIN PUBLIC KEY-----
		fmt.Println("-------------------------------解析PKCS8格式的PrivateKey-----------------------------------------")
		pkixPrivateKey, err = x509.ParsePKCS8PrivateKey(pemBlock.Bytes)

	}
	if err != nil {
		panic(err)
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	privateKey := pkixPrivateKey.(*rsa.PrivateKey)
	return privateKey, nil
}

// 获取公钥
func ReadRSAPublicKey(readFile []byte) (*rsa.PublicKey, error) {
	var err error
	// 使用pem解码
	pemBlock, _ := pem.Decode(readFile)
	var pkixPublicKey interface{}
	if pemBlock.Type == "RSA PUBLIC KEY" {
		//fmt.Println("-------------------------------解析PKCS1格式的PublicKey-----------------------------------------")
		// -----BEGIN RSA PUBLIC KEY-----
		pkixPublicKey, err = x509.ParsePKCS1PublicKey(pemBlock.Bytes)
	} else if pemBlock.Type == "PUBLIC KEY" {
		// -----BEGIN PUBLIC KEY-----
		//fmt.Println("-------------------------------解析PKCS8格式的PublicKey-----------------------------------------")
		pkixPublicKey, err = x509.ParsePKIXPublicKey(pemBlock.Bytes)
	}
	if err != nil {
		return nil, err
	}
	publicKey := pkixPublicKey.(*rsa.PublicKey)
	return publicKey, nil
}

// 验证
func RsaVerySignWithSha256(data, signData []byte, pubKey *rsa.PublicKey) bool {
	hashed := sha256.Sum256(data)
	err := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], signData)
	if err != nil {
		panic(err)
	}
	return true
}

// 公钥加密
func RsaEncrypt(data []byte, pub *rsa.PublicKey) []byte {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if err != nil {
		panic(err)
	}
	return ciphertext
}

// 私钥解密
func RsaDecrypt(ciphertext []byte, prvKey *rsa.PrivateKey) []byte {
	data, err := rsa.DecryptPKCS1v15(rand.Reader, prvKey, ciphertext)
	if err != nil {
		panic(err)
	}
	return data
}

func (l *logic) GetUserInfo(req userReqEntity.Info) (res userRespEntity.Info, err myError.Error) {
	userErr = myError.NewUserError()
	profile, er := l.getUserProfileCache(req.SearchTerm)
	if er != nil {
		log.Println("GetUserInfo err :", err)
	}
	res = userRespEntity.Info{profile}
	if len(res.UserId) > 0 && len(res.WalletAddress) > 0 {
		return res, nil
	}
	domainList, _, _ := l.getDomainList()
	apiName := config.GetCfgValueFromYml(constant.ApiFileName, constant.ApiKeyName)
	// 创建一个 userRespEntity.Info 类型的 channel，用于接收每个 goroutine 的结果
	results := make(chan userRespEntity.Info, len(domainList))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var resPon UserSearchResPon
	for i := 0; i < len(domainList); i++ {
		var data = req.SearchTerm
		// 获取 私钥 公钥 源数据
		pubKeyStr, err := domainService.NewService(l.Ctx).GetServicePublicKey(domainList[i])
		if err != nil {
			fmt.Println("公钥获取失败：", err)
		}
		pubKeyRead := []byte(pubKeyStr)
		//读取公钥
		pubKey, er := ReadRSAPublicKey(pubKeyRead)
		if er != nil {
			fmt.Println("公钥读取失败：", er)
		}
		ciphertext := RsaEncrypt([]byte(data), pubKey)

		body := PostBody{
			req.SearchTerm, constant.ApiLimit, base64.StdEncoding.EncodeToString(ciphertext)}
		postBody, _ := json.Marshal(body)

		go func(i int) {
			select {
			case <-ctx.Done():
				return //  如果上下文已经取消，则该goroutine终止
			default:

				api := fmt.Sprintf("%s%s%s", "https://", domainList[i], apiName)
				req := curl.Request{
					Url:        api,
					ReqTimeOut: constant.CurlTimeOut,
					BodyData:   postBody,
					Headers:    map[string]string{"Content-Type": "application/json"},
				}
				start := time2.Now()
				resp, er := curl.Post(req)
				end := time2.Now()
				interval := end.Sub(start)
				//执行时间
				fmt.Println(domainList[i], "执行时间：", interval)
				if er != nil {
					userErr.SetCodeMsg(errMsg.Fail)
					return
				}
				//Success
				if len(resp.Body) > 0 {
					jsonIter.Unmarshal(resp.Body, &resPon)
					if len(resPon.Results) == 0 {
						return
					} else if resp.StatusCode == 401 {
						log.Println("resp :"+api, string(resp.Body))
					}

					res = userRespEntity.Info{
						userDbEntity.UserProfile{
							UserId:        resPon.Results[0].UserId,
							DisplayName:   resPon.Results[0].DisplayName,
							AvatarUrl:     resPon.Results[0].AvatarUrl,
							WalletAddress: resPon.Results[0].WalletAddress,
						},
					}
					results <- res // 将结果发送到 channel
					cancel()       // 调用 cancel() 终止所有 goroutine
					return
				}
			}
		}(i)
	}

	select {
	case tmp := <-results: // 阻塞等待任一响应
		if len(tmp.UserProfile.UserId) > 0 && len(tmp.UserProfile.WalletAddress) > 0 {
			fmt.Println("Success! Data received:", tmp)
			last, _ := json.Marshal(tmp)
			if err != nil {
				log.Println("err :", err)
			}
			//写进缓存
			userCache.NewCache(l.Ctx).SetCacheByAddress(res.WalletAddress, string(last))

			return tmp, nil
			// 处理成功结果...
		}

	case <-time2.After(constant.TimeAfter):
		fmt.Println("context cancelled")
		return
	}
	return
}

// getDomainList
//
//	@Description:  //domainList = append(domainList, "matrix.org", "a.dweba.com", "172.25.11.243")
//	@receiver l
//	@return list
//	@return total
//	@return err
func (l *logic) getDomainList() (list []string, total int64, err myError.Error) {
	domainList := make([]string, 0)
	domainData, err := domainService.NewService(l.Ctx).GetDomainList()
	if err != nil {
		return list, 0, nil
	}
	for _, Service := range domainData.List {
		if Service != nil { // Check if the pointer is not nil to avoid dereference of nil pointer
			domainList = append(domainList, Service.Domain)
		}
	}
	return domainList, total, nil
}
