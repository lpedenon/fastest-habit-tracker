package cmd

import (
	"fmt"
	"os"
	// "path/filepath"
	"strings"
	"time"
)

func currentDate() string {
	return time.Now().Format("2006-01-02")
}

func readLastUpdatedDate(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func writeLastUpdatedDate(filePath string) error {
	return os.WriteFile(filePath, []byte(currentDate()), 0644)
}

func isLastUpdatedToday(filePath string) bool {
	lastDate, err := readLastUpdatedDate(filePath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return currentDate() == lastDate
}

func writeToCsv(filePath string, data [][]string, append bool) error {
	var file *os.File
	var err error

	if append {
		file, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	} else {
		file, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	}

	if err != nil {
		return err
	}
	defer file.Close()

	for _, record := range data {
		_, err = fmt.Fprintln(file, strings.Join(record, ","))
		if err != nil {
			return err
		}
	}
	return nil
}
