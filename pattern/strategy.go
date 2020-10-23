// 策略模式
package pattern

import "fmt"

// 策略接口
type iStrategy interface {
	exec()
}

// 具体策略1
type fifo struct{}

func (f *fifo) exec() {
	fmt.Println("exec fifi iStrategy")
}

// 具体策略2
type lru struct{}

func (l *lru) exec() {
	fmt.Println("exec lru iStrategy")
}

// 具体策略3
type lfu struct{}

func (l *lfu) exec() {
	fmt.Println("exec lfu iStrategy")
}

// 持有策略的客户端
type client1 struct {
	s iStrategy
}

func (c client1) do() {
	c.s.exec()
}

func strategyTest() {
	s1, s2, s3 := &fifo{}, &lru{}, &lfu{}
	c := client1{s1}
	c.do()
	fmt.Println("----------------------")
	c.s = s2
	c.do()
	fmt.Println("----------------------")
	c.s = s3
	c.do()
}
