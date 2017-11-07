package services

// import (
// 	"io"
// 	"log"
// 	"net"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/jinzhu/gorm"
// 	"github.com/koodinikkarit/seppo/db"
// 	"github.com/koodinikkarit/seppo/matias_service"
// 	"golang.org/x/net/context"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/grpclog"
// 	"google.golang.org/grpc/reflection"
// )

// type MatiasServiceServer struct {
// 	getDB func() *gorm.DB
// }

// func StartMatiasService(
// 	port string,
// 	getDB func() *gorm.DB,
// ) {
// 	lis, err := net.Listen("tcp", ":"+port)
// 	if err != nil {
// 		grpclog.Fatalf("failed to listen: %v", err)
// 	}

// 	s := grpc.NewServer()
// 	MatiasService.RegisterMatiasServer(s, &MatiasServiceServer{
// 		getDB: getDB,
// 	})

// 	reflection.Register(s)
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }

// func (s *MatiasServiceServer) RequestMatiasKey(
// 	ctx context.Context,
// 	in *MatiasService.RequestMatiasKeyRequest,
// ) (
// 	*MatiasService.RequestMatiasKeyResponse,
// 	error,
// ) {
// 	res := &MatiasService.RequestMatiasKeyResponse{}
// 	newDb := s.getDB()

// 	randString, _ := GenerateRandomString(10)

// 	newMatiasClient := db.MatiasClient{
// 		ClientKey: randString,
// 	}

// 	newDb.Create(&newMatiasClient)

// 	res.Key = newMatiasClient.ClientKey

// 	return res, nil
// }

// func (s *MatiasServiceServer) InsertEwSongIds(
// 	stream MatiasService.Matias_InsertEwSongIdsServer,
// ) error {
// 	for {
// 		in, err := stream.Recv()
// 		if err == io.EOF {
// 			return nil
// 		}
// 		if err != nil {
// 			return err
// 		}
// 	}
// }

// func (s *MatiasServiceServer) ChangeEwSongIds(
// 	stream MatiasService.Matias_ChangeEwSongIdsServer,
// ) error {
// 	for {
// 		in, err := stream.Recv()
// 		if err == io.EOF {
// 			return nil
// 		}
// 		if err != nil {
// 			return err
// 		}
// 	}
// }

// func (s *MatiasServiceServer) SyncEwSongs(
// 	stream MatiasService.Matias_SyncEwSongsServer,
// ) error {
// 	for {
// 		in, err := stream.Recv()
// 		if err == io.EOF {
// 			return nil
// 		}
// 		if err != nil {
// 			return err
// 		}
// 	}
// }

// func (s *MatiasServiceServer) RequestEwChanges(
// 	in *MatiasService.RequestEwDatabaseChangesRequest,
// 	stream MatiasService.Matias_RequestEwChangesServer,
// ) error {

// }
