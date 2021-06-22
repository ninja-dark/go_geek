package main

import (
	"fmt"
)
func main() {
	//var input int
	//fmt.Print("Введите числа\n")
	//fmt.Scanln(&input)
	slice := []int{9,5,3,1,8,56}
	//slice = append(slice, input)
	insertionSort(slice)
}
func insertionSort(arr []int){
	len := len(arr)
	for i := 1; i < len; i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr [i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}

	fmt.Println("Sorted")
	for _, val := range arr {
		fmt.Println(val)
	}
}

