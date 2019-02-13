package main

import (
	"context"
	"fmt"
	"github.com/ernestoaparicio/blog-example/blog/blogpb"
	"google.golang.org/grpc"
	"io"
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

	//createBlog(c)
	//readBlog(c)
	//updateBlog(c)
	//deleteBlog(c)
	listBlogs(c)
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

func readBlog(client blogpb.BlogServiceClient) {
	res, err := client.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "5c6352e46e429469273571d4"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Blog has been read: %v", res)
}

func updateBlog(client blogpb.BlogServiceClient) {
	blog := &blogpb.Blog{
		Id:       "5c6352e46e429469273571d4",
		AuthorId: "Ernie",
		Title:    "My updated blog 2",
		Content:  "Content of my updated blog",
	}
	res, err := client.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Blog has been updated: %v", res)
}

func deleteBlog(client blogpb.BlogServiceClient) {
	res, err := client.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: "5c63532e6e429469273571d5"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Blog id %v has been deleted.", res.BlogId)
}

func listBlogs(client blogpb.BlogServiceClient) {
	fmt.Sprintf("List all blogs.")

	stream, err := client.ListBlog(context.Background(), &blogpb.ListBlogRequest{})

	if err != nil {
		log.Fatalf("Error while calling ListBlog service.")
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetBlog())
	}
}
