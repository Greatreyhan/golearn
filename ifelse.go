package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("is even")
	} else {
		fmt.Println("is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("All is even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, " is Negative")
	} else if num < 10 {
		fmt.Println(num, " has 1 digit")
	} else {
		fmt.Println(num, " has multiple digits")
	}
}
