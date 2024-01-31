package domain

import (
	"context"
	"time"
)

type Like struct {
	Uuid      string
	UserUuid  string
	PostUuid  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
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
	GetByPostUuid(context.Context, string, uint64) (*[]Like, uint64, error)
	Set(context.Context, *Like) error
	Unset(context.Context, string, string) error
}

type ILikeService interface {
	GetByPostUuid(context.Context, string, uint64) (*[]Like, *Pagination, error)
	Set(context.Context, *SetLikeRequest) error
	Unset(context.Context, string, string) error
}
