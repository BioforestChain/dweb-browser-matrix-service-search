package domainService

import (
	"context"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/entity/resp/domainRespEntity"
	myError "github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/error"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/model/professional/domainModel"

	"github.com/gin-gonic/gin"
)

var domainListErr myError.Error

type service struct {
	Ctx    context.Context
	GCtx   *gin.Context
	userId uint32
}

func NewService(ctx context.Context) *service {
	return &service{Ctx: ctx}
}

// GetCookbookList 列表
func (l *service) GetDomainList() (resp domainRespEntity.OnlineList, err myError.Error) {
	list, total, err := domainModel.NewModel(l.Ctx).GetDomainList()
	if err != nil {
		return resp, err
	}
	resp.List = list
	resp.Total = total
	return resp, nil
}

// GetServicePublicKey
func (l *service) GetServicePublicKey(domain string) (resp string, err myError.Error) {
	info, err := domainModel.NewModel(l.Ctx).GetServicePublicKey(domain)
	if err != nil {
		return resp, err
	}
	return info.PublicKey, nil
}
