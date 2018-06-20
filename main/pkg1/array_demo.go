/**
数组和切片
*/

package pkg1

import "fmt"

func init() {
	// 1、直接实例化具有具体元素的数组。
	/* a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	} */
	var arr1 [5]int        // 得到的是一个值类型
	var arr2 = new([5]int) // new 出来的是一个指针
	// fmt.Printf("%T\n", arr1)
	// fmt.Printf("%T\n", arr2)
	// arr1[2] = 1
	// arr2[2] = 1

	// 2、引用类型的传递属于 引用传递（指针传递）
	arr11 := arr1
	arr22 := arr2
	arr11[2] = 2
	arr22[2] = 2
	// fmt.Printf("%v\n", arr11) // [0 0 2 0 0]
	// fmt.Printf("%v\n", arr22) // &[0 0 2 0 0]
	// fmt.Printf("%v\n", arr1) // [0 0 0 0 0]
	// fmt.Printf("%v\n", arr2) // &[0 0 2 0 0]

	// 3、基本类型的传递属于值传递
	/* 	aa := 1
	   	bb := aa
	   	bb = 2
	   	fmt.Println(bb)
	   	fmt.Println(aa) */

	// 4、切片的声明：
	/* 	slice1 := [...]int{1, 2, 3, 4} // 这是数组
	   	slice2 := []int{1, 2, 3, 4}    // 这是 切片
	   	// slice1[4] = 5	// 会报错
	   	slice2[4] = 5
	   	fmt.Printf("%T\n", slice1)
	   	fmt.Printf("%T\n", slice2) */

	// 5、切片cap（切片的容量）
/* 	slice3 := []int{2, 3, 5, 7, 11}
	fmt.Println(cap(slice3))
	slice4 := slice3[:4]
	fmt.Println(slice4)
	slice4 = slice3[0:]
	fmt.Println(slice4)
	slice4[2] = 12
	fmt.Println(slice4)
	slice4 = slice4[:cap(slice4)]
	fmt.Println(slice4)
	slice4[3] = 13
	fmt.Println(slice4)
	fmt.Println("cap(slice4)", cap(slice4))
	slice4 = slice4[1:cap(slice4)]
	fmt.Println(slice4)
	fmt.Println(cap(slice4)) */

	// 6、用make 来生成 切片
	slice5 := make([]int, 5, 10)
	fmt.Println(slice5)

}
