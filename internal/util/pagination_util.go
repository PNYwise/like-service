package util

import (
	"math"

	"github.com/PNYwise/like-service/internal/domain"
)

func GeneratePagination(paginationRequest *domain.Pagination, count uint64) *domain.Pagination {
	totalPages := uint64(math.Ceil(float64(count) / float64(paginationRequest.Take)))
	pagination := new(domain.Pagination)
	pagination.Page = paginationRequest.Page
	pagination.ItemCount = count
	pagination.PageCount = totalPages

	return pagination
}
