package main

import "fmt"

func linearSearch(list[]int, key int) bool{
	for _, item := range list{
		if item == key{
			return true
		}
	}
	return false
}

func main(){
	items := []int {22,45,67,54,26,89,54,67,32,76,44}
	key := 10
	fmt.Println(linearSearch(items, key))

}



//Worst cast complexity is O(n). 
//Time increases as the number of elements increases
