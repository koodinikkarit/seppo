package service

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	"github.com/cskr/pubsub"
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/seppo_service"
)

type SeppoServiceServer struct {
	getDB  func() *gorm.DB
	pubSub *pubsub.PubSub
}

func StartSeppoService(
	getDB func() *gorm.DB,
	pubSub *pubsub.PubSub,
	port string,
) {
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
		getDB:  getDB,
		pubSub: pubSub,
	})

	reflection.Register(s)
	//go s.Serve(lis)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
