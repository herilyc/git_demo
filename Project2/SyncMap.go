package main

import (
	"fmt"
	"sync"
	"time"
)

func read(p *sync.Map) {
	//遍历
	p.Range(func(key, value interface{}) bool {
		fmt.Println("iterator: ", key, value)
		return true
	})
}

func SyncMap() {
	var smap sync.Map

	//Store()函数，向smap中添加
	smap.Store("store", "storage")
	smap.Store("load", "load")
	smap.Store("del", "delete")
	smap.Store(1, 1)
	smap.Store("time", time.Local)

	read(&smap)

	//Load()函数返回两个值（value,bool）
	fmt.Println(smap.Load("load"))
	fmt.Println(smap.Load(1))
	fmt.Println(smap.Load(time.Local)) //false
	fmt.Println(smap.Load("time"))     //true

	smap.Delete("del")

	read(&smap)

}
