package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/bozd4g/fb.grpc/pkg/proto/userservice"
	"github.com/manifoldco/promptui"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	fmt.Println("Connecting to server..")

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	fmt.Println("Connected! \n")

	prompt := promptui.Select{
		Label: "Select a method",
		Items: []string{"Exit", "Call Create method with fake values", "Call GetAll method every 5 seconds"},
	}

	position, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	c := pb.NewUserServiceClient(conn)
	ctx := context.Background()
	for true {
		if position == 0 {
			break
		} else
		if position == 1 {
			callCreate(ctx, c)
		} else
		if position == 2 {
			callGetAll(ctx, c)
		} else {
			break
		}

		position, _, _ = prompt.Run()
	}

	conn.Close()
}

func callCreate(ctx context.Context, client pb.UserServiceClient) {
	fmt.Println("Requesting..")

	r, err := client.Create(ctx, &pb.CreateRequest{
		Name:     fmt.Sprintf("Name%d", rand.Int()),
		Surname:  fmt.Sprintf("Surname%d", rand.Int()),
		Email:    fmt.Sprintf("Email%d", rand.Int()),
		Password: fmt.Sprintf("Password%d", rand.Int()),
	})
	if err != nil {
		log.Fatalf("could not create: %v", err)
	}

	statusCode := r.GetStatusCode()
	fmt.Println(fmt.Sprintf("Response code: %s", http.StatusText(int(statusCode))))
	fmt.Println("\n=============================================")
}

func callGetAll(ctx context.Context, client pb.UserServiceClient) {
	for true {
		fmt.Println("Requesting..")

		r, err := client.GetAll(ctx, &pb.GetAllRequest{})
		if err != nil {
			log.Fatalf("could not getAll: %v", err)
		}

		jsonUsers, err := json.Marshal(r.GetUsers())
		dst := &bytes.Buffer{}
		if err := json.Indent(dst, jsonUsers, "", "  "); err != nil {
			panic(err)
		}

		fmt.Println("[Users]")
		fmt.Println(dst.String())
		fmt.Println("=============================================")

		time.Sleep(5 * time.Second)
	}
}
