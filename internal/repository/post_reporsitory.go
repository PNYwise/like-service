package repository

import (
	"context"
	"log"

	"github.com/PNYwise/like-service/internal/domain"
	like_service "github.com/PNYwise/like-service/proto"
	"google.golang.org/grpc"
)

type postRepository struct {
	postClient like_service.PostClient
}

func NewPostRepository(postClient *grpc.ClientConn) domain.IPostRepository {
	return &postRepository{
		postClient: like_service.NewPostClient(postClient),
	}
}

// Exist implements domain.IPostRepository.
func (p *postRepository) Exist(ctx context.Context, postUuid string) (bool, error) {
	in := &like_service.Uuid{Uuid: postUuid}
	response, err := p.postClient.Exist(ctx, in)
	if err != nil {
		log.Fatalf("Error executing Grpc: %v", err)
		return false, err
	}
	return response.GetValue(), nil
}
