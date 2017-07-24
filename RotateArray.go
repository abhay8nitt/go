package main

import "fmt"
/**
  Rotate an array of size N by given steps
*/
func main() {
	input :=[]int{1,2,3,4,5}
	steps:=3
	rotateArray(input,steps)
	rotateArray([]int{},steps)
	fmt.Println("Resultant array:",input)
}

func rotateArray(input []int,steps int){
	if len(input) <=0 {
		return
	}
	rotate(input,0, len(input)-1)
	rotate(input,0, steps-1)
	rotate(input, steps ,  len(input)-1)
}

func rotate(input []int, start int, end int){
	if start > end {
		return
	}
	for start<end{
		temp:=input[start]
		input[start] = input[end]
		input[end] = temp
		start++
		end--
	}
}


