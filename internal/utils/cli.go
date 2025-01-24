package utils

import (
	"context"
	"flag"
	"fmt"

	pb "github.com/brianykl/gopm/proto"
)

func Usage() {
	fmt.Println("usage: client <start|stop|list> ...")
}

func RunStart(client pb.ProcessManagerClient, ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("start", flag.ContinueOnError)
	var autoRestart string
	fs.StringVar(&autoRestart, "auto-restart", "never", "auto restart policy (never|always|on-failure)")

	// e.g. `client start -auto-restart=always myapp ping google.com`
	err := fs.Parse(args)
	if err != nil {
		fmt.Println("bomba")
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

func ParseFlag(args []string) {

}
