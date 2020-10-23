package main

import (
	"fmt"
	"net"
	"strconv"
)

type data struct {
	name string
}

func main() {

	fmt.Println(strconv.FormatInt(128, 16))
	fmt.Println(strconv.ParseInt("o80", 0, 64))

}

func fun1() {
	i, j := 0, 0
	if true {
		j, k := 1, 1
		fmt.Println(j, k)
	}
	fmt.Println(i, j)
}

type Student struct {
	name string
	age  uint8
}

func equalsIgnoreCase(a, b string) (bool, string) {
	m, n := len(a), len(b)
	for i := 0; i < m || i < n; i++ {
		if i == m {
			return false, b
		}
		if i == n {
			return false, a
		}
		x, y := offset(a[i]), offset(b[i])
		if x > y {
			return false, a
		} else if x < y {
			return false, b
		}
	}
	return true, ""
}

func offset(c byte) byte {
	switch {
	case 'a' <= c && c <= 'z':
		return c - 'a'
	case 'A' <= c && c <= 'Z':
		return c - 'A'
	default:
		return c
	}
}

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
)

func (b ByteSize) String() string {
	switch {
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

type intSlice []int

func (p *intSlice) append(data ...int) {
	/*slice := *p
	slice = append(slice, data...)
	*p = slice*/
	*p = append(*p, data...)
}

func (p *intSlice) write(buf []int) (int, error) {
	slice := *p
	slice = append(slice, buf...)
	*p = slice
	return len(buf), nil
}

func exampleTcp() {
	listener, e := net.Listen("tcp", "127.0.0.1:8092")
	if e != nil {
		fmt.Println("error=", e)
		return
	}
	defer listener.Close()
	conn, e := listener.Accept()
	if e != nil {
		fmt.Println("error=", e)
		return
	}
	defer conn.Close()
	buff := make([]byte, 1024)
	n, e := conn.Read(buff)
	if e != nil {
		fmt.Println("error=", e)
		return
	}
	fmt.Println("receive data:", string(buff[0:n]))
	conn.Write([]byte("well done"))
	str := "aaa"
	u := str[0]
	print(u)
}

type queen []int

func (q *queen) enQueen(a int) {
	*q = append(*q, a)
}

func (q *queen) deQueen() {
	slice := *q
	slice = slice[:len(slice)-1]
	*q = slice
}
