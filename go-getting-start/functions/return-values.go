package main 

func main() {
	println("... Single Return Values ...")
	sum := addSingleReturn(1,3,5,9)
	println(sum)

	println("... Multiple Return Values ...")
	numTerms, sum := addMultipleReturn(1,3,5,9)
	println("Add:", numTerms, "terms for a total of", sum)
	
	println("... Named Return Parameters ...")
	numTerms, sum = addNamedParameters(1,3,5,9)
	println("Add:", numTerms, "terms for a total of", sum)
	
}

func addSingleReturn(terms ...int) int {
	result := 0
	for _, term := range terms {
		result += term
	}

	return result
}

func addMultipleReturn(terms ...int) (int, int) {
	result := 0
	for _, term := range terms {
		result += term
	}

	return len(terms), result
}

func addNamedParameters(terms ...int) (numTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}
	numTerms = len(terms)

	return
}