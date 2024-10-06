package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "path/filepath"
  "os"
  "encoding/csv"
)

var ListCmd = &cobra.Command{
  Use: "list",
  Short: "list current methods",
  Run: func(cmd *cobra.Command, args []string) {

    // get current dir
    currentDir, err := os.Getwd();
    if err != nil {
      fmt.Println("error getting dir", err)
      return
    }

    // get data dirpath
    dataDir := filepath.Join(currentDir, "data")
    filePath := filepath.Join(dataDir, "habits.csv")

    file, err := os.Open(filePath)
    if err != nil {
      fmt.Println("Error opening file", err)
      return
    }
    defer file.Close()  // ensure file is closed

    reader := csv.NewReader(file)

    header, err := reader.Read()
    if err != nil {
        fmt.Println("Error reading header:", err)
        return
    }
    fmt.Println("Header:", header)

    records, err := reader.ReadAll()
    if err != nil {
      fmt.Println("error reading csv file", err)
      return
    }

    for i, record := range records {
      if record[5] == "false" {
        fmt.Printf("%d: %s\n", i+1, record)
      }
    }

  },
}
