package matias

import (
	"log"
	"net"

	"github.com/cskr/pubsub"
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/go-clientlibs/matias"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type MatiasServiceServer struct {
	getDB  func() *gorm.DB
	pubSub *pubsub.PubSub
}

func StartMatiasService(
	getDB func() *gorm.DB,
	pubSub *pubsub.PubSub,
	port string,
) {
	lis, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Printf("Matias service listen failed port %v error %v", port, err)
	}

	s := grpc.NewServer()
	MatiasService.RegisterMatiasServer(s, &MatiasServiceServer{
		getDB:  getDB,
		pubSub: pubSub,
	})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
