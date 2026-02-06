package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	myScore := 121

	sqRoot := math.Sqrt(144)
	sqRoot1 := math.Sqrt(float64(myScore))

	fmt.Println("Square Root of my Score:", sqRoot)
	fmt.Println("Square Root of my Score:", sqRoot1)
	
	firstName := "Vijay"
	lastName := "Subramanian"
	
	myFullName := firstName + " " + lastName
	
	 fmt.Println("My Name is ", myFullName)
	 fmt.Println("My Name is ", strings.ToUpper(myFullName), " in UPPERCASE")

}