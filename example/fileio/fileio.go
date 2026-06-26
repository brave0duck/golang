// 간단한 파일입출력 예제.
// file_open파일의 플래그 - CREATE, RDWR, APPEND (없으면 만들고, 읽고쓸수있고, 추가해라)
package main

import (
	"bufio"
	"fmt"
	"os"
)

func file_open(file_name string) *os.File {
	f, err := os.OpenFile(file_name, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.FileMode(0644))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed open file : %v", err)
		return nil
	}
	file_info := f.Fd()
	fmt.Println("파일 디스크립터 : ", file_info)
	return f
}
func file_read(file_name string) string {
	f, err := os.OpenFile(file_name, os.O_RDONLY, os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer f.Close()

	f_stat, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	data := make([]byte, f_stat.Size())
	byte_size, err := f.Read(data)
	if err != nil {
		fmt.Println("error occurred : ", err)
		return ""
	} else {
		fmt.Println(byte_size, "바이트를 읽어왔습니다")
	}

	return string(data)
}
func file_write(f *os.File, text string) bool {

	_, err := f.Write([]byte(text))
	if err != nil {
		fmt.Println("Error occurred : ", err)
		return false
	}
	return true
}
func main() {
	var file_name string
	var mytext string

	fmt.Println("--------------------------------")
	fmt.Println("간단한 파일 입출력 연습")
	fmt.Println("--------------------------------")
	fmt.Print("파일이름을 뭘로 할까요? ")

	fmt.Scanf("%s\n", &file_name)

	f := file_open(file_name)
	if f == nil {
		fmt.Println("==> 파일을 여는데 실패했습니다")
		return
	} else {
		fmt.Println("==> 파일을 여는데 성공했습니다")
	}
	defer f.Close()

	fmt.Print("파일에 저장하고 싶은 문구를 입력하세요 : ")
	reader := bufio.NewReader(os.Stdin)
	mytext, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("키보드에서 문자열을 읽어오는데 실패했습니다")
	}

	if file_write(f, mytext) {
		fmt.Println("파일에 저장하였습니다.")
	} else {
		fmt.Println("파일 저장에 실패했습니다")
	}

	fmt.Printf("[%s] 파일의 내용을 출력합니다\n", file_name)
	print_text := file_read(file_name)

	fmt.Printf("파일 내용 : %s\n", print_text)
	fmt.Println("프로그램을 종료합니다...")
}
