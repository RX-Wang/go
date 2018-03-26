package version

import (
	"fmt"
	"runtime"
)

/**
注释
*/
func Version() {
	fmt.Printf("%s", runtime.Version())
}
