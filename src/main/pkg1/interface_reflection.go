/**
接口 + 反射
*/
package pkg1

import (
	"fmt"
	"math"
)

// 1、接口的实现
/* type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
} */

// 2、类型断言
type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

// 练习题 1：
/*
type Simpler interface {
	Get() int
	Put(int)
}

type Simple struct {
	i int
}

func (p *Simple) Get() int {
	return p.i
}

func (p *Simple) Put(u int) {
	p.i = u
}

func fI(it Simpler) int {
	it.Put(5)
	return it.Get()
}
*/
func init() {
	// 1、 接口的实现

	/* 	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	   	q := &Square{5}      // Area() of Square needs a pointer
	   	// shapes := []Shaper{Shaper(r), Shaper(q)}
	   	// or shorter

	   			// 如果 Rectangle 和 Square 没有实现接口，那么这里会报错：
	   			// cannot use r (type Rectangle) as type Shaper in array or slice literal:
	   	    // 	Rectangle does not implement Shaper (missing Area method)
	   			// cannot use q (type *Square) as type Shaper in array or slice literal:
	   	    // 	Square does not implement Shaper (missing Area method)

	   	shapes := []Shaper{r, q}

	   	fmt.Println("Looping through shapes for area ...")
	   	for _, s := range shapes {
	   		fmt.Printf("Shape details: %#v\n", s)
	   		fmt.Println("Area of this shape is: ", s.Area())
			 }
	*/
	// 2、类型断言
	/* var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1 //接口值 areaInf 的类型 动态变成 Shaper
	// Is Square the type of areaIntf?
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	areaIntf = new(Circle) //接口值 areaInf 的类型 动态变成 Circle
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Printf("areaIntf does not contain a variable of type Circle---%T\n", u)
	} */

	// 3、 type-switch 判断 接口变量的类型
/* 	switch areaIntf.(type) {
	case *Circle:
		fmt.Printf("type is Circle\n")
	default:
		fmt.Println("没有匹配项")
	} */
	


	// 练习题 1：
	/* var s Simple
	fmt.Println(fI(&s))
	*/
}
