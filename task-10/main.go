package main

import (
	"fmt"
	"sort"
)

func main() {

	temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	mapAnswer := make(map[int][]float32)

	for _, elem := range temperatures {
		// int округляет к нулю, поэтому для отрицательных и положительных чисел
		// мы получаем ожидаемый результат
		// PS правда, для чисел от -10 до 10 не понятно, что делать (видимо, они
		// в диапазоне с 0)
		index := int(elem/10) * 10
		mapAnswer[index] = append(mapAnswer[index], elem)
	}

	answerKeys := []int{}
	for k := range mapAnswer {
		answerKeys = append(answerKeys, k)
	}
	sort.Ints(answerKeys)

	for _, k := range answerKeys {
		fmt.Printf("%d: %v\n", k, mapAnswer[k])
	}
}
