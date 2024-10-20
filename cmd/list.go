package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"time"
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

		habits, err := loadHabits(filePath)

		for i, habit := range habits {
			if !habit.isLastCompletedToday() {
				habits[i].completed = false
			}
			if habits[i].completed {
				continue
			}

			for _, day := range habit.daysOfWeek {
				if day == time.Now().Weekday() {
					fmt.Printf("%s - streak %s\n", habit.name, habit.streak)
				}
			}
		}
		habits.save(filePath)
	},
}
