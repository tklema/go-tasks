package main

import (
	"fmt"
)

func remove(arr []int, index int) []int {
	copy(arr[index:], arr[index+1:])
	return arr[:len(arr)-1]
}

func main() {

	ind := 2

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr)
	for range len(arr) - ind {
		arr = remove(arr, ind)
		fmt.Println(arr)
	}
}
