package main

import "fmt"

type Month struct {
	id     int
	name   string
	lenght int
}

type Day struct {
	id            int
	name          string
	numberOfTusks int
}

// func calculateDay() string {
// 	for i := 0; i <= 365; i++ {

// 	}
// }

func printTable(arrMonth [12]Month) {
	for _, p := range arrMonth {
		fmt.Printf("%d. %s x with lenght(%d)\n", p.id, p.name, p.lenght)
	}
}

func main() {
	arrDay := make([]string, 7, 7)
	arrDay = append(arrDay, "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday")

	arrMonth := [...]Month{
		{1, "January", 31},
		{2, "February", 28},
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
	printTable(arrMonth)
}
