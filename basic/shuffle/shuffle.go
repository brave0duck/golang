// 배열a의 모든 요소의 순서를 뒤섞는 shuffle함수를 작성. n은 요소의 개수
package main

import (
	"fmt"
	"math/rand"
)

func shuffle(sample []int) {
	for i := 0; i < 100; i++ {
		pick := rand.Intn(len(sample)-1)
		sample[0], sample[pick] = sample[pick], sample[0]
	}
}
func main() {
	sample := make([]int,20)
	for i := 0; i < 20; i++ {
		sample = append(sample, rand.Intn(99)+1)
		fmt.Printf("%d ", sample[i])
	}

	fmt.Printf("\nnow. shuffle it\n")
	shuffle(sample)

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%d ", sample[i])
	}
	fmt.Println("")
}
