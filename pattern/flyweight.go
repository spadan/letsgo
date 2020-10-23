// 享元模式
package pattern

type iDress interface {
	getColor() string
}

type dressType uint8

const (
	terrorist dressType = iota
	police
)

type dress struct {
	color string
}

func (d dress) getColor() string {
	return d.color
}

type terroristDress struct {
	dress
}

type policeDress struct {
	dress
}

type dressFactory struct {
	cache map[dressType]iDress
}

func (f *dressFactory) getDressByType(dressType dressType) iDress {
	if v, ok := f.cache[dressType]; ok {
		return v
	}
	switch dressType {
	case terrorist:
		f.cache[dressType] = terroristDress{dress{"red"}}
	case police:
		f.cache[dressType] = policeDress{dress{"green"}}
	default:
		panic("unknown dress type")
	}
	return f.cache[dressType]
}
