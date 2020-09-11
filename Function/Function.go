package Function

import "fmt"

type calculation func(int, int) int

func Test() {
	muiltParams("aaaa", "bbbb", "cccc", "dddd")
	a := sum
	b := sub
	fmt.Println(a(10, 20))
	fmt.Println(b(20, 10))
	//函数做完参数
	fmt.Println(cal3(33, 44, a))
	//函数做为返回值
	fmt.Println(op("+")(123, 456))

	//匿名函数
	//匿名函数保存到变量
	m := func(a, b int) int {
		return a + b
	}
	fmt.Println(m(1111, 2222))
	//匿名函数立即执行
	v := func(a, b int) int {
		return a + b
	}(3333, 4444)
	fmt.Println(v)

	//闭包
	f := adder()
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60
	f1 := adder()
	fmt.Println(f1(30))
	fmt.Println(f1(33))
	fmt.Println(f1(44))

	//defer
	fmt.Println(ff1())
	fmt.Println(ff2())
	fmt.Println(ff3())
	fmt.Println(ff4())
	//defer面试题
	x := 1
	y := 2
	defer calcDefer("AA", x, calcDefer("A", x, y))
	x = 10
	defer calcDefer("BB", x, calcDefer("B", x, y))
	y = 20

	//panic/recover
	funcA()
	funcB()
	funcC()
}

func funcA() {
	fmt.Println("Func A")
}
func funcB() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Recover panic in funcB")
		}
	}()

	panic("panic in funcB")
}
func funcC() {
	fmt.Println("Func C")
}

func calcDefer(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

//5
func ff1() int {
	x := 5
	defer func() {
		x++
	}()
	return x //相当于returnValue = 5
}

//6
func ff2() (x int) {
	defer func() {
		x++
	}()
	return 5 //相当于returnValue = x
}

//5
func ff3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //相当于returnValue = y = 5
}

//5
func ff4() (x int) {
	defer func(x int) {
		//x为值传递，不影响匿名函数外面
		x++
	}(x)
	return 5 //相当于returnValue = x
}

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func cal3(a, b int, op func(int, int) int) int {
	return op(a, b)
}
func op(s string) func(int, int) int {
	switch s {
	case "+":
		return sum
	case "-":
		return sub
	default:
		return nil
	}
}

func sum(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}

func muiltParams(x ...string) {
	for index, value := range x {
		fmt.Println("参数", index, "，参数值：", value)
	}
}

func calc1(a, b int) (int, int) {
	sum := a + b
	sub := a - b

	return sum, sub
}

func calc2(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b

	return
}
