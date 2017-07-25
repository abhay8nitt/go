package main

import (
	"fmt"
	"../go/problems"
)

func main(){
	_problems()
}

/**
  Test data
 */
func _problems(){
	input :=[]int{1,2,3,4,5}
	steps:=3
	problems.RotateArray(input,steps)
	problems.RotateArray([]int{},steps)
	fmt.Println("Resultant array:",input)

	input =  []int{3,4,5,1,2}
	key:= 5
	index := problems.Search(input,key)
	fmt.Println("Found at: ",index)

	key = 1
	index = problems.Search(input,key)
	fmt.Println("Found at: ",index)

	key = 6
	index = problems.Search(input,key)
	fmt.Println("Found at:",index)

	X,Y:= problems.GetPairInSortedArray([]int{1,2,3,4,5},9)
	fmt.Println("X =",X,",Y=",Y)
}
