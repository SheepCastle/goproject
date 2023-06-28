package main

import "fmt"

type Data struct { // Data 타입 구조체
	value int
	data  [200]int
}

func ChangeData(arg Data) {
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data
	ChangeData(data)
	fmt.Printf("value = %d\n", data.value)         // 인수로 data를 넣습니다.
	fmt.Printf("data[100] = %d\n", data.data[100]) // data의 두 필드 출력
}
