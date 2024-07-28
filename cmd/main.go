package main

import (
	"fmt"
)

type Month struct {
	id           int
	name         string
	lenght       int
	lenghtOfYear int
}

type Day struct {
	id            int
	name          string
	numberOfTasks int
}

func printMonth(arrMonth [12]Month) {
	for _, p := range arrMonth {
		fmt.Printf("%d. %s with lenght(%d)\n", p.id, p.name, p.lenght)
	}
}

func controls(arrMonth [12]Month, arrDays [][]Day) {
	fmt.Println("Chose month")
	var n int
	fmt.Scan(&n)
	for i := 1; i <= arrMonth[n-1].lenght; i++ {
		fmt.Printf("%d. %s, Task(%d)\n", arrDays[n][i].id, arrDays[n][i].name, arrDays[n][i].numberOfTasks)
	}
}

func main() {
	arrDayOnWeek := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	arrMonth := [12]Month{
		{1, "January", 31, 31},
		{2, "February", 29, 59},
		{3, "March", 31, 91},
		{4, "April", 30, 121},
		{5, "May", 31, 152},
		{6, "June", 30, 183},
		{7, "July", 31, 213},
		{8, "August", 31, 244},
		{9, "September", 30, 274},
		{10, "October", 31, 305},
		{11, "November", 30, 335},
		{12, "December", 31, 366},
	}

	arrDays := make([][]Day, 12)
	dayOfWeekIndex := 0
	for i := 0; i < len(arrMonth); i++ {
		arrDays[i] = make([]Day, arrMonth[i].lenght)
		for j := 0; j < arrMonth[i].lenght; j++ {
			arrDays[i][j] = Day{
				id:            j,
				name:          arrDayOnWeek[dayOfWeekIndex%7],
				numberOfTasks: 0,
			}
			dayOfWeekIndex++
		}
	}
	printMonth(arrMonth)
	controls(arrMonth, arrDays)
}
