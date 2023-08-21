package main

import "fmt"

// Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
// 接口可以让我们将不同的类型绑定到一组公共的方法上，从而实现多态和灵活的设计。
// Go 语言中的接口是隐式实现的，也就是说，如果一个类型实现了一个接口定义的所有方法，那么它就自动地实现了该接口。因此，我们可以通过将接口作为参数来实现对不同类型的调用，从而实现多态。
// type Phone interface {
// 	call()
// }

// type NokiaPhone struct {
// }

// func (nokiaPhone NokiaPhone) call() {
// 	fmt.Println("I am Nokia, I can call you!")
// }

// type IPhone struct {
// }

// func (iPhone IPhone) call() {
// 	fmt.Println("I am iPhone, I can call you!")
// }

// func main() {
// 	var phone Phone

// 	phone = new(NokiaPhone)
// 	phone.call()

// 	phone = new(IPhone)
// 	phone.call()

// }

type Shape interface {
	area() float64
	area2() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) area2() float64 {
	return r.width * r.height * 2
}

type Circle2 struct {
	radius float64
}

func (c Circle2) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	var s Shape

	s = Rectangle{width: 10, height: 5}
	fmt.Printf("矩形面积: %f\n", s.area())
	fmt.Printf("矩形面积: %f\n", s.area2())

	c := Circle2{radius: 3}
	fmt.Printf("圆形面积: %f\n", c.area())
}
