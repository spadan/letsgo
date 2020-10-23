// 模板模式
package pattern

import "fmt"

// 自定义扩展接口，模板中定义的自定义扩展
type iCustom interface {
	step1()
	step2()
}

// 封装整体逻辑
type handler struct {
	// 提供自定义扩展逻辑
	sub iCustom
}

// 整体通用流程
func (h *handler) handle() {
	fmt.Println("do some common things")
	h.sub.step1()
	fmt.Println("do some common things")
	h.sub.step2()
}

// 自定义实现1
type myCustom1 struct{}

func (s *myCustom1) step1() {
	fmt.Println("myCustom1 exec step 1")
}

func (s *myCustom1) step2() {
	fmt.Println("myCustom1 exec step 2")
}

// 自定义实现2
type myCustom2 struct{}

func (s *myCustom2) step1() {
	fmt.Println("myCustom2 exec step 1")
}

func (s *myCustom2) step2() {
	fmt.Println("myCustom2 exec step 2")
}

func templateTest() {
	c1 := &myCustom1{}
	h1 := handler{c1}
	h1.handle()
	fmt.Println("-----------------------------")
	c2 := &myCustom2{}
	h2 := handler{c2}
	h2.handle()
}
