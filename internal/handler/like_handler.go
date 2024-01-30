package handler

import (
	"context"
	"log"

	"errors"

	"github.com/PNYwise/like-service/internal/domain"
	like_service "github.com/PNYwise/like-service/proto"
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
func (*likeHandler) Set(context.Context, *like_service.LikeRequest) (*emptypb.Empty, error) {
	panic("unimplemented")
}

// Unset implements like_service.LikeServer.
func (*likeHandler) Unset(context.Context, *like_service.LikeRequest) (*emptypb.Empty, error) {
	panic("unimplemented")
}
