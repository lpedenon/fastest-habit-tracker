package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "path/filepath"
  "os"
  "encoding/csv"
)

var AddCmd = &cobra.Command{
  Use: "add [task]",
  Short: "add a new habit",
  Args: cobra.ExactArgs(5),
  Run: func(cmd *cobra.Command, args []string) {
    task := append(args, "false");
    fmt.Printf("added task %s\n", task)

    // get current dir
    currentDir, err := os.Getwd();
    if err != nil {
      fmt.Println("error getting dir", err)
      return
    }

    // get data dirpath
    dataDir := filepath.Join(currentDir, "data")
    filePath := filepath.Join(dataDir, "habits.csv")

    file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
      fmt.Println("Error opening file", err)
      return
    }
    defer file.Close()  // ensure file is closed

    //create a csv writer
    writer := csv.NewWriter(file)
    defer writer.Flush()  // ensures all data is written (always include this) because writer keeps buffer

    //write task to csv
    err = writer.Write(task)
    if err != nil {
      fmt.Println("error writing to file", err)
    } else {
      fmt.Printf("wrote task: %s\n", task)
    }

  },
}


