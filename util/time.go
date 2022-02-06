package util

import (
	"fmt"
	"time"
)

func TimeFunc(f func()) {
	st := time.Now()
	f()
	fmt.Println("\nfunc exec time: ", time.Since(st))
}
