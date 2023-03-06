package main

import (
	"DataStructure/customstruct"
	"fmt"
)

func main() {
	fmt.Println("Start Make IntList")
	v1 := customstruct.MakeIntList(1, 2, "3", 4, "5")

	for e := v1.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}

	v2 := customstruct.MakeStringList(1, 2, "3", 4, "5")

	for e := v2.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
}
