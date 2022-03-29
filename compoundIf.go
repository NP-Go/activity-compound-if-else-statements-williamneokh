package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Insert your code here

	var year int
	for {
		fmt.Println("Key in the year to check if it's a leap year")
		_, _ = fmt.Scanln(&year)
		checkYear(year)
		fmt.Println("------------")
	}

}
func checkYear(year int) string {
	if year%4 == 0 {
		if year%100 != 0 {
			fmt.Println(year, "is a leap year")
			return strconv.Itoa(year) + " is a leap year"
		} else if year%100 == 0 {
			if year%400 == 0 {
				fmt.Println(year, "is a leap year")
				return strconv.Itoa(year) + " is a leap year"
			} else {
				fmt.Println(year, "is not a leap year")
				return strconv.Itoa(year) + " is not a leap year"
			}
		}
	} else {

		fmt.Println(year, "is not a leap year")
		return strconv.Itoa(year) + " is not a leap year"
	}
	return strconv.Itoa(year) + " is not a leap year"
}
