package main 

func main() {
	addFunc := func(terms ...int) (numTerms int, sum int) {
		for _, term := range terms {
			sum += term
		}
		numTerms = len(terms)

		return
	}
	
	println("... Anonymous Functions ...")
	numTerms, sum := addFunc(1,3,5,9)
	println("Add:", numTerms, "terms for a total of", sum)
	
}