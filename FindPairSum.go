package main

import "fmt"


func main() {
  X,Y:= getPairInSortedArray([]int{1,2,3,4,5},9)
  fmt.Println("X =",X,",Y=",Y)
}

/**
Given a sorted array and a sum find a pair (X,Y) such that X+Y = sum
*/
func getPairInSortedArray(input []int, sum int) (X int, Y int) {
	X = -1
	Y = -1
	size:= len(input)
	if size == 0 {
		return
	}
	low  := 0
	high := size - 1

	for low <= high{
		temp:= input[low] + input[high]
		if temp == sum{
			X = input[low]
			Y = input[high]
			break
		}else if temp < sum{
			low++
		}else{
			high--
		}
	}
	return
}
