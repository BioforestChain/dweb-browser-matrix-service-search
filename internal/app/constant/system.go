package constant

import time2 "time"

const RedisPrefix = "matrix:"

const (
	ApiFileName = "curl"
	ApiKeyName  = "matrixClientUserSearch"
	ApiLimit    = 50
	CurlTimeOut = 1 * time2.Millisecond
	TimeAfter   = 5000 * time2.Millisecond
)
