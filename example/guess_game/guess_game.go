// 1. usr input 1-100
// 2. while usr guess number
// 3. chance count
// 4. program exit if correct number
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var usr_number int = 0
	var ncount int = 5

	fmt.Println("-----------------------------------")
	fmt.Println("guess number game")
	fmt.Println("must correct within 5 chance")
	fmt.Println("----------------------------")

	com_number := rand.Intn(99) + 1

	for ncount > 0 {

		fmt.Printf("Guess number 1-100 (%d remain)\n", ncount)
		if _, err := fmt.Scanf("%d", &usr_number); err != nil {
			fmt.Printf("only number can input : %v\n", err)
			continue
		}

		if usr_number > com_number {
			fmt.Println("no .. it's high")
		} else if usr_number < com_number {
			fmt.Println("no .. it's low")
		} else {
			fmt.Println("Amazing! you Correct")
			fmt.Printf("You correct %dth try\n", 6-ncount)
			break
		}
		ncount--

	}
	if ncount == 0 {
		fmt.Printf("Unfortune guy, Correct number was [%d]\nProgram exit...", com_number)
	}
}
