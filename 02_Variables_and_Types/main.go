package main

import "fmt"

func main() {
	var name string = "Vijay"
	var age int= 25
	var isLearning bool = true
	var height float32 = 5.9
	var myCompany string = "Asthrix"

	fmt.Println("Name:", name)
	fmt.Println("Company:", myCompany )
	fmt.Println("Age:", age)
	fmt.Println("Is Learning Go:", isLearning)
	fmt.Println("Height:", height)

	//short declaration

	currentCompany := "Susanoox"
	working,year := "Developer", 2

	fmt.Printf("currently working in %s as a %s for %d years\n", currentCompany, working, year)


	isAdmin := true
	isLogged := false
	hasSubscription := true

	// Logical AND
	canAccess := isAdmin && isLogged && hasSubscription
	fmt.Println("Can Access:", canAccess)

	// Logical OR
	canView := isAdmin || isLogged || hasSubscription
	fmt.Println("Can View:", canView)

	// Logical NOT
	isGuest := !isLogged
	fmt.Println("Is Guest:", isGuest)




}