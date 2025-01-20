package main

import (
	"fmt"
	"os"

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

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
