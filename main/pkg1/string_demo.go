/**
字符串
*/
package pkg1

// "fmt"
// "strings"
// "strconv"

func init() {
	/**
	1、字符串拆分
	*/

	/* str := "The quick   brown  fox jumps  over the   lazy dog"

	// strings.Fields(str) 会默认根据 str 中 的空格进行拆分，并返回一个  slice。
	sl := strings.Fields(str)
	fmt.Printf("use Fields to split in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()

	// strings.Split(str, sep) 会根据指定的 sep 来对str进行拆分，并返回一个  slice。
	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str3) */

	/**
	2、字符串 与其他类型之间的转换
	*/
	/* var orig string = "666"
	var an int
	var newS string
	var float_64 float64 = 12.11111111111111111111111

	// 获取当前操作系统 ints 的字节长度
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	// string -> int
	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	// int -> string
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)

	fmt.Printf("float64 -> string: %s\n", strconv.FormatFloat(float_64, 'b', 6, 64))
	fmt.Printf("float64 -> string: %s\n", strconv.FormatFloat(float_64, 'e', 6, 64))
	fmt.Printf("float64 -> string: %s\n", strconv.FormatFloat(float_64, 'f', 6, 64))
	fmt.Printf("float64 -> string: %s\n", strconv.FormatFloat(float_64, 'g', 6, 64)) */

}
