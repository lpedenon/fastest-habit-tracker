package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
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
		perfectFilePath := filepath.Join(dataDir, "perfect.csv")

		// get records from csv
		records, err := readFromCsv(filePath)
		if err != nil {
			fmt.Println("error reading from csv", err)
			return
		}

		perfectStreak, err := readFromCsv(perfectFilePath)
		if err != nil {
			fmt.Println("error reading from csv", err)
			return
		}

		highestStreak := 0
		var habitOfHighestStreak string

		for _, record := range records[1:] {
			// skip header
			streak, err := strconv.Atoi(record[2])
			if err != nil {
				fmt.Println("error converting streak to int", err)
			}
			if streak > highestStreak {
				highestStreak = streak
				habitOfHighestStreak = record[0]
			}

			fmt.Printf("%s - streak %s\n", record[0], record[2])
		}
		fmt.Printf("\nhighscore: %s %d\n", habitOfHighestStreak, highestStreak)

		perfectStreakNumberStr := perfectStreak[0][0]
		fmt.Printf("Perfect streak: %s\n", perfectStreakNumberStr)

	},
}
