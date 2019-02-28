package main

import (
	"fmt"
)

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2 ; i++ {
		if input[i] != input[len(input)-i-1]{
			return false
		}
	}
	return true
		
}

func main(){

	input := "Akshay"
	fmt.Println("---Palindrome Check---")
	fmt.Println("Input: ", input)
	if isPalindrome(input) == true{
		fmt.Println(input, "is a Palindrome")
	} else{
		fmt.Println(input, "is not a Palindrome")
	}
		
}