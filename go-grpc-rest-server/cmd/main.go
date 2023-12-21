package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/digvijay17july/golang-projects/go-grpc-rest-server/app/proto"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	proto.UnimplementedUserServiceServer
}

func (grpcServer *GrpcServer) GetUser(ctx context.Context, req *proto.Request) (*proto.User, error) {
	return &proto.User{
		Id:    req.Id,
		Name:  "Orion",
		Email: "orion@example.com",
	}, nil

}

func main() {
	list, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, &GrpcServer{})
	go func() {
		if err := s.Serve(list); err != nil {
			panic(err)
		}
	}()
	startHttpServer()
}

func startHttpServer() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))

}
