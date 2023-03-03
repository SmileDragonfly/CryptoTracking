package main

import (
	"fmt"
	"runtime"
)

func getCurrentFuncname() string {
	pc, _, _, _ := runtime.Caller(1)
	return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
}
