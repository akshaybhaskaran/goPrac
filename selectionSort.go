package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	// Create an empty size of 20 elements of type int
	elements := make([]int, 20)
	
	// Filling the slice with random integers
	rand.Seed(time.Now().UnixNano())
	for i:= 0; i < 20; i++{
		elements[i] = rand.Intn(999) - rand.Intn(999)
	}

	fmt.Println("Unsorted Slice: \n", elements)
	selectionSort(elements)
	fmt.Println("Sorted Slice: \n", elements)
	
}

func selectionSort(elem [] int) []int{

	len := len(elem)

	for i:=0; i<len; i++ {
		var minIndex = i
		for j:=i; j<len; j++{
			//starting at i, we search for the smallest item that exists between index i and last element. 
			if elem[j] < elem[minIndex]{
				//when this element has been found, its exchanged with the data found at index i
				minIndex = j
			}
		}
		//swapping the i'th element with minimum index
		elem[i], elem[minIndex] = elem[minIndex], elem[i]
		
	}
	

	return elem
}


//best complexity: O(n^2)
//worst complexity: O(n^2)