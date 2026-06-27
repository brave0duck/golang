// 1부터 n까지의 합 출력
package main

import (
	"fmt"
)

func main() {
	var usr int
	var sum_result int

	fmt.Println("1 to n sum : (0이상)")
	if _, err := fmt.Scanf("%d", &usr); err != nil {
		fmt.Printf("Input only number : %v\n", err)
		return
	} else {
		if usr < 0 {
			fmt.Printf("The number must be 0 or greater\n")
			return
		}
	}

	for i := 1; i <= usr; i++ {
		sum_result += i
	}
	fmt.Println("sum total : ", sum_result)
}
