// guess number game
package main

import (
	"fmt"
	"math/rand"

	"github.com/fatih/color"
)

func main() {
	var chance int = 5
	var usr_guess int = 0
	com := rand.Intn(100) + 1

	fmt.Println(color.BlueString("==========================="))
	fmt.Println(color.YellowString("welcome to guess number game"))
	fmt.Println(color.BlueString("==========================="))

	for i := 1; i <= chance; i++ {
		fmt.Println(color.YellowString("Try to guees number (1-100) : "))
		fmt.Scanf("%d", &usr_guess)

		if 0 > usr_guess || usr_guess > 100 {
			fmt.Println(color.RedString("wrong input.. only 1-100 number"))
			continue
		} else {
			switch {
			case usr_guess > com:
				fmt.Println(color.RedString("high"))
			case usr_guess < com:
				fmt.Println(color.BlueString("low"))
			case usr_guess == com:
				fmt.Println(color.GreenString("Great! you win. it's correct"))
				goto SUCCESS_EXIT
			}

		}
		fmt.Println(color.YellowString("remain chance %d\n", chance-i))

	}

	fmt.Print("computer random number was ")
	fmt.Print(color.RedString(fmt.Sprintf("%d", com)))
	fmt.Print("  unfortunely... program exit\n")

SUCCESS_EXIT:
}
