package Pointer

import "fmt"

/*
1.对变量进行取地址（&）操作，可以获得这个变量的指针变量。
2.指针变量的值是指针地址。
3.对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
*/

func Test() {
	a := 10
	//&取地址
	b := &a
	//*根据地址取值
	c := *b
	fmt.Println("a的地址：", &a)
	//b变量存放的是a的地址
	fmt.Println("b变量存放的是a的地址：", b)
	//b自己的地址
	fmt.Println("b自己的地址：", &b)
	//根据b变量存放的地址值，取值
	fmt.Println("根据b变量存放的地址值，取值：", c)
	//根据b自己的地址，*取b上存放的值
	fmt.Println("根据b自己的地址，*取b上存放的值：", *&b)

	//指针传值（类似引用类型）
	p := 100
	modify1(p)
	fmt.Println(p)
	modify2(&p)
	fmt.Println(p)

	//引用类型变量必须初始化才能使用
	//new
	var r *int = new(int)
	*r = 123
	fmt.Println("r值:", *r)
	r1 := new(int)
	//0
	fmt.Println(*r1)
	//make: make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身
	map1 := make(map[string]int)
	map1["张三"] = 123
	fmt.Println(map1)
}

//指针传值
func modify1(x int) {
	x += 100
}
func modify2(x *int) {
	*x += 100
}
