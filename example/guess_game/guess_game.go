// guess number game
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
	fmt.Println("You have to get it right within 5 tries.")
	fmt.Println("-----------------------------------")

	com_number := rand.Intn(99) + 1 // random 1-100

	for ncount > 0 {

		fmt.Printf("Guess number 1-100 (%d remain)\n", ncount)
		if _, err := fmt.Scanf("%d", &usr_number); err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		if 0 > usr_number || usr_number > 100{
			fmt.Println("== only 1-100 number. retry")
			continue
		}
		if usr_number > com_number {
			fmt.Println("---> it's high")
		} else if usr_number < com_number {
			fmt.Println("---> it's low")
		} else {
			fmt.Printf("Amazing! you Correct!\n")
			fmt.Printf("You correct %dth try\n", 6-ncount)
			break
		}
		ncount--

	}
	if ncount == 0 {
		fmt.Printf("Unfortunately, You couldn't guess it. The correct answer is [%d]\n", com_number)
	}
	fmt.Printf("Program exit...\n")
}
