package main

import (
	"context"
	"fmt"
	pb "protos/service"

	grpc "google.golang.org/grpc"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type gRPCServer struct {
	pb.UnimplementedServiceServer

	courses []pb.CourseDTO
}

func (g *gRPCServer) NewCourse(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*pb.CourseReply, error) {
	reply := pb.CourseReply{ReplyType: pb.CourseReply_OK,
		Id:               0,
		Name:             "name",
		ExpectedWorkload: 9000,
		Rating:           5.0,
	}

	return &reply, nil
}

func (c *gRPCServer) GetCourse(ctx context.Context, in *pb.CourseRequest, opts ...grpc.CallOption) (*CourseReply, error) {

	return nil, nil
}

func main() {
	a := pb.UnimplementedServiceServer{}
	fmt.Println(a)
}
