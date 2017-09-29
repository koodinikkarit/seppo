package seppo

import (
	"log"
	"net"

	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/matias_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

type MatiasServiceServer struct {
	databaseService *SeppoDB.DatabaseService
}

func NewMatiasService(port string, databaseService *SeppoDB.DatabaseService) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	MatiasService.RegisterMatiasServer(s, &MatiasServiceServer{
		databaseService: databaseService,
	})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
