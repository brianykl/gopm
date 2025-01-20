package main

import (
	"fmt"
	"os"

	"github.com/brianykl/gopm/internal/process"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "gopm",
		Short: "a simple process manager in go",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello world!")
		},
	}
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "start a process",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				fmt.Println("usage: gopm start <name> <command> [args...]")
				return
			}
			name := args[0]
			command := args[1]
			processArgs := args[2:]
			_, err := process.StartProcess(name, command, processArgs...)
			if err != nil {
				fmt.Println(err)
			}
		},
	}

	// stopCmd := &cobra.Command{
	// 	Use:   "stop",
	// 	Short: "stop a process",
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		if len(args) < 1 {
	// 			fmt.Println("usage: gopm stop <name>")
	// 			return
	// 		}
	// 		name := args[0]
	// 		err := process.StopProcess(name)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 	},
	// }

	rootCmd.AddCommand(startCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
