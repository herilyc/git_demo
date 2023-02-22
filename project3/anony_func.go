package main

import "fmt"

// 遍历切片的每个元素, 通过给定函数进行元素访问
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

func AF() {
	//使用匿名函数打印切片内容
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})

	//匿名函数的定义与赋值，f可以被调用，若定义在外部会出错
	f := func(v int) {
		fmt.Println(v)
	}

	//visit([]string{"jack", "rose", "judy", "tom"}, f)

	var t = []int{2, 3, 4, 5}
	for _, v := range t {
		f(v)
	}
}
