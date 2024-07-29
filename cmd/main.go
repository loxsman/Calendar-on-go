package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Day struct {
	id            int
	name          string
	numberOfTasks int
	tasks         []string
}

type Month struct {
	id     int
	name   string
	length int
	days   []Day
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
		fmt.Printf("%d. %s with length(%d)\n", p.id, p.name, p.length)
	}
}

func printDays(month Month) {
	for i := 0; i < month.length; i++ {
		fmt.Printf("%d. %s, Tasks(%d)\n", month.days[i].id, month.days[i].name, month.days[i].numberOfTasks)
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
			if k < 1 || k > arrMonth[n].length {
				fmt.Println("Invalid day. Please try again.")
				continue
			}
			k-- // Convert to zero-based index
			for {
				clearConsole()
				fmt.Printf("Tasks for %d %s:\n", k+1, arrMonth[n].name)
				for i, task := range arrMonth[n].days[k].tasks {
					fmt.Printf("%d. %s\n", i+1, task)
				}
				fmt.Println("\nEnter new task or enter '0' to go back")
				s, _ := reader.ReadString('\n')
				s = strings.TrimSpace(s)
				if s == "0" {
					break
				}
				if len(s) != 0 {
					arrMonth[n].days[k].tasks = append(arrMonth[n].days[k].tasks, s)
					arrMonth[n].days[k].numberOfTasks++
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
	arrDayOnWeek := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	dayOfWeekIndex := 0
	for i := 0; i < len(arrMonth); i++ {
		arrMonth[i].days = make([]Day, arrMonth[i].length)
		for j := 0; j < arrMonth[i].length; j++ {
			arrMonth[i].days[j] = Day{
				id:            j + 1,
				name:          arrDayOnWeek[dayOfWeekIndex%7],
				numberOfTasks: 0,
				tasks:         []string{},
			}
			dayOfWeekIndex++
		}
	}
	return arrMonth
}

func main() {
	arrMonth := createMonths()
	arrMonth = createDays(arrMonth)
	controls(arrMonth)
}
