package pkg1

// import "fmt"

func init() {
	// fallthrough  关键字会使 分支顺延至下一个，无论匹配不匹配。
	/* var num1 int = 10
	switch num1 {
		case 0:
		case 98, 99:
			fmt.Println("num1 is 98 or 99")
		default:
			fmt.Println("...")
	} */
	/* k := 6
	switch k {
		case 4: fmt.Println("was <= 4"); fallthrough;
		case 5: fmt.Println("was <= 5"); fallthrough;
		case 6: fmt.Println("was <= 6"); fallthrough;
		case 7: fmt.Println("was <= 7"); fallthrough;
		case 8: fmt.Println("was <= 8"); fallthrough;
		default: fmt.Println("default case")
	} */
}