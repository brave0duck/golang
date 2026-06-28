/*수많은 마라톤 선수들이 마라톤에 참여하였습니다.
단 한 명의 선수를 제외하고는 모든 선수가 마라톤을 완주하였습니다.
마라톤에 참여한 선수들의 이름이 담긴 배열 participant와 완주한 선수들의 이름이 담긴 배열 completion이 주어질 때,
완주하지 못한 선수의 이름을 return 하도록 solution 함수를 작성해주세요.

["leo", "kiki", "eden"]	["eden", "kiki"]	"leo"
["marina", "josipa", "nikola", "vinko", "filipa"]	["josipa", "filipa", "marina", "nikola"]	"vinko"
*/
// 탈락자를 가려내는 프로그램

package main

import (
	"fmt"
	"slices"
)

func solution(total []string, race []string) string {
	for _, v := range total {
		if isin := slices.Contains(race, v); isin != true {
			return v
		}
	}
	return ""
}
func main() {
	// var total []string = []string{"leo", "kiki", "eden"}
	// var race []string = []string{"eden", "kiki"}

	var total []string = []string{"marina", "josipa", "nikola", "vinko", "filipa"}
	var race []string = []string{"josipa", "filipa", "marina", "nikola"}

	a := solution(total, race)
	if a == "" {
		fmt.Println("No one cut off")
	} else {
		fmt.Println("cut off :", a)
	}
}
