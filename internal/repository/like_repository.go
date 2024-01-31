package repository

import (
	"context"
	"log"
	"time"

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
func (*likeRepository) GetByPostUuid(ctx context.Context, uuid string, page uint64) (*[]domain.Like, uint64, error) {
	panic("unimplemented")
}

// Set implements domain.ILikeRepository.
func (l *likeRepository) Set(ctx context.Context, like *domain.Like) error {
	query :=
		`INSERT INTO likes (user_uuid, post_uuid, created_at)
		VALUES ($1, $2, $3)
		RETURNING uuid, user_uuid, post_uuid, created_at`
	err := l.db.QueryRow(
		ctx,
		query,
		like.UserUuid, like.PostUuid, time.Now(),
	).Scan(&like.Uuid, &like.UserUuid, &like.PostUuid, &like.CreatedAt)
	if err != nil {
		log.Fatalf("err: %v", err)
		return err
	}
	return nil
}

// Unset implements domain.ILikeRepository.
func (*likeRepository) Unset(ctx context.Context, userUuid string, postUuid string) error {
	panic("unimplemented")
}
