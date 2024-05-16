package domainDbEntity

import (
	"gorm.io/gorm"
	"time"
)

type Service struct {
	Id        uint32         `json:"id"`     // User ID
	Domain    string         `json:"domain"` // domain
	Remark    string         `json:"remark"` // remark
	Ip        string         `json:"ip"`     // ip
	PublicKey string         `json:"public_key"`
	CreateAt  time.Time      `json:"create_at"` // Created Time
	UpdateAt  time.Time      `json:"update_at"` // Updated Time
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
