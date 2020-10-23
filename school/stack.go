// 简单栈
package school

import "errors"

type Stack []interface{}

func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (interface{}, error) {
	n := len(*s)
	if n > 0 {
		ret := (*s)[n-1]
		*s = (*s)[:n-1]
		return ret, nil
	}
	return -1, errors.New("stack is empty")
}
