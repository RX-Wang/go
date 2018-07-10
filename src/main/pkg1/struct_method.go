/**
结构体 和 方法
*/

package pkg1

import (
	"fmt"
	// "unsafe"
	"reflect"
	"strconv"
	"time"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

// 2、结构体的 Tag
type TagType struct { // tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

// 3、结构体的匿名字段 + 内嵌结构体
type innerS struct {
	in1 int
	in2 int
	b   int
}

type outerS struct {
	b      int
	c      float32
	int    // 结构体的匿名字段，名字就是类型名。
	innerS // 内嵌结构体
}

// 4、命名冲突（结构体内嵌时）-- 重载
type aStruct struct{ a, b int }
type bStruct struct {
	aStruct
	a float32
}

// 5.1、 结构体的方法  以 innerS 为接受者-- ins 的作用相当于 this
func (ins *innerS) sum() int {
	return ins.in1 + ins.in2
}

// 5.2、 非结构体方法
type intVector []int

func (v intVector) sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

// 6、给非本地类型 定义方法  -- 先给类型起一个别名，然后赋方法
type myTime struct {
	aField    int
	time.Time // 匿名字段
}

// 重载String 方法， 遍历 结构体
func (t myTime) String() string {
	var s = "{ "
	v := reflect.ValueOf(t)
	count := v.NumField()
	for i := 0; i < count; i++ {
		fmt.Println("--------", v.Field(i), v.Field(i).Kind())
		switch v.Field(i).Kind() {
		case reflect.Int:
			s = s + strconv.FormatInt(v.Field(i).Int(), 10) + ", "
		case reflect.Struct:
			s = s + v.Field(i).String()
		}
	}
	return s + " }"
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}

func init() {
	/* ms := new(struct1) // 得到：指针类型
	fmt.Println(ms)    // 得到的一个指向默认struct1实例的指针，实例的各个字段为默认值。
	fmt.Println(new([10]int))
	fmt.Println(new(map[string]int))
	ms.i1 = 10
	ms.f1 = 15.5
	(*ms).str = "Chris" // 注意这种赋值方式，也是可以的！

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Printf("%v\n", ms)
	stru := struct1{1, 1.1, "1.1"}   // 得到：值类型
	stru2 := &struct1{2, 2.2, "2.2"} // 得到： 指针类型
	fmt.Printf("%v\n", stru)
	fmt.Printf("%v\n", stru2)
	fmt.Println("ms实例占的内存大小：", unsafe.Sizeof(ms))
	fmt.Println("stru实例占的内存大小：", unsafe.Sizeof(stru))
	fmt.Println("stru2实例占的内存大小：", unsafe.Sizeof(stru2)) */

	// 1. 工厂方法 -- 利用 包的功能，将 结构体matrix 做成私有的，只能通过 工厂方法创建实例
	// matrixer := matrix.NewMatrix("Diana")
	// fmt.Printf("得到的 matrix的实例对象：%v\n", matrixer)

	// 2、reflect 获取 结构的类型
	/* tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	} */

	// 3、结构体匿名字段 + 内嵌结构体
	// 3.1、 使用 new 进行实例化

	/* outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60 // 给匿名字段赋值
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)
	fmt.Println("outer is :", outer) */

	// 3.2、使用字面量方式实例化
	/* 	outer2 := outerS{6, 7.5, 60, innerS{5, 10, 1}} // innerS：内嵌结构体
	fmt.Println("outer2 is:", outer2) */

	// 4、命名冲突-- 重载
	/* 	var bS = new(bStruct)
	   	bS.aStruct.a = 1233  // 给内嵌结构体的字段赋值
	   	bS.a = 1.23   // 重载 a 字段
	   	fmt.Println(bS.aStruct.a, bS.aStruct.b, bS.b, bS.a)

	   	bSS := bStruct{aStruct{123, 321}, 1.3}
	   	fmt.Println(bSS.aStruct, bSS.a) */

	// 5.1、结构体 + 方法
	/* ins := innerS{12, 13, 1}
	fmt.Println("结构体+方法--求和：", ins.sum())

	// 5.2、 非结构体方法
	fmt.Println("非结构体 方法：", intVector{1, 2, 3}.Sum()) */

	// 6、给非本地类型 定义方法  -- 先给类型起一个别名，然后赋方法
	mt := myTime{12, time.Now()}
	fmt.Println("mt is:", mt) // 默认调用 String 方法。
	fmt.Printf("Full time now is :%v\n", mt.Time.String())
	fmt.Printf("First 3 chars is :%v\n", mt.first3Chars())
}

/* func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
} */
