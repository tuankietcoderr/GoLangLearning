package main

import (
	"fmt"
	"time"
)

func IfStatement() int {
	statement := 5
	if statement < 5 {
		return statement
	}
	return 5
}

func IfDeclareAndStatement() int {
	if statement := 5; statement < 5 {
		return statement
	}
	return 5
}

func SwitchStatement() {
	switch os := 1; os {
	case 1:
		fmt.Println(1)
	default:
		fmt.Println("Not")
	}
}

func SwitchNoCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 10:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good night")
	}
}
