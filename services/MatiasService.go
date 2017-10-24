package services

import (
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/matias_service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

type MatiasServiceServer struct {
	getDB func() *gorm.DB
}

func StartMatiasService(
	port string,
	getDB func() *gorm.DB,
) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	MatiasService.RegisterMatiasServer(s, &MatiasServiceServer{
		getDB: getDB,
	})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *MatiasServiceServer) InsertEwSongIds(
	ctx context.Context,
	in *MatiasService.InsertEwSongIdsRequest,
) (
	*MatiasService.InsertEwSongIdsResponse,
	error,
) {
	res := &MatiasService.InsertEwSongIdsResponse{}
	return res, nil
}

func (s *MatiasServiceServer) ChangeEwSongIds(
	ctx context.Context,
	in *MatiasService.ChangeEwSongIdsRequest,
) (
	*MatiasService.ChangeEwSongIdsResponse,
	error,
) {
	res := &MatiasService.ChangeEwSongIdsResponse{}
	return res, nil
}

func (s *MatiasServiceServer) SyncEwDatabase(
	ctx context.Context,
	in *MatiasService.SyncEwDatabaseRequest,
) (
	*MatiasService.SyncEwDatabaseResponse,
	error,
) {
	res := &MatiasService.SyncEwDatabaseResponse{}
	return res, nil
}
