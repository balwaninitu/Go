package main

func sum(a, b int) int {
	return (a + b)
}

func minus(a, b int) int {
	return a - b
}

func multi(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

//modified

func calculator(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}
