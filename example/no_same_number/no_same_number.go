// arr = [1, 1, 3, 3, 0, 1, 1] 이면 [1, 3, 0, 1] 을 return 합니다.
// arr = [4, 4, 4, 3, 3] 이면 [4, 3] 을 return 합니다.
// 배열 arr에서 연속적으로 나타나는 숫자는 제거하고 남은 수들을 return 하는 solution 함수를 완성해 주세요.
package main

import (
	"fmt"
)

func solution(a []int) []int {
	var answer []int
	var compare int = -1

	for _, v := range a {
		if compare != v {
			answer = append(answer, v)
			compare = v
		}
	}
	return answer
}
func main() {
	t := []int{4, 4, 4, 3, 3}
	a := solution(t)
	fmt.Println(a)
}
