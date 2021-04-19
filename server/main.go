package main

import (
	"context"
	"fmt"
	pb "github.com/memeoAmazonas/gRPC-protobuf/proto"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	pb.UnimplementedWishListServiceServer
}

func (s *server) Create(_ context.Context, req *pb.CreateWishListReq) (*pb.CreateWishListResp, error) {
	fmt.Println("creating the wish list " + req.WishList.Name)
	return &pb.CreateWishListResp{
		WishListId: req.WishList.Id,
	}, nil
}

func (s *server) Add(_ context.Context, _ *pb.AddItemReq) (*pb.AddItemResp, error) {
	return nil, nil
}

func (s *server) List(_ context.Context, _ *pb.ListWishListReq) (*pb.ListWishListResp, error) {
	return nil, nil
}
func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}
	serv := grpc.NewServer()
	pb.RegisterWishListServiceServer(serv, &server{})
	if err = serv.Serve(listener); err != nil {
		panic("cannot initializer" + err.Error())
	}

}
