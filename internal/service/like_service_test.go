package service

import (
	"context"
	"testing"
	"time"

	"github.com/PNYwise/like-service/internal/domain"
	"github.com/PNYwise/like-service/internal/repository/_mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetByPostUuid(t *testing.T) {
	// Create a mock repository
	mockRepo := new(_mock.LikeRepositoryMock)
	mockPostRepo := new(_mock.PostRepositoryMock)

	// Create a like service with the mock repository
	likeService := NewLikeService(mockRepo, mockPostRepo)

	// Define a fake UUID
	fakeUuid := uuid.New().String()
	fakeUserUuid := uuid.New().String()
	fakePostUuid := uuid.New().String()

	ctx := context.Background()

	// Create an example list of like
	fakelikes := []domain.Like{
		{
			Uuid:      fakeUuid,
			UserUuid:  fakeUserUuid,
			PostUuid:  fakePostUuid,
			CreatedAt: time.Now(),
		},
	}

	// Create an example pagination
	fakePage := uint64(1)
	fakeLimit := uint64(15)
	fakeOffset := (fakePage - 1) * fakeLimit
	fakePagination := domain.Pagination{
		Take:      fakeOffset,
		ItemCount: uint64(0),
		PageCount: uint64(0),
	}

	// Expect the GetByPostUuid method to be called with the correct argument
	mockRepo.On("GetByPostUuid", ctx, fakePostUuid, fakePage).Return(&fakelikes, fakeOffset, nil)

	// Call the GetByPostUuid method of the post service
	resultLikes, resultPagination, err := likeService.GetByPostUuid(ctx, fakePostUuid, fakePage)

	// Assert that the mock repository's GetByPostUuid method was called with the correct argument
	mockRepo.AssertExpectations(t)

	// Assert that the returned likes and error match the expected values
	assert.NoError(t, err)
	assert.Equal(t, fakelikes, *resultLikes)
	assert.Equal(t, fakePagination, *resultPagination)
	assert.Equal(t, fakeOffset, resultPagination.Take)
}

func TestSet(t *testing.T) {
	// Create a mock repository
	mockRepo := new(_mock.LikeRepositoryMock)
	mockPostRepo := new(_mock.PostRepositoryMock)

	// Create a like service with the mock repository
	likeService := NewLikeService(mockRepo, mockPostRepo)

	fakeUserUuid := uuid.New().String()
	fakePostUuid := uuid.New().String()
	ctx := context.Background()

	// Create a sample like request
	likeRequest := domain.LikeRequest{
		UserUuid: fakeUserUuid,
		PostUuid: fakePostUuid,
	}

	fakelikes := domain.Like{
		UserUuid: fakeUserUuid,
		PostUuid: fakePostUuid,
	}

	// Expect the Set method to be called with the correct argument
	mockRepo.On("Set", ctx, &fakelikes).Return(nil)
	mockRepo.On("Exist", ctx, likeRequest.UserUuid, likeRequest.PostUuid).Return(false, nil)

	// Expect the Exist method to be called with the correct argument
	mockPostRepo.On("Exist", ctx, fakePostUuid).Return(true, nil)

	// Call the Set method of the like service
	err := likeService.Set(ctx, &likeRequest)

	// Assert that the mock repository's Set method was called with the correct argument
	mockRepo.AssertExpectations(t)
	mockPostRepo.AssertExpectations(t)

	// Assert that the returned like and error match the expected values
	assert.NoError(t, err)
}

func TestUnset(t *testing.T) {
	// Create a mock repository
	mockRepo := new(_mock.LikeRepositoryMock)
	mockPostRepo := new(_mock.PostRepositoryMock)

	// Create a like service with the mock repository
	postService := NewLikeService(mockRepo, mockPostRepo)

	// Define a fake UUID
	fakeUserUuid := uuid.New().String()
	fakePostUuid := uuid.New().String()
	likeRequest := domain.LikeRequest{
		UserUuid: fakeUserUuid,
		PostUuid: fakePostUuid,
	}

	ctx := context.Background()

	// Expect the Exist method to be called with the correct argument
	mockRepo.On("Exist", ctx, fakeUserUuid, fakePostUuid).Return(true, nil)

	// Expect the Unset method to be called with the correct argument
	mockRepo.On("Unset", ctx, fakeUserUuid, fakePostUuid).Return(nil)

	// Call the Unset method of the likes service
	err := postService.Unset(ctx, &likeRequest)

	// Assert that the mock repository's Unset method was called with the correct argument
	mockRepo.AssertExpectations(t)

	// Assert that the returned likes and error match the expected values
	assert.NoError(t, err)
}
