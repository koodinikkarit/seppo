package seppo

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	"github.com/koodinikkarit/seppo/seppo_service"

	"github.com/koodinikkarit/seppo/db"
)

type SeppoServiceServer struct {
	databaseService *SeppoDB.DatabaseService
}

func CreateSeppoService(port string, databaseService *SeppoDB.DatabaseService) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// creds, err := credentials.NewServerTLSFromFile("./ssl/server.crt", "./ssl/server.key")
	// if err != nil {
	// 	log.Fatalf("Failed to generate credentials %v", err)
	// }

	s := grpc.NewServer()
	SeppoService.RegisterSeppoServer(s, &SeppoServiceServer{
		databaseService: databaseService,
	})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
