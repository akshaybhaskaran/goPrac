package main

import "fmt"
import "math/rand"
import "time"

func main(){

	// Create an empty size of 20 elements of type int
	arr := make([]int, 20)
	
	// Filling the slice with random integers
	rand.Seed(time.Now().UnixNano())
	for i:= 0; i < 20; i++{
		arr[i] = rand.Intn(999) - rand.Intn(999)
	}

	fmt.Println("Unsorted Slice: \n", arr)
	quickSort(arr)
	fmt.Println("Sorted Slice :\n", arr)
	
}

func quickSort(arr [] int) []int{
	if len(arr) < 2{
		return arr
	}

	//setting left and right pointers
	left, right := 0, len(arr)-1

	//selecting random pivot value
	pivot := rand.Int() % len(arr)

	//swapping pivot and right pointers
	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i,_ := range arr{
		if arr[i] < arr[right]{
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	
	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr 
}


//worstcase: O(n^2)
//best case: O(nlogn)