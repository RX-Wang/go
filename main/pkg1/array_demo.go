/**
数组和切片
*/

package pkg1

import (
	"fmt"
)

func init() {
	// 1、直接实例化具有具体元素的数组。
	/* a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	} */
	var arr1 [5]int	// 得到的是一个值类型
	var arr2 = new([5]int) // new 出来的是一个指针
	// fmt.Printf("%T\n", arr1)
	// fmt.Printf("%T\n", arr2)
	// arr1[2] = 1
	// arr2[2] = 1
	arr11 := arr1
	arr22 := arr2
	arr11[2] = 2
	arr22[2] = 2
	// fmt.Printf("%v\n", arr11) // [0 0 2 0 0]
	// fmt.Printf("%v\n", arr22) // &[0 0 2 0 0]
	// fmt.Printf("%v\n", arr1) // [0 0 0 0 0]
	// fmt.Printf("%v\n", arr2) // &[0 0 2 0 0]
	aa := 1
	bb := aa
	bb = 2
	fmt.Println(bb)
	fmt.Println(aa)
}
