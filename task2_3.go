package main

import (
	"fmt"
)

/* task 2_3 */

func main(){
	var a int32
	fmt.Println("Введите 3х значное число")
	fmt.Scanln(&a)
	s:= a/100
	d:= (a/10)%10
	e:= a%10
	fmt.Println("Сотни", s)
	fmt.Println("Десятки", d)
	fmt.Println("Единицы", e)

}