package main

import (
	"fmt"
	"math/rand"
	"time"
)
func guess(input, secret_num int){
	
	if input < secret_num{
		fmt.Println("Your guess is low!")
	}else if input > secret_num{
		fmt.Println("Your guess is high")
	}else if input >= 50{
		fmt.Print("Please guess within 50")
	}else{
		fmt.Println("Nope! You're not there!")
	}
}

func main(){

	fmt.Println("--<GUESS THE NUMBER GAME>--")
	
	//Generating a random number. A seed value produces varying sequences
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	secret_num := random.Intn(50)

	fmt.Println("Secret num: ",  secret_num)

	fmt.Println("You have 5 chances to guess the number!")
	var input int
	for i := 0; i <= 5; i++{
		fmt.Print("Guess #",i,": ")
		fmt.Scanf("%d", &input)
		if input == secret_num{
			fmt.Println("You guess it right!")
			break
		}else{
			guess(input, secret_num)
		}
	}
}



