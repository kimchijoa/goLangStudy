package customstruct

import "container/list"

func GetList() *list.List {
	v := list.New()
	return v
}
