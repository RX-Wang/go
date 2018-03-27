package pkg1

import (
	"fmt"
	"runtime"
)

/**
注释
*/
func Version() {
	fmt.Printf("%s\n", runtime.Version())
}
