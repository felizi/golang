package main 

func main() {
	foo := 1
	if foo == 1 {
		println("bar")
	} else {
		println("buz")
	}

	if soo := 1; soo == 1 {
		println("bar")
	} else {
		println("buz")
	}

	switch bar := 2; bar {
	case 1:
		println("one")
	case 2:
		println("two")
	}

	boo := 3
	switch {
	case boo == 1:
		println("one")
	case boo == 2:
		println("two")
	case boo > 2:
		println("sup two")
	}
}