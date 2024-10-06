package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "list current methods",
	Run: func(cmd *cobra.Command, args []string) {

		// get current dir
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println("error getting dir", err)
			return
		}

		// get data dirpath
		dataDir := filepath.Join(currentDir, "data")
		filePath := filepath.Join(dataDir, "habits.csv")
		lastUpdatedFilePath := filepath.Join(dataDir, "last_updated.txt")

		// get records from csv
		records, err := readFromCsv(filePath)
		if err != nil {
			fmt.Println("error reading from csv", err)
			return
		}

		if !isLastUpdatedToday(lastUpdatedFilePath) {
			err = writeLastUpdatedDate(lastUpdatedFilePath)
			if err != nil {
				fmt.Println("error writing last updated date", err)
				return
			}

			for _, record := range records {
				record[5] = "false"
			}
			err = writeToCsv(filePath, records, false)
		}

		for i, record := range records {
			// skip header
			if i == 0 {
				continue
			}
			if record[5] == "false" {
				fmt.Printf("%d: %s\n", i, record[0])
			}
		}

	},
}
