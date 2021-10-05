package main

import (
	"context"
	"log"
	pb "protos/service"
	"time"

	grpc "google.golang.org/grpc"
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
	r, err := c.GetCourse(ctx, &pb.CourseRequest{
		Action:   pb.CourseRequest_GET,
		CourseId: &id,
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r)
}
