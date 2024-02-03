package _mock

import (
	"context"

	"github.com/PNYwise/like-service/internal/domain"
	"github.com/stretchr/testify/mock"
)

// MockLikeRepository is a mock implementation of ILikeRepository
type LikeRepositoryMock struct {
	mock.Mock
}

// GetByPostUuid mocks the GetByPostUuid method of ILikeRepository
func (m *LikeRepositoryMock) GetByPostUuid(ctx context.Context, postUuid string, page uint64) (*[]domain.Like, uint64, error) {
	args := m.Called(ctx, postUuid, page)
	return args.Get(0).(*[]domain.Like), args.Get(1).(uint64), args.Error(2)
}

// Create mocks the Set method of ILikeRepository
func (m *LikeRepositoryMock) Set(ctx context.Context, like *domain.Like) error {
	args := m.Called(ctx, like)
	return args.Error(0)
}

func (m *LikeRepositoryMock) Unset(ctx context.Context, userUuid string, postUuid string) error {
	args := m.Called(ctx, userUuid, postUuid)
	return args.Error(0)

}
func (m *LikeRepositoryMock) Exist(ctx context.Context, userUuid string, postUuid string) (bool, error) {
	args := m.Called(ctx, userUuid, postUuid)
	return args.Get(0).(bool), args.Error(1)
}
