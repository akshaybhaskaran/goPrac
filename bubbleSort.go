package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	// Create an empty size of 20 elements of type int
	slice := make([]int, 20)
	
	// Filling the slice with random integers
	rand.Seed(time.Now().UnixNano())
	for i:= 0; i < 20; i++{
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}

	fmt.Println("Unsorted Slice: \n", slice)
	bubbleSort(slice)
}

func bubbleSort(items [] int){
	
	n := len(items)

	//initially set swapped = true
	swapped := true

	for swapped{
		// set swapped to false
		swapped = false

		for i := 0; i < n-1; i++{
			if items[i] > items[i+1]{
				//swapping easily using Go's tuple assignment
				items[i], items[i+1] = items[i+1], items[i]
				// change swapped = true
				// because if loop ends and swapped is still false, 
				// algorithm will assume the list is fully sorted
				swapped = true

			}
		}
	}

	fmt.Println("Sorted Slice: \n", items)

}


//Worst Case Time Complexity:	O(n^2)
//Best Case Time Complexity:	O(n)