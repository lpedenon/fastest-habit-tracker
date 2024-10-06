
package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "path/filepath"
  "os"
  "encoding/csv"
)

var CompleteCmd = &cobra.Command{
  Use: "complete",
  Short: "complete a habit",
  Args: cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {

    changed := false;

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


    records, err := reader.ReadAll()
    if err != nil {
      fmt.Println("error reading csv file", err)
      return
    }

    for _, record := range records[1:] {
      if record[0] == args[0] {
        record[5] = "true";
        changed = true;
      }
    }

    if !changed {
      return;
    }

    file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644);
    if err != nil {
      fmt.Println("Error opening file", err)
      return
    }
    defer file.Close()  // ensure file is closed

    file.Truncate(0)
    file.Seek(0, 0)

    writer := csv.NewWriter(file)
    err = writer.WriteAll(records) // Write all records back
    if err != nil {
      fmt.Println("error writing csv file", err);
    }
    writer.Flush()

  },
}

