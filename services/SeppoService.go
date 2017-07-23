package seppo

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	SeppoService "github.com/koodinikkarit/seppo/seppo_service"

	"github.com/koodinikkarit/seppo/db"
)

type SeppoServiceServer struct {
	createSongChannel chan seppo.CreateSongInput
	databaseService   *seppo.DatabaseService
}

func CreateSeppoService(port string, databaseService *seppo.DatabaseService) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	SeppoService.RegisterSeppoServer(s, &SeppoServiceServer{
		databaseService: databaseService,
	})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
