package main

import "fmt"


func main(){
	items := []int{10,20,45,62,12,10,19,22,35,40,66}
	key := 35
	fmt.Println(binarySearch(items, key))
}

func binarySearch(array_list[] int, key int)bool{

	low := 0
	high := len(array_list) - 1

	for low <= high{
		median := (low + high) / 2
		
		if array_list[median] < key{
			low = median + 1
		}else {
			high = median - 1
		}
	}

	if low == len(array_list) || array_list[low] != key{
		return false
	}

	return true

}

//Time complexity is O(log n)
