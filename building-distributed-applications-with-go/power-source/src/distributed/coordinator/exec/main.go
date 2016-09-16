package main

import (
	"fmt"

	"./../../coordinator"
)

func main() {
	ql := coordinator.NewQueueListener()
	go ql.ListenForNewSource()
	var a string
	fmt.Scanln(&a)
}
