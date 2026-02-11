package main

import "fmt"

func main() {

	// Array declaration and initialization
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	fmt.Println("Array:", arr)

	check := [5]string{"Go", "Python", "Java", "C++", "JavaScript"}
	fmt.Println("String Array:", check)

	// Slice declaration and initialization
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)

	// Slices are dynamic and can be resized
	slice = append(slice, 6)
	fmt.Println("Updated Slice:", slice)

	// Slices can be created from arrays
	subSlice := arr[1:4] // This will create a slice from index 1 to 3 (20, 30, 40)
	fmt.Println("Sub Slice from Array:", subSlice)

	// Modifying the subSlice will affect the original array
	subSlice[0] = 25
	fmt.Println("Modified Sub Slice:", subSlice)
	fmt.Println("Original Array after modification:", arr)

	// Length and capacity of slices
	fmt.Println("Length of slice:", len(slice))
	fmt.Println("Capacity of slice:", cap(slice))

	//make function to create a slice with specific length and capacity
	// newSlice := make([]int, 4, 5) // length 4 and capacity 5
	newSlice := make([]int, 3, 5) // length 0 and capacity 5
	newSlice1 := []int{}
	// newSlice = append(newSlice, 10)
	// newSlice[0] = 1
	// newSlice[1] = 2
	// newSlice[2] = 3
	// newSlice[3] = 4
	// newSlice = append(newSlice, 10, 20, 30, 40,50,60)
	// newSlice = append(newSlice, 10, 20, 30, 40,50,60,70,80,90, 100,110,120,130,140,150,160,170,180,190,200,210)
	fmt.Println("New Slice with make:", newSlice)
	fmt.Println("New Slice1 without make:", newSlice1)
	fmt.Println("Length of new slice:", len(newSlice))
	fmt.Println("Capacity of new slice:", cap(newSlice))

	todos := []string{"Learn GOLang,", "Workout,", "Assign Task for Interns,"}

	more := []string{"WakeUp Early,", "Play Shuttle"}

	todos = append(todos, more...)

	fmt.Println("Appending Two Slices:", todos)

	//for range

	views := []int{20,50,65,35,90}

	totalViews := 0
	averageViewsPerDay := 0

	for i, v := range views{
		fmt.Println("Days", i+1, "Views", v)
		totalViews = totalViews + v
		// averageViewsPerDay = totalViews/i+1
		
	}
	averageViewsPerDay = totalViews/len(views)
	fmt.Println("TotalViews", totalViews, "Average Views/day", averageViewsPerDay)
}