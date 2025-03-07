package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pm "github.com/brianykl/gopm/internal/process"
	pb "github.com/brianykl/gopm/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProcessManagerServer struct {
	pb.UnimplementedProcessManagerServer
	manager *pm.ProcessManager
}

func NewProcessManagerServer(manager *pm.ProcessManager) *ProcessManagerServer {
	return &ProcessManagerServer{manager: manager}
}

func (pms *ProcessManagerServer) StartProcess(ctx context.Context, req *pb.StartRequest) (*pb.ProcessResponse, error) {
	pi, err := pms.manager.StartProcess(req.Name, req.AutoRestart, req.Command, req.Args...)
	if err != nil {
		return &pb.ProcessResponse{
			Success: false,
			Message: fmt.Sprintf("failed to start process: %v", err),
		}, err
	}
	return &pb.ProcessResponse{
		Success: true,
		Message: fmt.Sprintf("process %s started", pi.Name),
	}, nil
}

func (pms *ProcessManagerServer) StopProcess(ctx context.Context, req *pb.StopRequest) (*pb.ProcessResponse, error) {
	pi, err := pms.manager.GetProcess(req.Name)
	if err != nil {
		return &pb.ProcessResponse{
			Success: false,
			Message: fmt.Sprintf("invalid process name: %v", err),
		}, err
	}
	fmt.Printf("bongo pi=%+v\n", pi)
	err = pms.manager.StopProcess(pi, req.Force)
	if err != nil {
		return &pb.ProcessResponse{
			Success: false,
			Message: fmt.Sprintf("failed to stop process: %v", err),
		}, nil
	}

	return &pb.ProcessResponse{
		Success: true,
		Message: fmt.Sprintf("process %s stopped", pi.Name),
	}, nil
}

func (pms *ProcessManagerServer) ListProcess(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	processes, err := pms.manager.ListProcesses(req.Verbose)
	if err != nil {
		return nil, err
	}

	var pbProcesses []*pb.ProcessInfo
	for _, process := range processes {
		pbProcesses = append(pbProcesses, &pb.ProcessInfo{
			Name:   process.Name,
			Pid:    int32(process.PID),
			Status: process.Status,
		})
	}

	return &pb.ListResponse{Processes: pbProcesses}, nil
}

func (pms *ProcessManagerServer) StreamLogs(req *pb.LogRequest, stream pb.ProcessManager_StreamLogsServer) error {
	name := req.Name
	follow := req.Follow

	channel, ok := pms.manager.GetLogChannels()[name]
	if !ok {
		return status.Errorf(codes.NotFound, "no logs to process %s", name)
	}

	for {
		select {
		case line, open := <-channel:
			if !open {
				return nil
			}
			err := stream.Send(&pb.LogLine{Text: line})
			if err != nil {
				return err
			}

		default:
			if !follow {
				return nil
			}

			time.Sleep(100 * time.Millisecond)
		}
	}

}

func (pms *ProcessManagerServer) RemoveProcess(ctx context.Context, req *pb.RemoveRequest) (*pb.ProcessResponse, error) {
	pi, err := pms.manager.GetProcess(req.Name)
	if err != nil {
		return &pb.ProcessResponse{
			Success: false,
			Message: fmt.Sprintf("invalid process name: %v", err),
		}, err
	}
	fmt.Printf("pi=%+v\n", pi)

	if pi.Status == "running" {

	} else {
		err = pms.manager.RemoveProcess(pi)
	}

	return &pb.ProcessResponse{
		Success: true,
		Message: fmt.Sprintf("process %s removed", pi.Name),
	}, nil
}

func StartServer() {
	manager := pm.NewProcessManager()
	grpcServer := grpc.NewServer()
	pb.RegisterProcessManagerServer(grpcServer, NewProcessManagerServer(manager))

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	fmt.Println("process manager daemon listening on 50051...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
