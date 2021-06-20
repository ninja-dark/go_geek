package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	var a,b,res float64
	var op string

	fmt.Print("Введите первое число")
	fmt.Scanln(&a)
	if a < 0 {fmt.Println("Значени меньше нуля")
		os.Exit(1)
	}


	fmt.Print("Введите второе число")
	fmt.Scanln(&b)
	if b < 0 {fmt.Println("Значени меньше нуля")
		os.Exit(1)
	}

	fmt.Print("Введит арифметическую операцию")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "/":
		if b == 0.0 {
			panic ("error: you tried to divide by zero.")
		}
		res = a / b
	case "*":
		res = a * b
	case "^":
		res = math.Pow(a, b)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Println("Результат выполнения операции: \n", res)
}



