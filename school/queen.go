// 简单队列
package school

import "errors"

type Queen []interface{}

func (q *Queen) InQueen(v interface{}) {
	*q = append(*q, v)
}

func (q *Queen) DeQueen() (interface{}, error) {
	if len(*q) > 0 {
		ret := (*q)[0]
		*q = (*q)[1:]
		return ret, nil
	}
	return -1, errors.New("queen is empty")
}
