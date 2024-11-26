package utils

import (
	"math"

	"github.com/n0o01lh/llp/internals/core/domain"
	"gorm.io/gorm"
)

func Paginate(value interface{}, whereClause string, pagination *domain.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Where(whereClause).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func CalculateTotalPagesWhenUseWhereClause(totalRows int64, pagination *domain.Pagination) *domain.Pagination {
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.TotalPages = totalPages

	return pagination
}
