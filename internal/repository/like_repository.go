package repository

import (
	"context"
	"fmt"
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
func (l *likeRepository) GetByPostUuid(ctx context.Context, postUuid string, page uint64) (*[]domain.Like, uint64, error) {
	query := "SELECT uuid, user_uuid, post_uuid, created_at, deleted_at FROM likes WHERE post_uuid = $1 OFFSET $2 LIMIT $3"
	limit := uint64(15)
	offset := (page - 1) * limit
	row, err := l.db.Query(ctx, query, postUuid, offset, limit)
	if err != nil {
		fmt.Printf("Error executing query: %v", err)
		return nil, 0, err
	}
	var likes []domain.Like
	for row.Next() {
		var like domain.Like
		if err := row.Scan(&like.Uuid, &like.UserUuid, &like.PostUuid, &like.CreatedAt, &like.DeletedAt); err != nil {
			fmt.Printf("Error Scaning query: %v", err)
			return nil, 0, err
		}
		likes = append(likes, like)
	}
	return &likes, limit, nil
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
		fmt.Printf("Error executing query: %v", err)
		return err
	}
	return nil
}

// Unset implements domain.ILikeRepository.
func (l *likeRepository) Unset(ctx context.Context, userUuid string, postUuid string) error {

	query := "DELETE FROM likes WHERE user_uuid = $1 AND post_uuid = $2"
	if _, err := l.db.Exec(ctx, query, userUuid, postUuid); err != nil {
		fmt.Printf("Error executing query: %v", err)
		return err
	}
	return nil
}

// Exist implements domain.ILikeRepository.
func (l *likeRepository) Exist(ctx context.Context, userUuid string, postUuid string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM likes WHERE user_uuid = $1 AND post_uuid = $2)"
	var exist bool
	row, err := l.db.Query(ctx, query, userUuid, postUuid)
	if err != nil {
		fmt.Printf("Error executing query: %v", err)
		return false, err
	}
	for row.Next() {
		if err := row.Scan(&exist); err != nil {
			fmt.Printf("Error Scaning query: %v", err)
			return false, err
		}
	}
	return exist, nil

}
