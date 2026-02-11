package main

import "fmt"

func main() {
	//map[keytype]valueType{}

	users := map[string]string{
		"u1": "Vijay",
		"u2": "Ish",
		"u3": "Tom",
	}

	fmt.Println("Users", users)
	fmt.Println("Users1", users["u1"])

	marks := make(map[string]int)

	marks["maths"] = 90
	marks["english"] = 80

	fmt.Println("Marks", marks)

	delete(users, "u3")

	fmt.Println("After Deleting the User3 ", users)
}
