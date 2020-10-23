// 简单工厂模式
package pattern

import (
	"errors"
	"fmt"
)

type iGun interface {
	shot()
}

type gun struct {
	name string
}

func (g *gun) shot() {
	fmt.Printf("%s shot\n", g.name)
}

type ak47 struct {
	gun
}

type musket struct {
	gun
}

func newAk47() iGun {
	return &ak47{gun{"ak47"}}
}

func newMusket() iGun {
	return &musket{gun{"musket"}}
}

func getGun(gunType string) (iGun, error) {
	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "musket":
		return newMusket(), nil
	default:
		return nil, errors.New("unknown gun type")
	}
}
