package main

import (
	"fmt"
	"math"
)

/* task 2_2 */

func main(){
	var a float64
	var b float64 = 2
	fmt.Println("Введите площадь круга")
	fmt.Scanln(&a)
	c:= (math.Sqrt (a/math.Pi))
	d:= b*c
	l:= math.Pi*d
	fmt.Println("Диаметр окружности", d)
	fmt.Println("Длина окружности", l)
}