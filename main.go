package main

import "fmt"

func main() {
	fmt.Println("Hello, World")

	var name string = "Renu Bhati"
	var age int = 25
	fmt.Printf("Name: %s\nAge: %d\n", name, age)

	batsman := "David"
	runs := 34
	fmt.Printf("Batsmen: %s\nRuns: %d\n", batsman, runs)

	if runs >= 50 {
		fmt.Println("Half-century!")
	} else {
		fmt.Println("Not a half-century!")

	}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println(add(5, 6))
}
func add(a int, b int) int {
	return a + b
}

