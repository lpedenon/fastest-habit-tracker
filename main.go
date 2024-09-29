package main

import (
  // "fmt"
  // "math"
  //  "math/rand"
  // "strconv"
  "github.com/spf13/cobra"
  "os"
  "github.com/lpedenon/habit/cmd"
) 

func main() {
  rootCmd := &cobra.Command{
    Use: "habit",
    Short: "Simple habit tracker",
  }

  rootCmd.AddCommand(cmd.AddCmd)

  if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
  
}
