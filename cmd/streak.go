package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var StreakCmd = &cobra.Command{
	Use:   "streak",
	Short: "show streaks of all habits",
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

		// get records from csv
		records, err := readFromCsv(filePath)
		if err != nil {
			fmt.Println("error reading from csv", err)
			return
		}

		for i, record := range records {
			// skip header
			if i == 0 {
				continue
			}
			fmt.Printf("%s - streak %s\n", record[0], record[3])
		}

	},
}
