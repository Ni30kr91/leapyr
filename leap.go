package main

import "fmt"

func main() {

	var yr int
	fmt.Println("Enter the year:")
	fmt.Scanf("%d", &yr)
	if yr%4 == 0 && yr%100 != 0 || yr%400 == 0 {
		fmt.Println("The year is a leap year")
	} else {
		fmt.Println("The year is not a leap year")
	}
}
