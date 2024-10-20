package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var PerfectCmd = &cobra.Command{
	Use:   "perfect",
	Short: "show perfect days streak",
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

		habits, err := loadHabits(filePath)
		if err != nil {
			fmt.Println("error loading habits into structs", err)
			return
		}

		for _, habit := range habits {
			if !habit.completed {
				for _, day := range habit.daysOfWeek {
					if day == time.Now().Weekday() {
						fmt.Print("You have not completed all habits today")
						return
					}
				}
			}
		}

		perfectStreak, err := readFromCsv(perfectFilePath)
		if err != nil {
			fmt.Println("error reading from csv", err)
			return
		}

		perfectStreakNumberStr := perfectStreak[0][0]

		perfectStreakNumber, err := strconv.Atoi(perfectStreakNumberStr)
		if err != nil {
			fmt.Println("error converting streak to int", err)
			return
		}

		perfectStreakNumber++

		perfectStreak[0][0] = strconv.Itoa(perfectStreakNumber)

		fmt.Printf("Congratulations! You have completed all habits today. Your perfect streak is now %d\n", perfectStreakNumber)

		writeToCsv(perfectFilePath, perfectStreak, false)

	},
}
