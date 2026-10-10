package speak

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (Dog) Speak() string {
	return "woof"
}

func Say(s Speaker) string {
	if s == nil {
		return ""
	}
	return s.Speak()
}
