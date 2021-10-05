package main

import (
	"context"
	"log"
	"net"
	pb "protos/service"

	grpc "google.golang.org/grpc"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type gRPCServer struct {
	pb.UnimplementedServiceServer
}

func (g *gRPCServer) NewCourse(ctx context.Context, in *emptypb.Empty) (*pb.CourseReply, error) {
	log.Println("NewCourse() call recieved.")
	reply := pb.CourseReply{
		ReplyType:        pb.CourseReply_OK,
		Id:               0,
		Name:             "name",
		ExpectedWorkload: 9000,
		Rating:           5.0,
	}

	return &reply, nil
}

func (g *gRPCServer) GetCourse(ctx context.Context, in *pb.CourseRequest) (*pb.CourseReply, error) {
	log.Println("GetCourse() call recieved.")
	reply := pb.CourseReply{
		ReplyType:        pb.CourseReply_OK,
		Id:               *in.CourseId,
		Name:             "name",
		ExpectedWorkload: 9000,
		Rating:           5.0,
	}

	return &reply, nil
}

func (g *gRPCServer) DeleteCourse(ctx context.Context, in *pb.CourseRequest) (*pb.CourseReply, error) {
	log.Println("DeleteCourse() call recieved.")
	reply := pb.CourseReply{
		ReplyType:        pb.CourseReply_OK,
		Id:               *in.CourseId,
		Name:             "name",
		ExpectedWorkload: 9000,
		Rating:           5.0,
	}

	return &reply, nil
}

func (g *gRPCServer) EditCourse(ctx context.Context, in *pb.CourseRequest) (*pb.CourseReply, error) {
	log.Println("EditCourse() call recieved.")
	dto := *in.NewCourseState

	reply := pb.CourseReply{
		ReplyType:        pb.CourseReply_OK,
		Id:               *in.CourseId,
		Name:             dto.Name,
		ExpectedWorkload: dto.ExpectedWorkload,
		Rating:           dto.Rating,
	}

	return &reply, nil
}

func newServer() *gRPCServer {
	s := &gRPCServer{}
	return s
}

func main() {
	lis, err := net.Listen("tcp", "localhost:3333")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var options []grpc.ServerOption

	grpcServer := grpc.NewServer(options...)
	pb.RegisterServiceServer(grpcServer, &gRPCServer{})
	grpcServer.Serve(lis)
}
