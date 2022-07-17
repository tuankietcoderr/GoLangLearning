package main

import "fmt"

func Pointer() {
	i, j := 52, 2000

	p := &i         // p point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through pointer
	fmt.Println(i)  //see the new value of i

	p = &j // p point to j
	*p = *p / 37
	fmt.Println(j)
}
