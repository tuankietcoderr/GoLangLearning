package main

import "fmt"

func Defer() {
	v := 5
	defer fmt.Println(v) // Stack
	v += 2
	fmt.Println(v)
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	} // 9 -> 0
}
