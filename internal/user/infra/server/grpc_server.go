package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpcService "github.com/jonattasmoraes/titan/internal/user/infra/grpc"
	pb "github.com/jonattasmoraes/titan/internal/user/infra/proto"
	"github.com/jonattasmoraes/titan/internal/user/usecase"
)

func StartGrpcServer(getUserById *usecase.GetUserByIdUsecase) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, grpcService.NewUserGrpcServer(getUserById))

	reflection.Register(grpcServer)

	log.Printf("Listening and serving GRPC on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
