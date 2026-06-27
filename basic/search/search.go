// 단순한 선형검색
package main

import (
	"fmt"
	"math/rand"
)

func nsearch(a []int, find int) int {
	for i := 0; i < len(a); i++ {
		if a[i] == find {
			return i
		}
	}
	return -1
}
func main() {
	var n int
	var list []int
	fmt.Println("how many elements create? ")
	if _, err := fmt.Scanf("%d", &n); err != nil {
		fmt.Printf("Error : %v", err)
		panic(err)
	}

	for i := 0; i < n; i++ {
		num := rand.Intn(99) + 1
		list = append(list, num)
	}
	fmt.Printf("What number you want to find ? ")
	fmt.Scanf("%d", &n)

}
