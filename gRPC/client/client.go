package main

import (
	"context"
	"log"
	pb "protos/service"
	"time"

	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	address := "127.0.0.1:3333"

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var id int32 = 1

	// NewCourse
	newCourseReply, err := c.NewCourse(ctx, &emptypb.Empty{})

	handleError(err, "new course")
	log.Printf("NewCourse: %s", newCourseReply)

	// GetCourse call
	getCourseReply, err := c.GetCourse(ctx, &pb.CourseRequest{
		Action:   pb.CourseRequest_GET,
		CourseId: &id,
	})

	handleError(err, "edit course")
	log.Printf("GetCourse: %s", getCourseReply)

	// DeleteCourse call
	deleteCourseReply, err := c.DeleteCourse(ctx, &pb.CourseRequest{
		Action:   pb.CourseRequest_DELETE,
		CourseId: &id,
	})

	handleError(err, "delete course")
	log.Printf("DeleteCourse: %s", deleteCourseReply)

	newCourseState := pb.CourseDTO{
		Id:               id + 1,
		Name:             "Distributed Systems",
		ExpectedWorkload: 1000000,
		Rating:           5.0,
	}
	editCourseReply, err := c.EditCourse(ctx, &pb.CourseRequest{
		Action:         pb.CourseRequest_PUT,
		CourseId:       &id,
		NewCourseState: &newCourseState,
	})

	handleError(err, "edit course")
	log.Printf("EditCourse: %s", editCourseReply)
}

func handleError(err error, annotation string) {
	if err != nil {
		log.Fatalf("could not get  %s: %v", annotation, err)
	}
}
