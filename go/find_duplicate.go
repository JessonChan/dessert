package main

import "fmt"

func find_duplicate(ia []int) bool {
	for i := 0; i < len(ia); {
		if ia[i] < 0 {
			return true
		}
		t := ia[i]
		ia[i] = -1
		i = t
	}
	return false
}

func main() {
	ia := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 8}
	fmt.Println(find_duplicate(ia))
}
