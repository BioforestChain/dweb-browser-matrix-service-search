package userReqEntity

import (
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/entity"
)

// List 列表
type List struct {
	entity.PaginationSearch
}

// Item 单个操作
type Item struct {
	Id uint32 `json:"news_page_id" form:"news_page_id"  binding:"required" msg:"必填"`
}

// Update 更新
type Update struct {
	Id uint32 `json:"news_page_id"  binding:"required" msg:"必填"`
}
