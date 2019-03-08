package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	// Create an empty size of 20 elements of type int
	arr := make([]int, 20)
	
	// Filling the slice with random integers
	rand.Seed(time.Now().UnixNano())
	for i:= 0; i < 20; i++{
		arr[i] = rand.Intn(999) - rand.Intn(999)
	}

	fmt.Println("Unsorted Slice: \n", arr)
	insertionSort(arr)
	fmt.Println("Sorted Slice :\n", arr)
	
}

func insertionSort(arr [] int){
	var n = len(arr)
	for i := 1; i < n; i ++ {
		j := i 
		for j > 0{
			if arr[j-1] > arr[j]{
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j = j - 1
		}

	}
}





/*Pseudocode:
for i: 1 to length(array) - 1
	j = i
	while j > 0 and array[j-1] > array[j]
		swap array[j] and array [j-1]
		j = j - 1

*/


//Best case: O(n)
//Worst case: O(n^2)