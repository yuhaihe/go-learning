package speak

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (Dog) Speak() string {
	panic("TODO")
}

func Say(s Speaker) string {
	panic("TODO")
}
