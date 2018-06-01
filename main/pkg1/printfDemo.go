package pkg1

// import "fmt"

/**
	格式化输出 Demo
*/

/**
		*	%s：输出一个string；
		%t：输出一个bool；
		%v：按照其自身类型输出；
		%d: 格式化整数；
		%x 、%X ：输出16进制数字
		%g：格式化浮点型
		%f：输出浮点数
		%e：输出科学计数表示法
		%0d：用于规定输出定长的整数
		%n.mg：用于表示数字 n 并精确到小数点后 m 位，除了使用 g 之外，还可以使用 e 或者 f，
      例如：使用格式化字符串 %5.2e 来输出 3.4 的结果为 3.40e+00
*/

func init() {
	// fmt.Printf("string: %s\n", "this is a string");
	// fmt.Printf("bool: %t\n", true);
	// fmt.Printf("根据目标原本的类型进行输出: %v\n", "原始类型是string");
	// fmt.Printf("int: %d\n", 1234);
	// fmt.Printf("x16进制: %x\n", 0x1234);
	// fmt.Printf("X16进制: %X\n", 0x12344444);
	// fmt.Printf("浮点型: %g\n", 123.123);
	// fmt.Printf("浮点型: %f\n", 12345.1234);
	// fmt.Printf("科学计数法表示: %e\n", 1e5);
	// fmt.Printf("格式化的科学计数法表示: %5.2e\n", 3.1415926);
}