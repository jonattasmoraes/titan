package grpc

import (
	"context"

	pb "github.com/jonattasmoraes/titan/internal/user/infra/proto"
	"github.com/jonattasmoraes/titan/internal/user/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userGrpcServer struct {
	pb.UnimplementedUserServiceServer
	userService *usecase.GetUserByIdUsecase
}

func NewUserGrpcServer(userService *usecase.GetUserByIdUsecase) *userGrpcServer {
	return &userGrpcServer{
		userService: userService,
	}
}

func (s *userGrpcServer) GetUserByID(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.userService.Execute(req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	response := &pb.GetUserResponse{
		Id:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
	}

	return response, nil
}
