package pkg1

import (
	"fmt"
)

/**
变量
*/
var (
	aa int
	bb string
	cc bool
	dd *int
	ee *string
)

/**
未赋值的变量初始值为：
int 为 0，
float 为 0.0，
bool 为 false，
string 为空字符串，
指针为 nil；
*/
func Bianliang() {
	// fmt.Println(aa,bb,cc,dd,ee)

	// var i = "12" // 普通声明方式
	// 简短式声明：
	i := "12"
	j := i
	fmt.Printf("i : %d\n", &i) // &i: 用于获取i 在内存中得地址
	fmt.Printf("j : %d\n", &j)
}
