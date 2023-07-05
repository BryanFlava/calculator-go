package main

import "fmt"

func main() {
	var x, y int
	var operation string

	fmt.Println("Enter the first numbers:")
	fmt.Scanf("%d %d", &y)

	fmt.Println("Enter the seconds numbers:")
	fmt.Scanf("%d %d", &x)

	fmt.Println("Enter an operation (+, -, *, /):")
	fmt.Scanf("%s", &operation)

	switch operation {
	case "+":
		result := x + y
		fmt.Println("The result is:", result)
	case "-":
		result := x - y
		fmt.Println("The result is:", result)
	case "*":
		result := x * y
		fmt.Println("The result is:", result)
	case "/":
		result := x / y
		fmt.Println("The result is:", result)
	default:
		fmt.Println("Invalid operation.")
	}
}
