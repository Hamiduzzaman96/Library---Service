package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/Hamiduzzaman96/Library---Service/Library---Service/proto/bookpb"
	handlergrpc "github.com/Hamiduzzaman96/Library---Service/internal/handler/grpc" //given alias for same grpc name
	"github.com/Hamiduzzaman96/Library---Service/internal/repository/mysql"
	"github.com/Hamiduzzaman96/Library---Service/internal/usecase"
	"github.com/Hamiduzzaman96/Library---Service/pkg/database"
)

func main() {
	db := database.NewMySQL()
	repo := mysql.NewBookMySQlRepository(db)
	uc := usecase.NewBookUsecase(repo)
	bookHandler := handlergrpc.NewBookHandler(uc)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBookServiceServer(grpcServer, bookHandler)

	log.Println("Library gRPC server running on :50051")
	grpcServer.Serve(lis)

}
