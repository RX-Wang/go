package pkg1

/**
数值转换
*/

import (
	"fmt"
	"math"
)

/**
	通过位移运算 配合 iota 实现 存储单位的枚举
 */
type ByteSize float64

const (
	_           = iota // 通过赋值给空白标识符来忽略值
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func init() {
	// var a, error = uint8FromInt(-1);
	// fmt.Printf("a is : %d\n", a);
	// fmt.Println(error);
	// var fl = intFromFloat64(12.56);
	// var fl = intFromFloat64(-2147483649);
	// fmt.Printf("fl is:%d\n", fl);
	// fmt.Println("MinInt32:", math.MinInt32);
	// fmt.Println(1 ^ 1);
	// fmt.Println(1 ^ 0);
	// fmt.Println(2 ^ 3);
	// fmt.Println(^2); // 相当于JS 中的 ~ 按位取反
	// fmt.Println(^-3);
	/* fmt.Printf("%1.4e\n%1.4e\n%1.4e\n%1.4e\n%1.4e\n%1.4e\n%1.4e\n%1.4e\n", KB, MB,
		GB,
		TB,
		PB,
		EB,
		ZB,
		YB) */
}

/**
int 转 无符号 int8
*/
func uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 { // conversion is safe
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

/**
float64 转 int
*/
func intFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 { // x lies in the integer range
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}
