package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Day struct {
	ID            int      `json:"id,"`
	Name          string   `json:"name"`
	NumberOfTasks int      `json:"number_of_tasks"`
	Tasks         []string `json:"tasks"`
}

type Month struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Length int    `json:"length"`
	Days   []Day  `json:"days"`
}

func saveToFile(data [12]Month, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	return err
}

func loadFromFile(filename string) ([12]Month, error) {
	var data [12]Month
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return data, nil
		}
		return data, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return data, err
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(buffer, &data)
	return data, err
}

func clearConsole() {
	var clearCmd *exec.Cmd
	switch runtime.GOOS {
	case "linux", "darwin":
		clearCmd = exec.Command("clear")
	case "windows":
		clearCmd = exec.Command("cmd", "/c", "cls")
	default:
		return
	}
	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}

func printMonth(arrMonth [12]Month) {
	for _, p := range arrMonth {
		fmt.Printf("%d. %s with length(%d)\n", p.ID, p.Name, p.Length)
	}
}

func printDays(month Month) {
	for i := 0; i < month.Length; i++ {
		fmt.Printf("%d. %s, Tasks(%d)\n", month.Days[i].ID, month.Days[i].Name, month.Days[i].NumberOfTasks)
	}
}

func controls(arrMonth [12]Month) {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	for {
		clearConsole()
		printMonth(arrMonth)
		fmt.Println("Choose month, enter '0' to end program")
		fmt.Scan(&n)
		if n == 0 {
			saveToFile(arrMonth, "data.json")
			os.Exit(0)
		}
		if n < 1 || n > 12 {
			fmt.Println("Invalid month. Please try again.")
			continue
		}
		n-- // Convert to zero-based index
		for {
			clearConsole()
			printDays(arrMonth[n])
			fmt.Println("Choose day to add or view your tasks, or enter '0' to go back to the previous menu")
			fmt.Scan(&k)
			if k == 0 {
				break
			}
			if k < 1 || k > arrMonth[n].Length {
				fmt.Println("Invalid day. Please try again.")
				continue
			}
			k-- // Convert to zero-based index
			for {
				clearConsole()
				fmt.Printf("Tasks for %d %s:\n", k+1, arrMonth[n].Name)
				for i, task := range arrMonth[n].Days[k].Tasks {
					fmt.Printf("%d. %s\n", i+1, task)
				}
				fmt.Println("\nEnter new task or enter '0' to go back")
				s, _ := reader.ReadString('\n')
				s = strings.TrimSpace(s)
				if s == "0" {
					break
				}
				if len(s) != 0 {
					arrMonth[n].Days[k].Tasks = append(arrMonth[n].Days[k].Tasks, s)
					arrMonth[n].Days[k].NumberOfTasks++
				}
			}
		}
	}
}

func createMonths() [12]Month {
	return [12]Month{
		{1, "January", 31, nil},
		{2, "February", 29, nil},
		{3, "March", 31, nil},
		{4, "April", 30, nil},
		{5, "May", 31, nil},
		{6, "June", 30, nil},
		{7, "July", 31, nil},
		{8, "August", 31, nil},
		{9, "September", 30, nil},
		{10, "October", 31, nil},
		{11, "November", 30, nil},
		{12, "December", 31, nil},
	}
}

func createDays(arrMonth [12]Month) [12]Month {
	arrDayOnWeek := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	dayOfWeekIndex := 0
	for i := 0; i < len(arrMonth); i++ {
		arrMonth[i].Days = make([]Day, arrMonth[i].Length)
		for j := 0; j < arrMonth[i].Length; j++ {
			arrMonth[i].Days[j] = Day{
				ID:            j + 1,
				Name:          arrDayOnWeek[dayOfWeekIndex%7],
				NumberOfTasks: 0,
				Tasks:         []string{},
			}
			dayOfWeekIndex++
		}
	}
	return arrMonth
}

func main() {
	filename := "data.json"
	arrMonth, err := loadFromFile(filename)
	if err != nil {
		fmt.Println("Error loading data:", err)
		arrMonth = createMonths()
		arrMonth = createDays(arrMonth)
	} else if arrMonth[0].Days == nil {
		arrMonth = createDays(arrMonth)
	}
	controls(arrMonth)
}
