package handler

import (
	"context"
	"log"

	"errors"

	"github.com/PNYwise/like-service/internal/domain"
	like_service "github.com/PNYwise/like-service/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

type likeHandler struct {
	like_service.UnimplementedLikeServer
	likeService domain.ILikeService
	extConf     *domain.ExtConf
}

func NewLikeHandler(extConf *domain.ExtConf, likeService domain.ILikeService) *likeHandler {
	return &likeHandler{
		extConf:     extConf,
		likeService: likeService,
	}
}

// GetByPostUuid implements like_service.LikeServer.
func (l *likeHandler) GetByPostUuid(ctx context.Context, request *like_service.QueryLikeRequest) (*like_service.LikeResponse, error) {
	data, pagination, err := l.likeService.GetByPostUuid(ctx, request.GetPostUuid(), request.GetPage())
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("internal server error")
	}
	response := make([]*like_service.UserLikeResponse, len(*data))

	for i, resp := range *data {
		protoResp := &like_service.UserLikeResponse{
			UserUuid: resp.UserUuid,
			Name:     "Jhon Doe",
		}
		response[i] = protoResp
	}

	return &like_service.LikeResponse{
		PostUuid:         request.GetPostUuid(),
		UserLikeResponse: response,
		Pagiation: &like_service.Pagination{
			Page:        request.GetPage(),
			TotalRecord: pagination.ItemCount,
			TotalPage:   pagination.PageCount,
		},
	}, nil
}

// Set implements like_service.LikeServer.
func (l *likeHandler) Set(ctx context.Context, request *like_service.LikeRequest) (*emptypb.Empty, error) {
	setLikeRequest := &domain.SetLikeRequest{
		UserUuid: request.GetUserUuid(),
		PostUuid: request.GetPostUuid(),
	}
	if err := l.likeService.Set(ctx, setLikeRequest); err != nil {
		return nil, errors.New(err.Error())
	}
	return &empty.Empty{}, nil
}

// Unset implements like_service.LikeServer.
func (l *likeHandler) Unset(ctx context.Context, request *like_service.LikeRequest) (*emptypb.Empty, error) {
	if err := l.likeService.Unset(ctx, request.GetUserUuid(), request.GetPostUuid()); err != nil {
		return nil, errors.New(err.Error())
	}
	return &empty.Empty{}, nil
}
