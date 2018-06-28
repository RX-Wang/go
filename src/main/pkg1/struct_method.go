/**
结构体 和 方法
*/

package pkg1

import (
	"fmt"
	// "unsafe"

	"reflect"
)

type struct1 struct {
	i1  int "this is i1"
	f1  float32
	str string
}

func init() {
	ms := new(struct1) // 得到：指针类型
	fmt.Println(ms)    // 得到的一个指向默认struct1实例的指针，实例的各个字段为默认值。
	fmt.Println(new([10]int))
	fmt.Println(new(map[string]int))
	// stru := struct1{1, 1.1, "1.1"}   // 得到：值类型
	// stru2 := &struct1{2, 2.2, "2.2"} // 得到： 指针类型
	ms.i1 = 10
	ms.f1 = 15.5
	(*ms).str = "Chris" // 注意这种赋值方式，也是可以的！

	// fmt.Printf("The int is: %d\n", ms.i1)
	// fmt.Printf("The float is: %f\n", ms.f1)
	// fmt.Printf("The string is: %s\n", ms.str)
	// fmt.Printf("%v\n", ms)
	// fmt.Printf("%v\n", stru)
	// fmt.Printf("%v\n", stru2)
	// fmt.Println("ms实例占的内存大小：", unsafe.Sizeof(ms))
	// fmt.Println("stru实例占的内存大小：", unsafe.Sizeof(stru))
	// fmt.Println("stru2实例占的内存大小：", unsafe.Sizeof(stru2))

	// reflect --使用
	msType := reflect.TypeOf(ms)
	fmt.Println(msType)
	i1Field := msType.Field(0)
	fmt.Println(i1Field)
	fmt.Println(i1Field.Tag)

	// 1. 工厂方法 -- 利用 包的功能，将 结构体matrix 做成私有的，只能通过 工厂方法创建实例
	// matrixer := matrix.NewMatrix("Diana")
	// fmt.Printf("得到的 matrix的实例对象：%v\n", matrixer)
}
