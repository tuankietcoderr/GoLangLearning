package main

func ForLoop() (sum int64) {
	for i := 0; i < 10; i++ {
		sum += int64(i)
	}
	for sum < 10 {
		sum += sum
	}
	return sum
}

func WhileLoop() {
	statement := 1
	for statement < 5 {

	}
}

func ForeverLoop() {
	for {

	}
}
