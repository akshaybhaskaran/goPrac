package main

import (
	"fmt"
	//"strings"
)

func countVowels(input string) int{

	count := 0

	for _,a := range input{
		c := string(a)
		fmt.Println(c)
		if c == "a" || c == "e" || c == "i" || c == "o" || c == "u"{
			count = count + 1
		}
	}
	
	return count
}

func main(){

	input := "this is a sample text"
	fmt.Println("No. of vowels: ", countVowels(input))
}