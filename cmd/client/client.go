package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/brianykl/gopm/internal/utils"
	pb "github.com/brianykl/gopm/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if len(os.Args) < 2 {
		utils.Usage()
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect to daemon:", err)
		return
	}

	defer conn.Close()
	client := pb.NewProcessManagerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	command := os.Args[1]
	switch command {
	case "start":
		err := utils.RunStart(client, ctx, os.Args[2:])
		if err != nil {
			fmt.Println("error:", err)
		}

	case "stop":
		err := utils.RunStop(client, ctx, os.Args[2:])
		if err != nil {
			fmt.Println("error:", err)
		}

	case "list":
		err := utils.RunList(client, ctx, os.Args[2:])
		if err != nil {
			fmt.Println("error:", err)
		}

	default:
		fmt.Println("Unknown command:", command)
	}
}
