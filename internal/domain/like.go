package domain

import (
	"context"
	"database/sql"
	"time"
)

type Like struct {
	Uuid      string
	UserUuid  string
	PostUuid  string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

type SetLikeRequest struct {
	UserUuid string
	PostUuid string
}

type Pagination struct {
	Page      uint64
	Take      uint64
	ItemCount uint64
	PageCount uint64
}

func (p *Pagination) Skip() uint64 {
	return (p.Page - 1) * p.Take
}

type ILikeRepository interface {
	GetByPostUuid(ctx context.Context, postUuid string, page uint64) (*[]Like, uint64, error)
	Set(ctx context.Context, like *Like) error
	Unset(ctx context.Context, userUuid string, postUuid string) error
	Exist(ctx context.Context, userUuid string, postUuid string) (bool, error)
}

type ILikeService interface {
	GetByPostUuid(ctx context.Context, postUuid string, page uint64) (*[]Like, *Pagination, error)
	Set(ctx context.Context, request *SetLikeRequest) error
	Unset(ctx context.Context, userUuid string, postUuid string) error
}
