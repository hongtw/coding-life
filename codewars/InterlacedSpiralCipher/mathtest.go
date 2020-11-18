package main

import (
	"fmt"
	"math"
)

func main() {
	qtr := math.Pi / 2
	sin, cos := math.Sin, math.Cos
	fmt.Println("========= Sin ==========")
	fmt.Println(math.Round(sin(0)))
	fmt.Println(math.Round(sin(1 * qtr)))
	fmt.Println(math.Round(sin(2 * qtr)))
	fmt.Println(math.Round(sin(3 * qtr)))
	fmt.Println(math.Round(sin(4 * qtr)))
	fmt.Println(math.Round(sin(5 * qtr)))
	fmt.Println(math.Round(sin(6 * qtr)))
	fmt.Println(math.Round(sin(7 * qtr)))
	fmt.Println(math.Round(sin(8 * qtr)))

	fmt.Println("=========== Cos ==========")
	fmt.Println(cos(0))
	// fmt.Println(math.Cos(math.Pi / 2))
	// fmt.Println(math.Cos(math.Pi * 2 / 2))
	// fmt.Println(math.Cos(math.Pi * 3 / 2))

	// fmt.Println("=========== Tan ==========")
	// fmt.Println(math.Tan(0))
	// fmt.Println(math.Tan(math.Pi / 2))
	// fmt.Println(math.Tan(math.Pi * 2 / 2))
	// fmt.Println(math.Tan(math.Pi * 3 / 2))

	// fmt.Println("=========== (Cos - Sin)/2 + 0.5  ==========")
	// fmt.Println((math.Cos(0)-math.Sin(0))/2 + 0.5)
	// fmt.Println((math.Cos(math.Pi/2)-math.Sin(math.Pi/2))/2 + 0.5)
	// fmt.Println(math.Round((math.Cos(math.Pi*2/2)-math.Sin(math.Pi*2/2))/2 + 0.5))
	// fmt.Println((math.Round(math.Cos(math.Pi*3/2)-math.Sin(math.Pi*3/2))/2 + 0.5))
}
