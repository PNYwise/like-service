package handler

import (
	"context"

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
func (*likeHandler) GetByPostUuid(context.Context, *like_service.QueryLikeRequest) (*like_service.LikeResponse, error) {
	panic("unimplemented")
}

// Set implements like_service.LikeServer.
func (*likeHandler) Set(context.Context, *like_service.LikeRequest) (*emptypb.Empty, error) {
	panic("unimplemented")
}

// Unset implements like_service.LikeServer.
func (*likeHandler) Unset(context.Context, *like_service.LikeRequest) (*emptypb.Empty, error) {
	panic("unimplemented")
}
