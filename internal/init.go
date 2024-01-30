package internal

import (
	"github.com/PNYwise/like-service/internal/domain"
	"github.com/PNYwise/like-service/internal/handler"
	"github.com/PNYwise/like-service/internal/repository"
	"github.com/PNYwise/like-service/internal/service"
	like_service "github.com/PNYwise/like-service/proto"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
)

func InitGrpc(srv *grpc.Server, extConf *domain.ExtConf, db *pgx.Conn) {
	likeRepository := repository.NewLikeRepository(db)
	likeService := service.NewLikeService(likeRepository)
	likeHandlers := handler.NewLikeHandler(extConf, likeService)
	like_service.RegisterLikeServer(srv, likeHandlers)
}
