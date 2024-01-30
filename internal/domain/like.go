package domain

import "context"

type Like struct {
}

type CreateLikeRequest struct {
}

type ILikeRepository interface {
	GetByPostUuid(context.Context, string) (*[]Like, error)
	Set(context.Context, *Like) error
	Unset(context.Context, string, string) error
}

type ILikeService interface {
	GetByPostUuid(context.Context, string) (*[]Like, error)
	Set(context.Context, *CreateLikeRequest) error
	Unset(context.Context, string, string) error
}
