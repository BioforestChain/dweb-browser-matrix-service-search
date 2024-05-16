package userConstant

const (
	Address2UserRedisKey = "address2user:%s"
	Token2UserRedisKey   = "token2user:%s"
	UserTokenTTL         = 300 //token过期时间 5min
	UserInfoTTL          = 300 // 过期时间 5min
)
