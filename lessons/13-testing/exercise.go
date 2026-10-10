package fake

// Store 按 id 查名字。
type Store interface {
	Name(id int) (string, error)
}

// Greet 查到名字后返回 "hi " 加名字。查失败时把错误原样返回。
func Greet(s Store, id int) (string, error) {
	panic("TODO")
}
