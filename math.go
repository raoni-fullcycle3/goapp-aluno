package main

import "fmt"

func main() {
	fmt.Println(Soma(10, 10))
	fmt.Println(Subtracao(10, 10))
	fmt.Println(Multiplicacao(10, 10))
	fmt.Println(Divisao(100, 10))
}

func Soma(a int, b int) int {
	return a + b
}

func Subtracao(a int, b int) int {
	return a - b
}

func Multiplicacao(a int, b int) int {
	return a * b
}

func Divisao(a int, b int) int {
	return a / b
}
