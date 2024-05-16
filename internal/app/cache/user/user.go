package user

import (
	"context"
	"fmt"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/constant"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/constant/userConstant"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/pkg/support-go/helper/redis"
	"github.com/gin-gonic/gin"
)

type Cache struct {
	Ctx      context.Context
	GCtx     *gin.Context
	RedisCli *redis.RedisInstance
}

func NewCache(ctx context.Context) *Cache {
	redisCli, _ := redis.GetRedisInstance("default")
	return &Cache{Ctx: ctx, RedisCli: redisCli}
}

// GetCacheByAddress 获取缓存
func (c *Cache) GetCacheByAddress(address string) string {
	key := c.getUserCacheKey(address)
	value, _ := c.RedisCli.GetString(c.Ctx, key)
	return value
}

// SetCacheByAddress 设置缓存
func (c *Cache) SetCacheByAddress(address string, userProfile string) {
	key := c.getUserCacheKey(address)
	c.RedisCli.Set(c.Ctx, key, userProfile, userConstant.UserInfoTTL)
}

// getUserCacheKey 缓存的key
func (c *Cache) getUserCacheKey(address string) string {
	return constant.RedisPrefix + fmt.Sprintf(userConstant.Address2UserRedisKey, address)
}
