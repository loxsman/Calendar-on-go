package main

import "fmt"

type Month struct {
	id     int
	name   string
	lenght int
	lenghtOfYear int
}

type Day struct {
	id            int
	name          string
	numberOfTusks int
	idDay         int
}

// func calculateDay() string {
// 	for i := 0; i <= 365; i++ {
// 	}
// }

func printMonth(arrMonth [12]Month) {
	for _, p := range arrMonth {
		fmt.Printf("%d. %s with lenght(%d)\n", p.id, p.name, p.lenght)
	}
}

func controls(arrMonth [12]Month) {
	fmt.Println("Chose month")
	var i int
	switch i {
	case 1:

	}
}

func printDays(arrMonth [12]Month, arrDay []string) {
	for i := 1; i <= 12; i++ {
		if i == arrMonth[i].id {
			for j := 1; j <= arrMonth[i].lenght; j++ {
				fmt.Printf("%d. %s curr tasks(%d)\n", j, arrDay[0], p.lenght)
			}
		}
	}
}

func main() {
	arrDayOnWeek := make([]string, 7, 7)
	arrDayOnWeek = append(arrDayOnWeek, "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday")

	arrMonth := [12]Month{
		{1, "January", 31, 31},
		{2, "February", 28, 59},
		{3, "March", 31, 90},
		{4, "April", 30, 120},
		{5, "May", 31, 151},
		{6, "June", 30, 181},
		{7, "July", 31, 212},
		{8, "August", 31, 243},
		{9, "September", 30, 273},
		{10, "October", 31, 304},
		{11, "November", 30, 334},
		{12, "December", 31, 365},
	}

	arrDay := make([]Day, 365, 365)
	countMonth := 1
	for i := 1; i <= 365; i++{
		arrDay = append(arrDay, {i})
	}
	printMonth(arrMonth)
}
