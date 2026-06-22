package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {

	fmt.Printf("%v %v %v\n",
		color.RedString("Hello"),
		color.YellowString("fucking amazing"),
		color.GreenString("World!"))
}
