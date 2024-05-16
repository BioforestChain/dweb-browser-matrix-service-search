package error

import (
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/error/mysql"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/error/newsPageError"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/error/userError"
)

// NewMysqlError 实例化mysql错误
func NewMysqlError() Error {
	return &MyError{
		msgList: mysql.ErrorMessageList(),
	}
}

// NewUserError 实例化mysql错误
func NewUserError() Error {
	return &MyError{
		msgList: userError.ErrorMessageList(),
	}
}

func NewTypeError() Error {
	return &MyError{
		msgList: newsPageError.ErrorMessageList(),
	}
}
