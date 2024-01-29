package internal

import (
	"github.com/PNYwise/like-service/internal/domain"
	"github.com/PNYwise/like-service/internal/handler"
	like_service "github.com/PNYwise/like-service/proto"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
)

func InitGrpc(srv *grpc.Server, extConf *domain.ExtConf, db *pgx.Conn) {
	likeHandlers := handler.NewLikeHandler(extConf)
	like_service.RegisterLikeServer(srv, likeHandlers)
}
