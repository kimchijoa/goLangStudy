package customstruct

import (
	"container/list"
	"fmt"
	"reflect"
)

// 전달받은 인자 중 숫자만 리스트형태로 반환한다.
func MakeIntList(input ...interface{}) *list.List {
	sum := 0
	intList := list.New()
	fmt.Printf("nums 타입 : %T\n", input)

	for _, v := range input {
		fmt.Println("value 타입 : ", reflect.TypeOf(v))
		if reflect.TypeOf(v) == reflect.TypeOf(sum) {
			fmt.Println("add value : ", v)
			intList.PushBack(v)
		}
	}

	return intList
}

// 전달받은 인자 중 문자만 리스트형태로 반환한다.
func MakeStringList(input ...interface{}) *list.List {
	sum := ""
	stringList := list.New()
	fmt.Printf("nums 타입 : %T\n", input)

	for _, v := range input {
		fmt.Println("value 타입 : ", reflect.TypeOf(v))
		if reflect.TypeOf(v) == reflect.TypeOf(sum) {
			fmt.Println("add value : ", v)
			stringList.PushBack(v)
		}
	}

	return stringList
}
