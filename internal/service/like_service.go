package service

import (
	"context"
	"errors"
	"sync"

	"github.com/PNYwise/like-service/internal/domain"
	"github.com/PNYwise/like-service/internal/util"
)

func NewLikeService(likeRepo domain.ILikeRepository, postRepo domain.IPostRepository) domain.ILikeService {
	return &likeService{
		likeRepo: likeRepo,
		postRepo: postRepo,
	}
}

type likeService struct {
	likeRepo domain.ILikeRepository
	postRepo domain.IPostRepository
}

// GetByPostUuid implements domain.ILikeService.
func (l *likeService) GetByPostUuid(ctx context.Context, postUuid string, page uint64) (*[]domain.Like, *domain.Pagination, error) {
	if errs := util.Var(postUuid, "required,uuid4"); len(errs) > 0 && errs[0].Error {
		return nil, nil, util.ValidationErrMsg(errs)
	}
	likes, limit, err := l.likeRepo.GetByPostUuid(ctx, postUuid, page)
	if err != nil {
		return nil, nil, errors.New("Internal Server Error")
	}
	return likes, &domain.Pagination{
		Take:      limit,
		ItemCount: uint64(0),
		PageCount: uint64(0),
	}, nil
}

// Set implements domain.ILikeService.
func (l *likeService) Set(ctx context.Context, request *domain.LikeRequest) error {
	type setValidation struct {
		types  string
		result bool
		errors error
	}

	var (
		wg     sync.WaitGroup
		result = make(chan setValidation, 2)
		once   sync.Once
	)

	if errs := util.Validate(request); len(errs) > 0 && errs[0].Error {
		return util.ValidationErrMsg(errs)
	}

	validate := func(types string, defaults bool, existFunc func(ctx context.Context) (bool, error)) {
		defer wg.Done()
		exist, err := existFunc(ctx)
		if err != nil {
			result <- setValidation{types: types, result: defaults, errors: err}
			return
		}
		result <- setValidation{types: types, result: exist}
	}
	wg.Add(2)
	go validate("post", false, func(ctx context.Context) (bool, error) {
		return l.postRepo.Exist(context.Background(), request.PostUuid)
	})
	go validate("like", true, func(ctx context.Context) (bool, error) {
		return l.likeRepo.Exist(ctx, request.UserUuid, request.PostUuid)
	})
	go func() {
		wg.Wait()
		once.Do(func() {
			close(result)
		})
	}()
	for res := range result {
		switch res.types {
		case "post":
			if res.errors != nil {
				return errors.New("Internal Server Error")
			} else if !res.result {
				return errors.New("Post Not Found")
			}
		case "like":
			if res.errors != nil {
				return errors.New("Internal Server Error")
			} else if res.result {
				return errors.New("The post has been liked")
			}
		}
	}

	like := &domain.Like{
		UserUuid: request.UserUuid,
		PostUuid: request.PostUuid,
	}
	if err := l.likeRepo.Set(ctx, like); err != nil {
		return errors.New("Internal Server Error")
	}
	return nil
}

// Unset implements domain.ILikeService.
func (l *likeService) Unset(ctx context.Context, request *domain.LikeRequest) error {
	if errs := util.Validate(request); len(errs) > 0 && errs[0].Error {
		return util.ValidationErrMsg(errs)
	}
	exist, err := l.likeRepo.Exist(ctx, request.UserUuid, request.PostUuid)
	if err != nil {
		return errors.New("Internal Server Error")
	}
	if !exist {
		return errors.New("Like not found")
	}
	if err := l.likeRepo.Unset(ctx, request.UserUuid, request.PostUuid); err != nil {
		return errors.New("Internal Server Error")
	}
	return nil
}
