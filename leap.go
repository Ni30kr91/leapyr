package main

import "fmt"

func main() {

	var year int
	fmt.Println("Enter the year:")
	fmt.Scanf("%d", &year)
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Println("The year is a leap year")
	} else {
		fmt.Println("The year is not a leap year")
	}
}
