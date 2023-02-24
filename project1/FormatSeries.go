package main

import (
	"fmt"
	"strconv"
)

func FormatBool() {
	num := true
	str := strconv.FormatBool(num)
	fmt.Printf("\n\ntype:%T\t value:%#v\n", str, str)
}

func FormatInt() {
	var num int64 = -100
	str := strconv.FormatInt(num, 16)
	fmt.Printf("type:%T\t value:%#v\n", str, str)
}

func FormatUint() {
	var num uint64 = 100
	str := strconv.FormatUint(num, 16)
	fmt.Printf("type:%T\t value:%#v\n", str, str)
}

func FormatFloat() {
	var num float64 = 3.1425926
	str := strconv.FormatFloat(num, 'E', -1, 32)
	fmt.Printf("type:%T\t value:%#v\n", str, str)
}

func AppendInt() {
	b10 := []byte("int (base 10):")
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))

	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))
}

func AppendFloat() {
	f1 := []byte("float32 value:")
	f1 = strconv.AppendFloat(f1, 3.1415926, 'E', -1, 32)
	fmt.Println(string(f1))

	f2 := []byte("float64 value:")
	f2 = strconv.AppendFloat(f2, 3.1415926, 'f', -1, 64)
	fmt.Println(string(f2))
}

func AppendBool() {
	b1 := []byte("b1 bool:")
	b1 = strconv.AppendBool(b1, true)
	fmt.Println(string(b1))

	b2 := []byte("b2 bool:")
	b2 = strconv.AppendBool(b2, false)
	fmt.Println(string(b2))
}
