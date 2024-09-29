package main

import (
  // "fmt"
  // "math"
  //  "math/rand"
  // "strconv"
  "github.com/spf13/cobra"
  "os"
) 

func main() {
  rootCmd := &cobra.Command{
    Use: "habit",
    Short: "Simple habit tracker",
  }

  if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
  
}
