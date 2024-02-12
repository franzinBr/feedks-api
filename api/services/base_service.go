package services

import (
	"github.com/franzinBr/feedks-api/data/db"
	"gorm.io/gorm"
)

type BaseService struct {
	Db *gorm.DB
}

func NewBaseService() *BaseService {
	return &BaseService{
		Db: db.GetDB(),
	}
}
