package main

const (
	first = 1 << (10 * iota)
	second //= "the second"
	third //= "the third"
)
func main() {
	println(first)
	println(second)
	println(third)

}