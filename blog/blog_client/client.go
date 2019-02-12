package main

import (
	"context"
	"fmt"
	"github.com/ernestoaparicio/blog-example/blog/blogpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	createBlog(c)
}
func createBlog(client blogpb.BlogServiceClient) {
	blog := &blogpb.Blog{
		AuthorId: "Ernie",
		Title:    "My first blog",
		Content:  "Content of my first blog",
	}
	res, err := client.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Blog has been created: %v", res)
}
