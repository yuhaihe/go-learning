package shape

type Rect struct {
	W int
	H int
}

func (r Rect) Area() int {
	return r.W * r.H
}

type Address struct {
	City string
}

type User struct {
	Name string
	Address
}

func (u User) Label() string {
	return u.Name + "@" + u.City
}
