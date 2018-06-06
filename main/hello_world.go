package main

import (
	V "./pkg1" // 给包起一个别名
	// "./trains" // 直接使用包的名字
	"fmt"
)

var a = "G"

func main() {
	fmt.Println("hello world")
	// V.Version()
	// V.Test01("Diana", 10)
	// fmt.Println(V.Test02(1,2))
	// V.ForLoop()
	// V.Test03()
	// V.Changliang()
	// V.Bianliang()
	V.Bl = 123
	// fmt.Println(V.Bl)
	// fmt.Println(V.ClConst)
	// V.Goos()
	// fmt.Println(trains.Pi)
	// ============ 练习题：
	// n()
  // m()
	// n()
	// a = "G"
  // print(a)
	// f1()
	// print(`This is a raw string \n`) // 会原样输出
}

/* func n() { print(a) }

func m() {
   a = "O"
   print(a)
} */

/* func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(a)
} */