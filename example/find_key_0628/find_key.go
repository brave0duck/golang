// 배열a에서 key값을 찾아 그 인덱스를 배열에 차곡차곡 저장하고 찾는갯수를 리턴하는 함수 작성
// ex) a={1,9,2,9,4,6,7,9} -----> idx = {1,3,7} return 3;

package main

import (
	"fmt"
	"math/rand"
)

func find_key(a []int, key int) ([]int, int) {
	var answer []int
	var count int

	for i, v := range a {
		if v == key {
			answer = append(answer, i)
			count++
		}
	}
	return answer, count
}
func main() {
	var input int
	lst := make([]int, 30)

	for i := 0; i < 30; i++ {
		lst[i] = rand.Intn(20)
	}

	fmt.Println("1-30 랜덤배열에서 값을 찾아 그 인덱스를 배열로 반환합니다.\n 찾을값은 ? (1-20) : ")
	fmt.Scanf("%d\n", &input)

	find_list, list_size := find_key(lst, input)
	fmt.Printf("찾은 값은 : %v, 총 %d개를 찾았습니다\n", find_list, list_size)

	fmt.Printf("랜덤배열 : %v\n",lst)
}
