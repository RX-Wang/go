package pkg1

import (
	"fmt"
)

/**
注释 ：函数接收参数
*/
/* func Test01(name string, age int) {
	fmt.Printf("这是Test01.go, I`m %s ,my age is %d\n", name, age)
} */

/**
	(p1, p2 int)：入参
	(r1 int, r2 string)：返回值
*/
/* func Test02(p1, p2 int) (r1 int, r2 string) {
	fmt.Println("Test02");
	return (p1 + p2), "Test02的第二个返回值，哈哈"
} */

func Test03() {
	const a int64 = 1 << 1
	fmt.Println(a)
}