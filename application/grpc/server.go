package grpc

import (
	"github.com/codeedu/esquenta-imersao-grpc/application/grpc/pb"
	"github.com/codeedu/esquenta-imersao-grpc/application/grpc/service"
	"github.com/codeedu/esquenta-imersao-grpc/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var ProductList = model.NewProducts()

func StartGrpcServer() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("could not connect", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	productService := service.NewProductGrpcServer(ProductList)
	pb.RegisterProductServiceServer(grpcServer, productService)


	log.Println("gRPC Server has been started")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("could not connect", err)
	}
}
