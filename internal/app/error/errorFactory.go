package error

import (
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/error/mysql"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/error/newsPageError"
)

// NewMysqlError 实例化mysql错误
func NewMysqlError() Error {
	return &MyError{
		msgList: mysql.ErrorMessageList(),
	}
}

// NewNewsPageError 实例化mysql错误
func NewNewsPageError() Error {
	return &MyError{
		msgList: newsPageError.ErrorMessageList(),
	}
}

func NewTypeError() Error {
	return &MyError{
		msgList: newsPageError.ErrorMessageList(),
	}
}
