package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Month struct {
	id     int
	name   string
	length int
}

type Day struct {
	id            int
	name          string
	numberOfTasks int
	tasks         []string
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

func printDays(arrDays [12][]Day, arrMonth [12]Month, n int) {
	for i := 0; i < arrMonth[n].length; i++ {
		fmt.Printf("%d. %s, Tasks(%d)\n", arrDays[n][i].id, arrDays[n][i].name, arrDays[n][i].numberOfTasks)
	}
}

func controls(arrMonth [12]Month, arrDays [12][]Day) {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	for {
		clearConsole()
		printMonth(arrMonth)
		fmt.Println("Chose month, enter '0' to end program")
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
			printDays(arrDays, arrMonth, n)
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
				for i, task := range arrDays[n][k].tasks {
					fmt.Printf("%d. %s\n", i+1, task)
				}
				fmt.Println("\nEnter new task or enter '0' to go back")
				s, _ := reader.ReadString('\n')
				s = strings.TrimSpace(s)
				if s == "0" {
					break
				}
				if len(s) != 0 {
					arrDays[n][k].tasks = append(arrDays[n][k].tasks, s)
					arrDays[n][k].numberOfTasks++
				}
			}
		}
	}
}

func createMonths() [12]Month {
	return [12]Month{
		{1, "January", 31},
		{2, "February", 29},
		{3, "March", 31},
		{4, "April", 30},
		{5, "May", 31},
		{6, "June", 30},
		{7, "July", 31},
		{8, "August", 31},
		{9, "September", 30},
		{10, "October", 31},
		{11, "November", 30},
		{12, "December", 31},
	}
}

func createDays(arrMonth [12]Month) [12][]Day {
	arrDayOnWeek := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	var arrDays [12][]Day
	dayOfWeekIndex := 0
	for i := 0; i < len(arrMonth); i++ {
		arrDays[i] = make([]Day, arrMonth[i].length)
		for j := 0; j < arrMonth[i].length; j++ {
			arrDays[i][j] = Day{
				id:            j + 1,
				name:          arrDayOnWeek[dayOfWeekIndex%7],
				numberOfTasks: 0,
				tasks:         []string{},
			}
			dayOfWeekIndex++
		}
	}
	return arrDays
}

func main() {
	arrMonth := createMonths()
	arrDays := createDays(arrMonth)
	controls(arrMonth, arrDays)
}
