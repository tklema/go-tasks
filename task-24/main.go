package main

import (
	"PointProject/pointPackage"
	"fmt"
)

func main() {
	pointA := pointPackage.NewPoint(2, 7)
	pointB := pointPackage.NewPoint(10, 11)

	fmt.Println(pointA.Distance(pointB))
}
