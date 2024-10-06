package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var AddCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "add a new habit",
	Args:  cobra.ExactArgs(5),
	Run: func(cmd *cobra.Command, args []string) {
		task := [][]string{append(args, "false")}
		fmt.Printf("added task %s\n", task)

		// get current dir
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println("error getting dir", err)
			return
		}

		// get data dirpath
		dataDir := filepath.Join(currentDir, "data")
		filePath := filepath.Join(dataDir, "habits.csv")

		err = writeToCsv(filePath, task, true)
		if err != nil {
			fmt.Println("error writing to csv")
		}

	},
}
