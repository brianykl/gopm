package utils

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/brianykl/gopm/internal/server"
	pb "github.com/brianykl/gopm/proto"
)

func Usage() {
	fmt.Println("usage: client <start|stop|list> ...")
}

func RunServer(args []string) error {
	// fs := flag.NewFlagSet("bg", flag.ContinueOnError)
	// var runInBg bool
	// fs.BoolVar(&runInBg, "background", false, "run server in background")
	// err := fs.Parse(args)
	// if err != nil {
	// 	return err
	// }

	// subcommand := fs.Args()
	// if len(subcommand) > 1 {
	// 	return fmt.Errorf("usage: gopm init <flag>")
	// }

	server.StartServer()
	return nil
}

func RunServerInBackground(args []string) error {
	cmd := exec.Command(os.Args[0], "init")
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start server in background: %v", err)
	}
	fmt.Printf("Server started in background (PID %d)\n", cmd.Process.Pid)
	return nil
}

func RunStart(client pb.ProcessManagerClient, ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("start", flag.ContinueOnError)
	var autoRestart string
	fs.StringVar(&autoRestart, "auto-restart", "never", "auto restart policy (never|always|on-failure)")

	// e.g. `client start -auto-restart=always myapp ping google.com`
	err := fs.Parse(args)
	if err != nil {
		// fmt.Println("bomba")
		return err
	}

	subcommand := fs.Args()
	if len(subcommand) < 2 {
		return fmt.Errorf("usage: client start <flag> <name> <cmd> [args...]")
	}

	name := subcommand[0]
	cmdToRun := subcommand[1]
	procArgs := subcommand[2:]
	req := &pb.StartRequest{
		Name:        name,
		Command:     cmdToRun,
		Args:        procArgs,
		AutoRestart: autoRestart,
	}
	res, err := client.StartProcess(ctx, req)
	if err != nil {
		fmt.Println(name, cmdToRun, procArgs)
		fmt.Println("bomba")
		return err
	}
	fmt.Println(res.Message)
	return nil
}

func RunStop(client pb.ProcessManagerClient, ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("stop", flag.ContinueOnError)
	var force bool
	fs.BoolVar(&force, "force", false, "force stop the process")

	err := fs.Parse(args)
	if err != nil {
		return err
	}

	subcommand := fs.Args()
	if len(subcommand) < 1 {
		return fmt.Errorf("usage: client stop <flag> <name>")
	}

	name := subcommand[0]
	req := &pb.StopRequest{
		Name:  name,
		Force: force,
	}
	res, err := client.StopProcess(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println(res.Message)
	return nil
}

func RunList(client pb.ProcessManagerClient, ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("list", flag.ContinueOnError)
	var verbose bool
	fs.BoolVar(&verbose, "verbose", false, "show more information")

	err := fs.Parse(args)
	if err != nil {
		return err
	}

	req := &pb.ListRequest{Verbose: verbose}
	res, err := client.ListProcess(ctx, req)
	if err != nil {
		return err
	}
	if len(res.Processes) == 0 {
		fmt.Println("no running processes.")
	} else {
		for _, p := range res.Processes {
			fmt.Printf("name: %s, PID: %d, status: %s\n", p.Name, p.Pid, p.Status)
		}
	}
	return nil
}

func RunLogs(client pb.ProcessManagerClient, ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("log", flag.ContinueOnError)
	var follow bool
	fs.BoolVar(&follow, "follow", false, "follow logs in real time")

	err := fs.Parse(args)
	if err != nil {
		return err
	}

	subcommand := fs.Args()
	if len(subcommand) < 1 || len(subcommand) > 2 {
		return fmt.Errorf("usage: client log <flag>")
	}
	name := subcommand[0]

	req := &pb.LogRequest{Name: name, Follow: follow}
	stream, err := client.StreamLogs(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to open log stream %s", err)
	}

	for {
		line, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error recieving log data: %v", err)
		}
		fmt.Println(line.Text)
	}
	return nil
}

func RunRemove(client pb.ProcessManagerClient, ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("remove", flag.ContinueOnError)
	var follow bool
	fs.BoolVar(&follow, "no-stop", false, "remove process without stopping it")

	err := fs.Parse(args)
	if err != nil {
		return err
	}

	subcommand := fs.Args()
	if len(subcommand) < 1 || len(subcommand) > 3 {
		return fmt.Errorf("usage: client remove <flag> <name>")
	}

	name := subcommand[0]

	req := &pb.RemoveRequest{Name: name}
	res, err := client.RemoveProcess(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println(res.Message)
	return nil
}
