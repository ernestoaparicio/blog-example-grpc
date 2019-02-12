package main

import (
	"context"
	"fmt"
	"github.com/ernestoaparicio/blog-example/blog/blogpb"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

var collection *mongo.Collection

type server struct {
}

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string            	`bson:"author_id"`
	Content  string            	`bson:"content"`
	Title    string            	`bson:"title"`
}

func main() {
	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Connecting to MongoDB")

	// connect to mongodb
	client, err := mongo.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Blog Service Started")

	collection = client.Database("mydb").Collection("blog")

	lis, err := net.Listen("tcp", "0000:50051")

	if err != nil {
		log.Fatalf("Failed to listen to: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)

	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {
		fmt.Println("Starting Server...")
		err2 := s.Serve(lis)
		if err2 != nil {
			log.Fatalf("Failed to serve: %v", err2)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("Closing MongoDB Connection")
	client.Disconnect(context.TODO())
	fmt.Println("End of Program")
}
