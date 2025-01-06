package main

import "fmt"

func vals() (int, int) {
	return 5, 6
}

func main() {
	a, b := vals()
	fmt.Println(a, b)

	_, c := vals()
	fmt.Println(c)
}
