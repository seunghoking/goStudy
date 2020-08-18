package main

import (
	"fmt"
	// 표준 입력을 받기위해 필요
	"os"
	// 입력으로 부터 한줄을 입력받기 위해서 read를 해야되서
	"bufio"
	// 스트링 찌꺼기 없애는
	"strings"
	// 문자열을 숫자로 바꿔주는 기능을 가지고 있음
	"strconv"
)

func main() {
	fmt.Println("숫자를 입력하세요")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n1, _ := strconv.Atoi(line)

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n2, _ := strconv.Atoi(line)

	fmt.Printf("입력하신 숫자는 %d %d 입니다.", n1, n2)

	fmt.Println("연산자를 입력하세요")

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	if line == "+" {
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	} else if line == "-" {
		fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	} else if line == "*" {
		fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	} else if line == "/" {
		fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	} else {
		fmt.Println("잘못 입력하셨습니다.")
	}
}
