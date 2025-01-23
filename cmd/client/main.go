package main

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/brianykl/gopm/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: client <start|stop|list> ...")
		return
	}
	command := os.Args[1]

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect to daemon:", err)
		return
	}
	defer conn.Close()
	client := pb.NewProcessManagerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	switch command {
	case "start":
		if len(os.Args) < 4 {
			fmt.Println("Usage: client start <name> <cmd> [args...]")
			return
		}
		name := os.Args[2]
		cmdToRun := os.Args[3]
		args := []string{}
		if len(os.Args) > 4 {
			args = os.Args[4:]
		}
		req := &pb.StartRequest{
			Name:    name,
			Command: cmdToRun,
			Args:    args,
		}
		res, err := client.StartProcess(ctx, req)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(res.Message)

	case "stop":
		if len(os.Args) < 3 {
			fmt.Println("Usage: client stop <name>")
			return
		}
		name := os.Args[2]
		req := &pb.StopRequest{Name: name}
		res, err := client.StopProcess(ctx, req)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(res.Message)

	case "list":
		req := &pb.ListRequest{}
		res, err := client.ListProcess(ctx, req)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if len(res.Processes) == 0 {
			fmt.Println("No running processes.")
		} else {
			for _, p := range res.Processes {
				fmt.Printf("Name: %s, PID: %d, Status: %s\n", p.Name, p.Pid, p.Status)
			}
		}

	default:
		fmt.Println("Unknown command:", command)
	}
}
