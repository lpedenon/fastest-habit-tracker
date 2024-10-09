package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var CompleteCmd = &cobra.Command{
	Use:   "complete",
	Short: "complete a habit",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		changed := false

		// get current dir
		currentDir, err := os.Getwd()
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
		defer file.Close() // ensure file is closed

		reader := csv.NewReader(file)

		records, err := reader.ReadAll()
		if err != nil {
			fmt.Println("error reading csv file", err)
			return
		}

		for _, record := range records[1:] {
			// if habit we inputted is found
			if record[0] == args[0] {
				record[5] = "true"
				changed = true

				lastCompletedDay := strings.TrimSpace(string(record[4]))
				lastCompletedTime, err := time.Parse("2006-01-02", lastCompletedDay)
				if err != nil {
					fmt.Println("error parsing date", err)
					return
				}
				lastCompletedTime = time.Date(lastCompletedTime.Year(), lastCompletedTime.Month(), lastCompletedTime.Day(), 0, 0, 0, 0, lastCompletedTime.Location())
				now := time.Now()

				yesterday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, -2)

				if lastCompletedTime.Before(yesterday) {
					record[3] = "1"
				} else {
					record[3] = string(int(record[3][0]) + 1)
				}

				record[4] = now.Format("2006-01-02")
			}
		}

		if !changed {
			return
		}

		file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Error opening file", err)
			return
		}
		defer file.Close() // ensure file is closed

		file.Truncate(0)
		file.Seek(0, 0)

		writer := csv.NewWriter(file)
		err = writer.WriteAll(records) // Write all records back
		if err != nil {
			fmt.Println("error writing csv file", err)
		}
		writer.Flush()

	},
}
