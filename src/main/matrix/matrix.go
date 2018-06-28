/**
	利用 包的功能，将 结构体matrix 做成私有的，只能通过 工厂方法创建实例
*/

package matrix

type matrix struct {
	name string
}

// 工厂方法，强制外围的包使用该方法进行实例化

func NewMatrix(params string) *matrix {
	m := new(matrix) // 初始化 m
	m.name = params
	return m
}