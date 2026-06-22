// 1. 사용자에게 정수형 숫자를 입력받는다 (클로저 사용)
// 2. 숫자가 0이면 종료
// 3. 2-9범위면 출력하고 아니면 2-9단을 모두 출력한다
package main

import (
	"fmt"
)

func gugu(n int) {
	for i := 1; i < 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
	fmt.Println("")
}
func main() {
	var usr int
	fmt.Println("----------------------------")
	fmt.Println(" Gu-Gu-Dan Program")
	fmt.Println("----------------------------")

	for {
		fmt.Println("What do you want print gu-gu-dan number ? (exit:0) ")

		fmt.Scanf("%d", &usr)

		if usr != 0 {
			if 2 <= usr && usr <= 9 {
				gugu(usr)
			} else {
				for i := 2; i <= 9; i++ {
					gugu(i)
				}
			}
		} else {
			fmt.Println("Program exit...")
			break
		}

	}

}
