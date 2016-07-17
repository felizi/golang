package main 

import "fmt"

func main() {
	foo := myStruct{}
	foo.myField = "bar"
	println(foo.myField)
	fmt.Println(foo)

	bar := myStruct{"foo"}
	println(bar.myField)
	fmt.Println(bar)

	buz := new(myStruct)
	buz.myField = "buz"
	println(buz.myField)
	println(buz)

	boo := &myStruct{"boo"}
	println(boo.myField)
	println(boo)
}

type myStruct struct {
	myField string
}