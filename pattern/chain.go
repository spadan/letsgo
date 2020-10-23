// 责任链模式
package pattern

import "fmt"

// 部门接口
type iDepartment interface {
	// 部门处理逻辑
	execute()
	// 下一个处理部门
	setNext(iDepartment)
}

type department struct {
	next iDepartment
}

func (d *department) setNext(dep iDepartment) {
	d.next = dep
}

type reception struct {
	department
}

func (r *reception) execute() {
	fmt.Println("前台接待")
	r.next.execute()
}

type doctor struct {
	department
}

func (r *doctor) execute() {
	fmt.Println("医生看病")
	r.next.execute()
}

type medical struct {
	department
}

func (r *medical) execute() {
	fmt.Println("药房发药")
	r.next.execute()
}

type cashier struct {
	department
}

func (r *cashier) execute() {
	fmt.Println("收银")
	r.next.execute()
}
