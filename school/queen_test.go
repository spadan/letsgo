package school

import (
	"fmt"
	"testing"
)

func TestQueen_InQueen(t *testing.T) {
	q := Queen{}
	q.InQueen(1)
	q.InQueen(2)
	q.InQueen(3)
	for {
		v, err := q.DeQueen()
		fmt.Println(v, err)
		if err != nil {
			break
		}
	}
}

func TestStack_Push(t *testing.T) {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	for len(s) > 0 {
		fmt.Println(s.Pop())
	}
}
