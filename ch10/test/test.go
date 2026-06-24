// 깊은 복사를 하려면 copy함수를 써라
// 순회를 도려면 i := range <슬라이스>를 써라

package main

import "fmt"

func main() {
	origin := []int{1, 2, 3}
	cp_origin := make([]int, len(origin))
	copy(cp_origin, origin)	//deep copy

	c1 := origin	// allocation is shadow copy
	c2 := origin[:] // same too

	origin[0] = 5

	fmt.Println(origin)
	fmt.Println(cp_origin)
	fmt.Println(c1)
	fmt.Println(c2)

}
