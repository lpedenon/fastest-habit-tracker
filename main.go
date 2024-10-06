package main

import (
	// "fmt"
	// "math"
	//  "math/rand"
	// "strconv"
	"github.com/lpedenon/habit/cmd"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "habit",
		Short: "Simple habit tracker",
	}

	rootCmd.AddCommand(cmd.AddCmd)
	rootCmd.AddCommand(cmd.ListCmd)
	rootCmd.AddCommand(cmd.CompleteCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}
