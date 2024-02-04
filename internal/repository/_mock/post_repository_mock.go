package _mock

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockLikeRepository is a mock implementation of ILikeRepository
type PostRepositoryMock struct {
	mock.Mock
}

func (p *PostRepositoryMock) Exist(ctx context.Context, postUuid string) (bool, error) {
	args := p.Called(ctx, postUuid)
	return args.Get(0).(bool), args.Error(1)
}
