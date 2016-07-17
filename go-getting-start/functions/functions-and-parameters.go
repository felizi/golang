package main

func main() {
	println("... Passing by value ...")
	message := "Hello"
	passingByValue(message)
	println(message)

	println("... Passing by reference ...")
	passingByReference(&message)
	println(message)

	println("... Variadic functions ...")
	variadicFunction("Hello", "Go", "from", "Pluralsight")
}

func passingByValue(message string) {
	println(message)
	message = "Hello Go"
}

func passingByReference(message *string) {
	println(*message)
	*message = "Hello Go"
}

func variadicFunction(messages ...string) {
	for _, message := range messages {
		println(message)
	}
}