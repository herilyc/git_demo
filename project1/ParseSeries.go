package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i, j float32 = 10, 0.5
	fmt.Println(i * j)

	//strconv.Itoa() => int->string  func Itoa(i int) string
	num := 100
	str := strconv.Itoa(num)
	fmt.Printf("type:%T\tvalue:%#v\n", str, str)

	//strconv.Atoi() => string->int  func Atoi(s string) (i int, err error)
	str1, str2 := "200", "s100"
	num1, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Printf("%v 转换失败！\n", str1)
	} else {
		fmt.Printf("type:%T\tvalue:%#v\n", num1, num1)
	}
	num2, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Printf("%v 转换失败！\n", str2)
	} else {
		fmt.Printf("type:%T\tvalue:%#v\n", num2, num2)
	}

	//strconv.ParseBool() => string->bool  func ParseBool(str string) (value bool, err error)
	str3, str4 := "100", "True"
	boo1, err := strconv.ParseBool(str3)
	if err != nil {
		fmt.Printf("str3: %v\n", err)
	} else {
		fmt.Println(boo1)
	}
	boo2, err := strconv.ParseBool(str4)
	if err != nil {
		fmt.Printf("str4:%v\n", err)
	} else {
		fmt.Println(boo2)
	}

	//strconv.ParseInt() => string->int  func ParseInt(s string, base int, bitSize int) (i int64, err error)
	//可以指定转换的进制【base】:2-36 base = 0,会从字符串前置判断
	//可以指定转换的返回结果类型【bitSize】:(0,8,16,32,64)=>(int,int8,int16,int32,int64)
	//err: type=*NumErr,若语法有误，err.Eroor = ErrSyntax; 若结果进出类型范围，err.Error = ErrRange
	str5 := "-11"
	num3, err := strconv.ParseInt(str5, 16, 0) //base=16:认为str5中为16进制的-11，bitSize=0:用int型返回
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num3)
	}

	//strconv.ParseUint() => string->uint  func ParseInt(s string, base int, bitSize int) (i uint64, err error)
	str6 := "11"
	num4, err := strconv.ParseInt(str6, 16, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num4)
	}

	//strconv.ParseFloat() => string->float func ParseFloat(s string, bitSize int) (f float64, err error)
	//【bitSize】: (1)、bitSize=32 -> float32	(2)、bitSize=64 -> float64
	str7 := "5"
	num5, err := strconv.ParseFloat(str7, 32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("num=%.3f", num5)
	}

	FormatBool()
	FormatInt()
	FormatUint()
	FormatFloat()

	AppendInt()
	AppendFloat()
	AppendBool()
}
