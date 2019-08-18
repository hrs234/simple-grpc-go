package main

import (
	"context"
	"../proto"
	"net"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	// mendefinisikan port server
	listener, err := net.Listen("tcp", ":4040")
	
	// error handling, ketika port gagal dibuat
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}

}

// fungsi untuk menjalankan penjumlahan
func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	// mendapatkan dari main.go yang berada di folder client
	a, b := request.GetA(), request.GetB()

	// logic dari sebuah function
	result := a + b

	// output dari sebuah proses logic
	return &proto.Response{Result: result}, nil
}

// fungsi untuk menjalankan perkalian
func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &proto.Response{Result: result}, nil
}