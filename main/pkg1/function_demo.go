// 这是一个 函数 的demo
package pkg1

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
)

var num int = 10
var numx2, numx3 int

func init() {
	// Pop()
	// fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))
	// var i1 int = MultiPly3Nums(2, 5, 6)
	// fmt.Printf("MultiPly 2 * 5 * 6 = %d\n", i1)

	// 2/3、两种返回值类型
	// numx2, numx3 = getX2AndX3(num)
	// PrintValues()
	// numx2, numx3 = getX2AndX3_2(num)
	// PrintValues()

	// 4、改变外部变量（指针传递）
	/* 	n := 0
	   	reply := &n
	   	Multiply(10, 5, reply)
	   	fmt.Println("Multiply:", *reply) // Multiply: 50
	   	fmt.Println("n :", n) */

	// 5、变长参数的函数
	/* x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	slice := []int{7, 9, 3, 5, 1}
	// 传递 slice 类型的变量时可以这样传递：slice...
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d\n", x) */

	// 6、 defer 和追踪
	// function1()
	// abc()
	// func1("Go")
	func2()

	// 练习题：
	/* ret1, err1 := MySqrt(-1)
	if err1 != nil {
		fmt.Println("Error! Return values are: ", ret1, err1)
	} else {
		fmt.Println("It's ok! Return values are: ", ret1, err1)
	}

	fmt.Print("Second example with 5: ")
	// if 变量赋值 + 条件判断
	if ret2, err2 := MySqrt(5); err2 != nil {
		fmt.Println("Error! Return values are: ", ret2, err2)
	} else {
		fmt.Println("It's ok! Return values are: ", ret2, err2)
	}
	// named return variables:
	fmt.Println(MySqrt2(5)) */
}

// 1、函数定义
func MultiPly3Nums(a int, b int, c int) int {
	// var product int = a * b * c
	// return product
	return a * b * c
}

func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

// 2、函数定义--未命名返回值 -- return 后面需要跟上需要返回的值
func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

// 3、函数定义--命名返回值 -- return 后面 不需要跟上需要返回的值--即使只有一个命名返回值，也需要使用 () 括起来
func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3
	return
}

// 4、改变外部变量（指针 传递）
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

// 5、 变长参数的函数

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

// 6、 defer 和追踪
func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!")
}

func abc() {
	i := 0
	defer fmt.Println(i) // 0
	i++
	fmt.Println(i) // 1
	return
}

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}

func func2() {
	i := 0
	defer ff(i)
	defer func() {
		fmt.Println("打印--3：", i)
	}()
	defer fmt.Println("打印--2：", i)
	i++
	fmt.Println("打印--1：", i)
}
func ff(i int) {
	fmt.Println("打印--4：", i)
}

//  练习：
func MySqrt(f float64) (float64, error) {
	//return an error as second parameter if invalid input
	if f < 0 {
		return float64(math.NaN()), errors.New("I won't be able to do a sqrt of negative number!")
	}
	//otherwise use default square root function
	return math.Sqrt(f), nil
}

//name the return variables - by default it will have 'zero-ed' values i.e. numbers are 0, string is empty, etc.
func MySqrt2(f float64) (ret float64, err error) {
	if f < 0 {
		//then you can use those variables in code
		ret = float64(math.NaN())
		err = errors.New("I won't be able to do a sqrt of negative number!")
	} else {
		ret = math.Sqrt(f)
		//err is not assigned, so it gets default value nil
	}
	//automatically return the named return variables ret and err
	return
}

/* // 该函数 编译错误
func (st *Stack) Pop() int {
	v := 0
	for ix := len(st) - 1; ix >= 0; ix-- {
			if v = st[ix]; v != 0 {
					st[ix] = 0
					return v
			}
	}
} */
