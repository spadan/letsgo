package pattern

import (
	"fmt"
	"sync"
	"testing"
)

func TestFactory_GetGun(t *testing.T) {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")
	ak47.shot()
	musket.shot()
}

func TestSingle_getInstance(t *testing.T) {
	var group sync.WaitGroup
	for i := 0; i < 30; i++ {
		group.Add(1)
		go func() {
			getInstance()
			group.Done()
		}()
	}
	group.Wait()
}

func TestAdaper(t *testing.T) {
	client := &client{}
	mac := &mac{}
	client.insertLightningConnector2Computer(mac)

	adapter := &adapter{&windows{}}
	client.insertLightningConnector2Computer(adapter)
}

func TestDecorator(t *testing.T) {
	p := &veggeMania{}
	fmt.Println(p.getPrice())
	tp := tomatoTopping{p}
	fmt.Println(tp.getPrice())
	cp := cheeseTopping{p}
	fmt.Println(cp.getPrice())
	ctp := cheeseTopping{&tomatoTopping{p}}
	fmt.Println(ctp.getPrice())
}

func TestChain(t *testing.T) {
	r := &reception{}
	d := &doctor{}
	r.setNext(d)
	m := &medical{}
	d.setNext(m)
	c := &cashier{}
	m.setNext(c)
	r.execute()
}

func TestTemplate(t *testing.T) {
	templateTest()
}

func TestStrategy(t *testing.T) {
	strategyTest()
}