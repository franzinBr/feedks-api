package helpers

import (
	"math"

	"github.com/franzinBr/feedks-api/api/dtos"
	"gorm.io/gorm"
)

func Paginate[T any](value interface{}, paginationRequest *dtos.PaginationRequest, paginationResponse *dtos.PaginationResponse[T], db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var TotalItems int64
	db.Model(value).Count(&TotalItems)

	paginationResponse.Page = paginationRequest.GetPage()
	paginationResponse.Limit = paginationRequest.GetLimit()
	paginationResponse.TotalItems = TotalItems
	paginationResponse.TotalPages = int64(math.Ceil(float64(TotalItems) / float64(paginationRequest.GetLimit())))
	paginationResponse.HasNextPage = int64(paginationResponse.Page) < paginationResponse.TotalPages
	paginationResponse.HasPreviousPage = paginationResponse.Page > 1

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(paginationRequest.GetOffset()).Limit(paginationRequest.GetLimit()).Order("id asc")
	}
}
