package main

import "fmt"

/* task 2_1 */

func main(){
	var a int32
	var b int32
	fmt.Println("Введите длину прямоугольника")
	fmt.Scanln(&a)
	fmt.Println("Введите ширину прямоугольника")
	fmt.Scanln(&b)
	s:= a*b
	fmt.Println(s)
}