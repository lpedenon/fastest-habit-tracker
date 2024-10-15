package cmd

import (
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

		habits, err := loadHabits(filePath)
		if err != nil {
			fmt.Println("error loading habits into structs", err)
		}
		for i, habit := range habits {
			if habit.name == args[0] {
				habits[i].completed = true
				changed = true

				lastCompletedDay := strings.TrimSpace(habit.lastCompleted)
				lastCompletedTime, err := time.Parse("2006-01-02", lastCompletedDay)
				if err != nil {
					fmt.Println("error parsing date", err)
					return
				}
				lastCompletedTime = time.Date(lastCompletedTime.Year(), lastCompletedTime.Month(), lastCompletedTime.Day(), 0, 0, 0, 0, lastCompletedTime.Location())
				now := time.Now()

				yesterday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, -2)

				if lastCompletedTime.Before(yesterday) {
					habits[i].streak = "1"
				} else {
					habits[i].streak = string(int(habits[i].streak[0]) + 1)
				}

				habits[i].lastCompleted = now.Format("2006-01-02")
			}
		}

		if !changed {
			return
		}
		habits.save(filePath)

	},
}
