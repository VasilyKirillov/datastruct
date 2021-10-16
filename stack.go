package datastruct

import (
	"fmt"
)

type Stack []interface{}

func New(elems ...interface{}) Stack {
	if elems == nil {
		var s []interface{}
		return s
	}
	s := make([]interface{}, len(elems))
	copy(s, elems)
	return s
}

func (s *Stack) Pop() interface{} {
	size := len(*s) - 1
	elem := (*s)[size]
	(*s)[size] = nil
	*s = (*s)[:size]
	return elem
}

func (s *Stack) Push(elem interface{}) {
	*s = append(*s, elem)
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s Stack) String() string {
	var res string = "["
	for i, el := range s {
		res += fmt.Sprint(el)
		if i < len(s)-1 {
			res += ","
		}
	}
	return res + "]"
}
