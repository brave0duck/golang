// 1. 정수 리스트가 주어졌을 때, 벡터를 이용하여 이 리스트의 중간값 (median, 정렬했을 때 가장 가운데 위치한 값),
// 그리고 최빈값 (mode, 가장 많이 발생한 값; 해시맵이 여기서 도움이 될 것입니다) 을 반환해 보세요.

// 2. 문자열을 피그 라틴 (pig Latin) 으로 변경해 보세요.
// 각 단어의 첫 번째 자음은 단어의 끝으로 이동하고 ‘ay’를 붙이므로, ‘first’는 ‘irst-fay’가 됩니다.
// 모음으로 시작하는 단어는 대신 끝에 ‘hay’를 붙입니다.
// (‘apple’은 ‘apple-hay’가 됩니다.) UTF-8 인코딩에 대한 세부 사항을 명심하세요!

// 3. 해시맵과 벡터를 이용하여 사용자가 회사 부서의 직원 이름을 추가할 수 있도록 하는 텍스트 인터페이스를 만들어 보세요.
// 예를 들어 ‘Add Sally to Engineering’이나 ‘Add Amir to Sales’ 같은 식으로요.
// 그 후 사용자가 모든 사람에 대해 알파벳 순으로 정렬된 목록이나 부서별 모든 사람에 대한 목록을 조회할 수 있도록 해보세요.

// 1에서 100사이의 정수리스트 20개 생성.
package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func median(p []int) int {
	size := len(p)
	mid := int(size / 2)
	if mid%2 == 0 {
		return int((p[mid-1] + p[mid+1]) / 2)
	} else {
		return p[mid]
	}
}
func mode(p []int) (int, int) {
	var find_mode int = 0

	m := make(map[int]int, 20)
	result := make(map[int]int)

	for _, value := range p {
		m[value] += 1
	}
	// 2,3,3,3,5 = 2:1, 3:3, 5:1
	for index, value := range m {
		if value > find_mode {
			result[index] = value
		}
	}
	return count_n, mode_n
}

func main() {
	var sample = make([]int, 20)
	for i, _ := range sample {
		sample[i] = rand.Intn(49) + 1
	}
	slices.Sort(sample)

	median_n := median(sample)
	fmt.Printf("\nmedian : %d\n", median_n)

	count_n, mode_n := mode(sample)
	fmt.Printf("mode : %d, count : %d\n", count_n, mode_n)
}
