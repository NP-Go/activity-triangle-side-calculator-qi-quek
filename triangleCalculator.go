package main

import (
	"fmt"
	"math"
)

func calculateSide(length1, length2, angle float64) float64 {

	//Insert the code here
	angleRad := angle * ((math.Pi) / 180)
	length3 := math.Sqrt(math.Pow(length1, 2) + math.Pow(length2, 2) - 2*length1*length2*math.Cos(angleRad))
	//

	//Do not remove this line
	fmt.Printf("The 3rd length of the triange is %.02f", length3)
	return length3
}

func main() {
	var length1 float64
	var length2 float64
	var angle float64

	fmt.Println("Enter the first length of the triangle: ")
	fmt.Scanln(&length1)
	fmt.Println("Enter the second length of the triangle: ")
	fmt.Scanln(&length2)
	fmt.Println("Enter the angle between the two lengths: ")
	fmt.Scanln(&angle)

	calculateSide(length1, length2, angle)
}
