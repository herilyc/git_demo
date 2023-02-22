package main

import (
	"fmt"
	"sort"
)

func Map() {
	//map定义时必须要初始化空间，否则无法使用
	m := make(map[string]int, 10)          //map[keyType]valueType
	m = map[string]int{"one": 1, "two": 2} //key : value
	m["three"] = 3                         //后续添加到map中
	fmt.Println(len(m))                    //map无法判断capacity,不可使用cap()函数
	for k, v := range m {
		fmt.Println(k, v)
	}

	//map的多键索引 map[keyType] []valueType || map[keyType] *[]valueType
	var m1 = map[string][]int{"one": {1, 2}, "two": {2, 3, 4}}
	m1["three"] = []int{3, 4, 5}
	fmt.Println(m1) //会按string型keyType的字典序进行输出
	fmt.Println(m1["two"])
}

func MapConMulti() {
	var map1 map[int]float64
	map1 = map[int]float64{3: 0.5, 4: 0.6, 5: 0.7}
	map1[1] = 3.1415926
	map1[2] = 0.123
	//map1[0] = 0.1
	//i代表的是map的key值，最好不要用，
	//可能会有i（key）的值无对应value的情况，会自动被变为零值
	for i := 1; i <= len(map1); i++ { //i := 0 开始就会出现无对应值变成0值的情况
		fmt.Println(map1[i])
	}

	//声明一个用于存放map1元素的切片
	var mSlice []int      //key
	var mSlice2 []float64 //value
	for i, f := range map1 {
		mSlice = append(mSlice, i)
		mSlice2 = append(mSlice2, f)
	}
	sort.Ints(mSlice) //根据map1中的key来进行排序
	sort.Float64s(mSlice2)
	fmt.Println(mSlice) //根据map1中的value来进行排序输出
	fmt.Println(mSlice2)

	for i := 1; i <= len(mSlice); i++ {
		fmt.Println(mSlice[i-1], mSlice2[mSlice[i-1]-1])
	}

	delete(map1, 1)
	fmt.Println(map1)
	fmt.Println(mSlice) //切片不会同步更新
}
