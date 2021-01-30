package service

import (
	"context"
	"github.com/codeedu/esquenta-imersao-grpc/application/grpc/pb"
	"github.com/codeedu/esquenta-imersao-grpc/model"
	"time"
)

type ProductGrpcServer struct {
	pb.UnimplementedProductServiceServer
	Products *model.Products
}

func (p *ProductGrpcServer) CreateProduct(ctx context.Context, in *pb.Product) (*pb.ProductResult, error) {
	product := model.NewProduct()
	product.Name = in.Name
	p.Products.Add(product)

	return &pb.ProductResult{
		Id:   product.ID,
		Name: product.Name,
	}, nil
}

func (p *ProductGrpcServer) List(req *pb.Empty, stream pb.ProductService_ListServer) error {
	for _, product := range p.Products.Product {
		time.Sleep(time.Second * 5)
		stream.Send(&pb.ProductResult{Name: product.Name, Id: product.ID})
	}
	return nil
}

func NewProductGrpcServer(products *model.Products) *ProductGrpcServer {
	return &ProductGrpcServer{
		Products: products,
	}
}
