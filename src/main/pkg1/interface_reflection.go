/**
接口 + 反射
*/
package pkg1

import (
	"fmt"
	"math"
	"reflect"
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

// 4.3、通过反射修改 属性值
type T struct {
	A int
	B string
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

	// 4.1、反射
	/*
		var x float64 = 3.4
		fmt.Println("type:", reflect.TypeOf(x))
		v := reflect.ValueOf(x)
		fmt.Println("value:", v)
		fmt.Println("type:", v.Type())
		fmt.Println("kind:", v.Kind())
		fmt.Println("value:", v.Float())
		fmt.Println(v.Interface())
		fmt.Printf("value is %5.2e\n", v.Interface())
		y := v.Interface().(float64)
		fmt.Println(y)
	*/

	// 4.2、通过反射 修改或设置值
	/*
		var x float64 = 3.4
		v := reflect.ValueOf(x)
		// setting a value:
		// v.SetFloat(3.1415) // Error: will panic: reflect.Value.SetFloat using unaddressable value
		fmt.Println("settability of v:", v.CanSet())
		v = reflect.ValueOf(&x) // Note: take the address of x.
		fmt.Println("type of v:", v.Type())
		fmt.Println("settability of v:", v.CanSet())
		v = v.Elem()
		fmt.Println("The Elem of v is: ", v)
		fmt.Println("settability of v:", v.CanSet())
		v.SetFloat(3.1415) // this works!
		fmt.Println(v.Interface())
		fmt.Println(v)
	*/

	// 4.3、通过反射修改 属性值
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t)
	// fmt.Println("s is : ", s) // s is :  &{23 skidoo}
	s = s.Elem()
	typeOfT := s.Type()
	// fmt.Println("ValueOf 之后 的 Type：", typeOfT) // ValueOf 之后 的 Type： pkg1.T
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v---Field:%v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface(), f)
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	// fmt.Println("t is now", t)

	// 练习题 1：
	/*
		var s Simple
		fmt.Println(fI(&s))
	*/
}
