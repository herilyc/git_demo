package main

import "fmt"

func Array() {
	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	b := [...]int{0, 1, 2, 3}
	var c = [4]int{0, 1, 2, 3}
	fmt.Println(b == c)

	//索引，遍历数组
	for i, value := range b {
		fmt.Println(i, value)
	}
	//仅打印元素
	for _, value := range a {
		fmt.Println(value)
	}
}

func DoubleArray() {
	var a = [4][4]int{
		{0, 1, 2, 3},
		{1, 2, 3},
		{0: 1, 3: 3},
	}
	fmt.Println(a)

	var b [2][2]string
	b = [2][2]string{{"name", "user"}, {0: "password"}}
	fmt.Println(b)

}
