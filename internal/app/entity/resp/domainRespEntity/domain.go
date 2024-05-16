package domainRespEntity

import (
	"github.com/BioforestChain/dweb-browser-matrix-service-search/internal/app/entity/db/domainDbEntity"
)

type OnlineList struct {
	List  []*domainDbEntity.Service `json:"list"`
	Total int64                     `json:"total"`
}

type ServiceInfo struct {
	Info domainDbEntity.Service `json:"info"`
}
