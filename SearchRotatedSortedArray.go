package main

import "fmt"

/**
 Find the index of an element in a rotated sorted array
 Input Array: [3,4,5,1,2]
 Search key: 5
 Output: 2
 Expected time complexity - O(logN)
 Expected space complexity - O(1)
 */
func main(){
    input:=  []int{3,4,5,1,2}
	key:= 5
	index := search(input,key)
	fmt.Println("Found at: ",index)

	key = 1
	index = search(input,key)
	fmt.Println("Found at: ",index)

	key = 6
	index = search(input,key)
	fmt.Println("Found at:",index)
}

func search(input []int, key int) (index int){
	if len(input) <= 0{
		index = -1
		return
	}
	low:= 0
	high:=len(input)-1
	index = -1
	for low < high {
		mid:= low + (high-low)/2;
		if input[mid] == key{
			index = mid
			return index
		}
		//Left subarray is sorted
		if input[low] <= input[mid]{
			   if key >= input[low] && key < input[mid]{
				   high = mid - 1
			   }else{
				   low  =  mid + 1
			   }
		}else{
			//Right subarray is sorted
			if key >= input[mid] && key < input[high]{
				low = mid + 1
			}else{
				high  =  mid - 1
			}
		}
	}
	return
}
