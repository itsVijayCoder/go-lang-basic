package main

import "fmt"

func main() {
	//for loop

	for i :=1; i <= 7; i++ {
		fmt.Println(i)
	}

	N := 10
	sum := 0

	for i := 1; i <= N; i++ {
		sum = sum + i
	}
	fmt.Println(sum)

	//switch statement

	daysInNo := 5

	switch daysInNo {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Choose the Days between 1 to 7")
	}
}