package domainModel

import (
	"context"
	"fmt"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/entity/db/domainDbEntity"
	myError "github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/error"
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/model/common"
	"gorm.io/gorm"
)

type model struct {
	Ctx context.Context
	DB  *gorm.DB
}

func NewModel(ctx context.Context) *model {
	return &model{ctx, common.ConnectionObject(ctx).DB}
}
func (m *model) TableName() string {
	return "service"
}

func (m *model) GetDomainList() (list []*domainDbEntity.Service, total int64, err myError.Error) {
	domainModule := domainDbEntity.Service{}
	db := m.DB.Model(&domainModule)

	if err := db.Count(&total).Error; err != nil {
		return list, total, common.CheckMysqlError(err)
	}
	fmt.Println("total", total)
	if err := db.Order("id desc").Find(&list).Error; err != nil {
		return list, total, common.CheckMysqlError(err)
	}
	fmt.Println("list", list)

	return list, total, nil
}

func (m *model) GetServicePublicKey(domain string) (domainDbEntity.Service, myError.Error) {
	domainModule := domainDbEntity.Service{}
	db := m.DB.Model(&domainModule)

	err := db.Select("public_key").
		Where("`domain` = ? ", domain).
		First(&domainModule).Error
	if err != nil {
		return domainModule, common.CheckMysqlError(err)
	}

	return domainModule, nil
}
