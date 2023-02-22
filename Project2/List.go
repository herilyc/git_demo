package main

import (
	"container/list"
	"fmt"
)

func readListN(list list.List) {
	for i := list.Front(); i != nil; i = i.Next() {
		//fmt.Println(i)
		fmt.Printf("%v\t", i.Value)
	}
	fmt.Println()
}

func readListP(list list.List) {
	for i := list.Back(); i != nil; i = i.Prev() {
		//fmt.Println(i)
		fmt.Printf("%v\t", i.Value)
	}
	fmt.Println()
}

func List() {
	//声明的两种方式
	var a list.List
	b := list.New() //通过new()来声明，返回的是一个指针
	for i := 0; i < 3; i++ {
		c := 'a' + i
		b.PushBack(c)
	}
	readListP(*b) //通过 * 来取指针

	a.PushFront(1)
	intElement := a.PushBack(2)
	a.PushBack("three")
	a.PushFront("zero")
	readListN(a)

	a.InsertAfter(2.5, intElement) //将2.5插入在intElement后面，intElement为指针类型，指向list b中和2，相当于在2后面插入2.5
	a.InsertBefore("one", intElement)
	readListN(a)

	a.Remove(intElement)
	readListN(a)

	a.PushBackList(b)
	readListN(a)
	readListP(a)
}
