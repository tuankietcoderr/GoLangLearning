package main

import "fmt"

func main() {
	// slices is a reference
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s, len(s), cap(s))
	capacity := 10
	length := 5
	s = make([]int, length, capacity)
	fmt.Println(s)
	s = append(s, 5)
	fmt.Println(s)
}
