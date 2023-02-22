package main

import "fmt"

// 自定义类型
type newInt int

// 可变参数的传参以及返回值的预定义
func variableVar(args ...newInt) (sum newInt) {
	for i := range args {
		sum += args[i]
	}
	return
}

// 自定义新func 类型
type newFunc func(newInt)

func simpleVariableFunc() {
	var a = []newInt{1, 2, 3, 4}
	f := variableVar
	fmt.Println(f(a...))
}
