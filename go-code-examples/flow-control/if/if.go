package main

import "fmt"

func main() {
	// Normal if statement
	num := 3
	if num == 3 {
		fmt.Println("num is 3")
	}

	// If statement with initial short statement
	for i := 0; i < 3; i++ {
		if j := 1; j < i {
			fmt.Println("Inside the if statement!")
		}
		fmt.Println("Inside the for loop!")
	}

	// If, else if and else blocks
	if v := 2; v == num {
		fmt.Println("v is equal to num")
	} else if v < num {
		fmt.Println("v is smaller than num")
	} else {
		fmt.Println("v is greater than num")
	}
}
