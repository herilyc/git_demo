package main

import (
	"fmt"
)

/*
	从数组或切片生成新的切片拥有如下特性：
	1、取出的元素数量为：结束位置 - 开始位置；
	2、取出元素不包含结束位置对应的索引，切片最后一个元素使用 slice[len(slice)] 获取；
	3、当缺省开始位置时，表示从连续区域开头到结束位置；
	4、当缺省结束位置时，表示从开始位置到整个连续区域末尾；
	5、两者同时缺省时，与切片本身等效；
	6、两者同时为 0 时，等效于空切片，一般用于切片复位。
*/

func slice() {
	//go数组的特殊用法（切片，引用类型）
	str1 := [3]string{"num1", "num2", "num3"}
	fmt.Println(str1)
	temp := str1[1:] //[1:]从下标为1的地方开始，将str1中后续所有元素赋给temp(与str1同类型)
	temp1 := str1[:] //[:]两边都无参数：表示原有的切片，即str1
	fmt.Println(temp)
	fmt.Println(temp1)

	temp[0] = "differ"
	temp[1] = "3"
	//temp[2] = "null" 		panic: runtime error: index out of range [2] with length 2
	fmt.Println(str1)

	var highRiseBuilding [30]int
	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}
	// 区间 => 左闭右开->【a,b)
	fmt.Println(highRiseBuilding[10:15])
	// 中间到尾部的所有元素
	fmt.Println(highRiseBuilding[20:])
	// 开头到中间指定位置的所有元素
	fmt.Println(highRiseBuilding[:2])

	//重置切片，清空拥有的元素
	fmt.Println(highRiseBuilding[0:0])

	//声明切片的方法，以及声明后是否有分配内存空间
	var list1 []int
	var list2 []int
	var list3 = []string{}
	fmt.Println(list1 == nil)
	fmt.Println(list2 == nil)
	fmt.Println(list3 == nil)
	fmt.Println(len(list3), cap(list3))
}

// make()函数，	make( []Type, size, cap )
/*
	其中 Type 是指切片的元素类型，
	size 指的是为这个类型分配多少个元素，
	cap 为预分配的元素空间数量，
	这个值设定后不影响 size，只是能提前分配空间，
	为的是降低多次分配空间造成的性能问题
*/
func sliceMake() {
	var a = make([]int, 2)
	b := make([]complex64, 2, 10)
	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))
}

/*
使用 append() 函数为切片动态添加元素时，
如果空间不足以容纳足够多的元素，切片就会进行“扩容”
容量的扩展规律是按原容量的 2 倍数进行扩充 （1，2，4，8，16，32，64 …………）
*/
func sliceAppend() {
	var a []int //声明一个空切片
	//在尾部添加
	a = append(a, 1)
	fmt.Println(a, len(a), cap(a))
	a = append(a, 1, 2, 3) //追加多个元素, 手写解包方式
	fmt.Println(a, len(a), cap(a))
	a = append(a, []int{4, 5, 6}...) //追加一个切片, 切片需要解包（用...）
	fmt.Println(a, len(a), cap(a))

	//在头部添加
	//在切片开头添加元素一般都会导致内存的重新分配，而且会导致已有元素全部被复制 1 次
	//性能差，一般不用
	a = append([]int{1}, a...)
	fmt.Println(a, len(a), cap(a))

	i, x := 3, 7
	a = append(a[:i], append([]int{x}, a[i:]...)...) // 在第i个位置插入x
	fmt.Println(a, len(a), cap(a))
	a = append(a[:i], append([]int{1, 2, 3}, a[i:]...)...) // 在第i个位置插入切片
	fmt.Println(a, len(a), cap(a))
}
func sliceCopy() {
	// 设置元素数量为10
	const elementCount = 10
	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)
	// 将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	fmt.Println(srcData)

	// 引用切片数据
	refData := srcData
	// 预分配足够多的元素切片
	copyData := make([]int, elementCount)
	// 将数据复制到新的切片空间中
	copy(copyData, srcData)
	// 修改原始数据的第一个元素
	srcData[0] = 999
	// 打印引用切片的第一个元素
	fmt.Println(refData[0])
	fmt.Println(refData)

	// 打印复制切片的第一个和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount-1])
	// 复制原始数据从4到6(不包含)
	copy(copyData, srcData[4:6])
	fmt.Println(copyData)
}

func sliceDelFromStart() {
	var a []int
	a = []int{1, 2, 3, 4, 5, 6, 7}
	var n int
	fmt.Scanln(&n)
	a = append(a[:0], a[1:]...) //删除开关一个元素，同a = a[1:]
	fmt.Println(a)
	a = a[n:] //删除前n个元素，同a = append(a[:0], a[n:]...)
	fmt.Println(a)
	copy(a, a[n:]) //可以删除前n位，但后续也会自动补全容量
	fmt.Println(a)
	a = a[:copy(a, a[n:])]
	fmt.Println(a)
}

func sliceDelBetween() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	var i, n int
	fmt.Scanln(&i)
	fmt.Scanln(&n)
	a = append(a[:i], a[i+n:]...) //从i开始删除n个元素
	fmt.Println(a)
	fmt.Println(copy(a[i:], a[i+1:]))
	copy(a[i:], a[i+1:])
	fmt.Println(a)
	a = a[:i+copy(a[i:], a[i+1:])] //再从新切片的i开始删除1个元素
	fmt.Println(a)
}

func sliceDelFromLast() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	var n int
	fmt.Scanln(&n)
	a = a[:len(a)-n]
	fmt.Println(a)
}

func DoubleSlice() {
	// 声明一个二维整型切片并赋值
	slice := [][]int{{10}, {100, 200}}
	// 为第一个切片追加值为 20 的元素
	slice[0] = append(slice[0], 20, 30)
	fmt.Println(slice)
	DSread(slice)
}

// DSread for循环中的高级break
func DSread(slice [][]int) {
JLoop:
	for j := 0; j < 2; j++ {
		for i := 0; i < 2; i++ {
			if i == 1 {
				break JLoop
			}
			fmt.Printf("%d\t", slice[i][j])
		}
		fmt.Println()
	}

}
