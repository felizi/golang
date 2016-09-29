package main

import (
	"./elements"
	"fmt"
)

func main () {
	// Board 12 columns x 4 rows, with code length 8 and don't permit repetitions
	board := elements.Create(12, 4, 8, false)
	fmt.Println(board)

	// Board 12 columns x 4 rows, with code length 8 and permit repetitions
	boardWithRepetition := elements.Create(12, 4, 8, true)
	fmt.Println(boardWithRepetition)
	fmt.Println(boardWithRepetition.Check([]int {1,2,3,4}))


}