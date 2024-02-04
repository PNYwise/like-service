package domain

import "context"

type IPostRepository interface {
	Exist(ctx context.Context, postUuid string) (bool, error)
}
