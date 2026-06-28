package main

import (
	"fmt"
	"sort"
	"strings"
)

func addEmployee(employee map[int][]string, deptID int, name string) {
	if name == "" {
		fmt.Println("이름을 입력하세요.")
		return
	}

	employee[deptID] = append(employee[deptID], name)
	sort.Strings(employee[deptID])
	fmt.Printf("%s를 부서 %d에 추가했습니다.\n", name, deptID)
}

func removeEmployee(employee map[int][]string, deptID int, name string) {
	names := employee[deptID]
	newNames := names[:0]
	removed := false

	for _, n := range names {
		if n == name {
			removed = true
			continue
		}
		newNames = append(newNames, n)
	}

	if !removed {
		fmt.Printf("%s는 부서 %d에 없습니다.\n", name, deptID)
		return
	}

	if len(newNames) == 0 {
		delete(employee, deptID)
	} else {
		employee[deptID] = newNames
	}

	fmt.Printf("%s를 부서 %d에서 삭제했습니다.\n", name, deptID)
}

func printDepartment(employee map[int][]string, deptID int) {
	names, ok := employee[deptID]	//이 표현식을 주목
	if !ok || len(names) == 0 {
		fmt.Printf("부서 %d에는 직원이 없습니다.\n", deptID)
		return
	}

	fmt.Printf("부서 %d 직원 목록:\n", deptID)
	for _, name := range names {
		fmt.Println("-", name)
	}
}

func printAll(employee map[int][]string) {
	if len(employee) == 0 {
		fmt.Println("등록된 직원이 없습니다.")
		return
	}

	deptIDs := make([]int, 0, len(employee))
	for deptID := range employee {
		deptIDs = append(deptIDs, deptID)
	}
	sort.Ints(deptIDs)

	fmt.Println("전체 직원 목록:")
	for _, deptID := range deptIDs {
		names := employee[deptID]
		sort.Strings(names)
		fmt.Printf("부서 %d: %s\n", deptID, strings.Join(names, ", "))
	}
}

func menu() {
	fmt.Println("1. 추가")
	fmt.Println("2. 삭제")
	fmt.Println("3. 부서 조회")
	fmt.Println("4. 전체 조회")
	fmt.Println("5. 종료")
}

func main() {
	employee := make(map[int][]string)

	for {
		menu()
		fmt.Print("선택: ")

		var input string
		fmt.Scanln(&input)

		switch input {
		case "1":
			fmt.Print("부서 번호: ")
			var deptID int
			fmt.Scanln(&deptID)

			fmt.Print("이름: ")
			var name string
			fmt.Scanln(&name)

			addEmployee(employee, deptID, name)
		case "2":
			fmt.Print("부서 번호: ")
			var deptID int
			fmt.Scanln(&deptID)

			fmt.Print("이름: ")
			var name string
			fmt.Scanln(&name)

			removeEmployee(employee, deptID, name)
		case "3":
			fmt.Print("부서 번호: ")
			var deptID int
			fmt.Scanln(&deptID)
			printDepartment(employee, deptID)
		case "4":
			printAll(employee)
		case "5":
			fmt.Println("종료합니다.")
			return
		default:
			fmt.Println("잘못된 입력입니다.")
		}
	}
}
