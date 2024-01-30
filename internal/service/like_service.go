package service

import (
	"context"

	"github.com/PNYwise/like-service/internal/domain"
)

func NewLikeService(likeRepo domain.ILikeRepository) domain.ILikeService {
	return &likeService{
		likeRepo: likeRepo,
	}
}

type likeService struct {
	likeRepo domain.ILikeRepository
}

// GetByPostUuid implements domain.ILikeService.
func (*likeService) GetByPostUuid(ctx context.Context, uuid string, page uint64) (*[]domain.Like, *domain.Pagination, error) {
	panic("unimplemented")
}

// Set implements domain.ILikeService.
func (*likeService) Set(context.Context, *domain.CreateLikeRequest) error {
	panic("unimplemented")
}

// Unset implements domain.ILikeService.
func (*likeService) Unset(context.Context, string, string) error {
	panic("unimplemented")
}
