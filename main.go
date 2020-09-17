package main

import "hello/Socket"

var m = 100

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)
const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
	g, h
)

func main() {
	//sayhi.Sayhi()

	////切片
	//Slice.Test()

	////Map
	//Map.Test()

	////函数
	//Function.Test()

	////指针
	//Pointer.Test()

	////结构体
	//Struct.Test()

	////接口
	//Interface.Test()

	////goroutine
	//Goroutine.Test()

	////线程安全和锁
	//Lock.Test()

	////并发和安全练习题
	//GoroutineExercise.Test()

	//网络
	Socket.Test()




	////面向对象
	//oop.Test()
}

