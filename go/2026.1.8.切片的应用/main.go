package main

import (
	"fmt"
	"math"
)

func maximum (numbers ...float64) float64{
	max := math.Inf(-1)
	for _,number := range numbers{
		if number > max{
			max =number
		}
	}
	return max
}

func inRange (min float64 , max float64 ,numbers ...float64) [] float64{
	var result [] float64
	for _,number := range numbers{
		if number >= min && number <= max {
			result = append(result,number)
		}
	}
	return result
}

func average(numbers ...float64) float64 {
	var sum float64 =0 
	for _,number := range numbers {
		sum += number
	}
	return sum/float64(len(numbers))
}

func main(){

	fmt.Println(maximum(234.342,432.434,789.563))
	fmt.Println(maximum(234.342,432.434,789.563,832.578))

	fmt.Println(inRange(1,100,12.5,32.8,0,-12.5))
	fmt.Println(inRange(122,140,132.5,333.8,340,-312.5))

	fmt.Println(average(122,140,132.5,333.8,340,-312.5))
	fmt.Println(average(1,100,12.5,32.8,0,-12.5))
}