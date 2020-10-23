package school


type Integer int

func (a Integer)Equals(b Integer) bool{
	return a== b
}

func (a *Integer) Increase(i Integer) {
	*a=*a+ i
}