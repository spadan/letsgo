// 装饰模式
package pattern

// 披萨接口
type pizza interface {
	getPrice() int
}

// 普通披萨
type veggeMania struct{}

func (p *veggeMania) getPrice() int {
	return 15
}

// 加奶油
type cheeseTopping struct {
	pizza pizza
}

func (p *cheeseTopping) getPrice() int {
	return p.pizza.getPrice() + 20
}

// 加土豆
type tomatoTopping struct {
	pizza pizza
}

func (p *tomatoTopping) getPrice() int {
	return p.pizza.getPrice() + 10
}
