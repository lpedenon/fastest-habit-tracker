package cmd

import (
	"fmt"
	"os"
	// "path/filepath"
	"encoding/csv"
	"strings"
	"time"
)

type Habit struct {
	name          string
	daysOfWeek    []weekday
	streak        string
	lastCompleted string
	completed     bool
}

type Habits []Habit

type weekday int

const (
	Monday weekday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func mapToWeekday(day string) weekday {
	switch day {
	case "M":
		return Monday
	case "T":
		return Tuesday
	case "W":
		return Wednesday
	case "R":
		return Thursday
	case "F":
		return Friday
	case "S":
		return Saturday
	case "N":
		return Sunday
	default:
		return -1 // Invalid day
	}
}

func (w weekday) String() string {
	switch w {
	case Monday:
		return "M"
	case Tuesday:
		return "T"
	case Wednesday:
		return "W"
	case Thursday:
		return "R"
	case Friday:
		return "F"
	case Saturday:
		return "S"
	case Sunday:
		return "N"
	default:
		return ""
	}
}

func loadHabits(filePath string) (Habits, error) {
	records, err := readFromCsv(filePath)
	if err != nil {
		return nil, err
	}

	habits := make(Habits, 0)
	for _, record := range records[1:] {
		days := strings.Split(record[1], "")
		var daysOfWeek []weekday
		for _, day := range days {
			daysOfWeek = append(daysOfWeek, mapToWeekday(day))
		}
		habit := Habit{
			name:          record[0],
			daysOfWeek:    daysOfWeek,
			streak:        record[2],
			lastCompleted: record[3],
			completed:     record[4] == "true",
		}
		habits = append(habits, habit)
	}
	return habits, nil
}

func (habits Habits) save(filePath string) error {
	data := make([][]string, 0)
	data = append(data, []string{"name", "daysOfWeek", "streak", "lastCompleted", "completed"})

	for _, habit := range habits {
		var daysOfWeekStrings []string
		for _, day := range habit.daysOfWeek {
			daysOfWeekStrings = append(daysOfWeekStrings, day.String())
		}

		data = append(data, []string{
			habit.name,
			strings.Join(daysOfWeekStrings, ""),
			habit.streak,
			habit.lastCompleted,
			fmt.Sprintf("%t", habit.completed),
		})
	}
	return writeToCsv(filePath, data, false)
}

func (habit Habit) isLastCompletedToday() bool {
	return currentDate() == habit.lastCompleted
}

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

func readFromCsv(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}
