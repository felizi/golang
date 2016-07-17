package main

func main() {
	// Simple loop
	for i := 0; i < 5; i++ {
		println(i)
	}

	i := 0
	for {
		i++
		println(i)
		if i > 5 {
			break
		}
	}

	i = 0
	for ;; {
		i++
		println(i)
		if i > 5 {
			break
		}
	}

	// Slice
	mySlice := []string{"foo", "bar", "buz"}
	for idx, v := range mySlice {
		println(idx, v)
	}

	// Map
	myMap := make(map[string]string)
	myMap["first"] = "foo"
	myMap["second"] = "bar"
	myMap["third"] = "buz"

	for k, v := range myMap {
		println(k, v)
	}
}