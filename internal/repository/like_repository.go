package repository

import (
	"context"

	"github.com/PNYwise/like-service/internal/domain"
	"github.com/jackc/pgx/v5"
)

type likeRepository struct {
	db *pgx.Conn
}

func NewLikeRepository(db *pgx.Conn) domain.ILikeRepository {
	return &likeRepository{
		db: db,
	}
}

// GetByPostUuid implements domain.ILikeRepository.
func (*likeRepository) GetByPostUuid(ctx context.Context, uuid string) (*[]domain.Like, error) {
	panic("unimplemented")
}

// Set implements domain.ILikeRepository.
func (*likeRepository) Set(ctx context.Context, like *domain.Like) error {
	panic("unimplemented")
}

// Unset implements domain.ILikeRepository.
func (*likeRepository) Unset(ctx context.Context, userUuid string, PostUuid string) error {
	panic("unimplemented")
}
