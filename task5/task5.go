package main

import (
	"fmt"
	"os"
)


func fibonacci(a uint) uint{
	if a == 0 {
		return 0
	}
	if a == 1 {
		return 1
	}
	return fibonacci(a - 1) + fibonacci(a - 2)
}

func main() {
	var a uint
	//fibonacciMap := map[int]int{}
	fmt.Print("Введите число для вычисления: ")
	if _, err := fmt.Scanln(&a); err != nil {
		fmt.Println("Введено не число")
		os.Exit(1)
	}
	fmt.Println("Result:", fibonacci(a))
}


