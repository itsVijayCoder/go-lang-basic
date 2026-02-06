package main

import (
	"fmt"
)

func main() {
	score:= 85

	if score >=90 {
		fmt.Println("Excellent! You scored an A.")
	} else if score >= 80 {
		fmt.Println("Great job! You scored a B.")
	} else if score >= 70 {
		fmt.Println("Good effort! You scored a C.")
	} else if score >= 60 {
		fmt.Println("You passed! You scored a D.")
	} else {
		fmt.Println("Better luck next time! You scored an F.")
	}

	// Short statement in if and give explanation
	if percentage := float64(score) / 100 * 100; percentage >= 90 {
		fmt.Printf("Your percentage is %.2f%%. Excellent! You scored an A.\n", percentage)
	} else if percentage := float64(score) / 100 * 100; percentage >= 80 {
		fmt.Printf("Your percentage is %.2f%%. Great job! You scored a B.\n", percentage)
	} else if percentage := float64(score) / 100 * 100; percentage >= 70 {
		fmt.Printf("Your percentage is %.2f%%. Good effort! You scored a C.\n", percentage)
	} else if percentage := float64(score) / 100 * 100; percentage >= 60 {
		fmt.Printf("Your percentage is %.2f%%. You passed! You scored a D.\n", percentage)
	} else {
		fmt.Printf("Your percentage is %.2f%%. Better luck next time! You scored an F.\n", float64(score)/100*100)
	}
}