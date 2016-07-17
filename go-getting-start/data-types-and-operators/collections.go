package main

import "fmt"

func main() {
	// Array
	myArray := [...] int{42, 37, 99}
	/*myArray[0] = 42
	myArray[1] = 27
	myArray[2] = 99*/
	fmt.Println(myArray)

	// Slice of array
	mySlice := myArray[:] // [2:] [:3]
	mySlice = append(mySlice, 100)

	fmt.Println(mySlice)
	//fmt.Println(len(myArray))

	// Slice
	mySliceMake := make([]float32, 100)
	mySliceMake[0] = 12.
	mySliceMake[1] = 15.
	mySliceMake[2] = 18.
	fmt.Println(mySliceMake)
	fmt.Println(len(mySliceMake))

	// Slice
	mySliceNew := []float32{14., 15., 19.}
	fmt.Println(mySliceNew)
	fmt.Println(len(mySliceNew))

	// Map
	myMap := make(map[int]string)
	fmt.Println(myMap)
	myMap[0] = "Foo"
	myMap[1] = "Bar"
	fmt.Println(myMap)
	fmt.Println(myMap[1])
	fmt.Println(myMap[99])

}