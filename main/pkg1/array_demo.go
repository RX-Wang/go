/**
数组和切片
*/

package pkg1

import (
	// "bytes"
	"fmt"
)

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

	// 6、用make / new 来生成 切片
	/* slice5 := make([]int, 5, 10) // 返回一个 长度为：5 容量为：10 的切片
	fmt.Printf("%T\n", slice5)

	// slice6 := new([]int)  // 或者：
	var slice6 *[]int = new([]int) // 返回一个 指向：长度为：0 ；容量为：0 的切片的指针

	fmt.Printf("%T\n", slice6)
	fmt.Println(slice6)
	fmt.Println(len(*slice6), "--", cap(*slice6))

	slice7 := new([5]int) // 返回一个 长度为 5 ；容量为：5 的数组
	fmt.Println(len(*slice7), "--", cap(*slice7))
	*slice6 = slice7[2:] // 改变 切片 slice6 的长度 为 slice7的长度
	fmt.Println(len(*slice6), "--", cap(*slice6))
	*slice6 = slice7[1:]
	fmt.Println(len(*slice6), "--", cap(*slice6)) */

	// 7、 bytes buffer
	/* var buffer *bytes.Buffer = new(bytes.Buffer)
	fmt.Printf("%T\n", buffer)
	fmt.Println(*buffer) */
	/* const golang = "golang is a good language"
	var buf bytes.Buffer
	for ind := range golang {
		buf.WriteByte(golang[ind])
	}
	fmt.Println(buf)
	fmt.Println(buf.Bytes()) // 将 buffer 转化成  []byte 切片
	fmt.Println(buf.String())
	bufBytes := []byte{1, 2, 3, 4, 5, 'a', 'b', 'c', 'd', 'e'}
	fmt.Println(bufBytes)
	buf1 := bytes.NewBuffer(bufBytes)
	fmt.Println(buf1) */

	// 8、切片的 append、remove、insert 操作
	/* sl := []byte{0, 1, 2, 3, 4, 5}
	fmt.Println("sl--", sl)
	byteArr := []byte{6, 7, 8, 9, 10}
	fmt.Println("byteArr--", byteArr)

	// 8.1、append -- 末尾追加  ，切片： slice... 相当于 EcmaScript 6中 的数组解构
	sl = append(sl, byteArr...)
	fmt.Println("new sl--append：", sl)

	// 8.2、remove -- 删除 index = 3 的值
	sl = append(sl[:3], sl[4:]...)
	fmt.Println("new sl--remove：", sl)

	// 8.3、insert -- 插入操作 在inde = 3 位置重新插入 3
	// slEndPart := sl[3:]  // 如果这样写的话，结果会是：[0 1 2 3 3 5 6 7 8 9 10] ，这是错误的，原因是：切片的赋值属于指针赋值，而非 值拷贝赋值。
	slEndPart := append([]byte{}, sl[3:]...) // 这样写是正确的，结果是：[0 1 2 3 4 5 6 7 8 9 10]
	fmt.Println("slEndPart", slEndPart)
	sl = append(sl[:3], 3)
	fmt.Println("append 后的 slEndPart：", slEndPart) // 这里可以验证 两个 · slEndPart · 的差别
	fmt.Println("new sl--insert1：", sl)
	sl = append(sl, slEndPart...)
	fmt.Println("new sl--insert2：", sl) */
	aa := [5]int{1,2,3,4,5}
	fmt.Println(len(aa[2:2]))
}
