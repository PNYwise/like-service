package service

import (
	"context"
	"errors"

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
func (l *likeService) Set(ctx context.Context, request *domain.SetLikeRequest) error {
	// TODO: validation
	like := &domain.Like{
		UserUuid: request.UserUuid,
		PostUuid: request.PostUuid,
	}
	if err := l.likeRepo.Set(ctx, like); err != nil {
		return errors.New("Internal Server Error")
	}
	return nil
}

// Unset implements domain.ILikeService.
func (l *likeService) Unset(ctx context.Context, userUuid string, postUuid string) error {
	if exist, _ := l.likeRepo.Exist(ctx, userUuid, postUuid); !exist {
		return errors.New("Like not found")
	}
	if err := l.likeRepo.Unset(ctx, userUuid, postUuid); err != nil {
		return errors.New("Internal Server Error")
	}
	return nil
}
