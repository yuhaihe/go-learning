package shape

type Rect struct {
	W int
	H int
}

func (r Rect) Area() int {
	panic("TODO")
}

type Address struct {
	City string
}

type User struct {
	Name string
	Address
}

func (u User) Label() string {
	panic("TODO")
}
