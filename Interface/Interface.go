package Interface

import "fmt"

//关于接口需要注意的是，只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗。
//因为空接口可以存储任意类型值的特点，所以空接口在Go语言中的使用十分广泛。
//关于接口需要注意的是，只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗。
func Test() {
	var x WashingMachine
	h := Haier{}
	x = h
	x.Wash()
	x.Dry()

	// 空接口作为map的值
	//因为空接口可以存储任意类型值的特点，所以空接口在Go语言中的使用十分广泛。
	m := make(map[string]interface{})
	m["name"] = "张三"
	m["age"] = 19
	m["married"] = false
	for k, v := range m {
		fmt.Println(k, ":", v)
	}

	//类型断
	for k, v := range m {
		AssertType(k, v)
	}
}
func AssertType(name string, x interface{}) {
	//t,ok:=x.(string)
	switch t := x.(type) {
	case string:
		fmt.Println(name, "的类型为：string,", t)
	case int:
		fmt.Println(name, "的类型为：int,", t)
	case bool:
		fmt.Println(name, "x的类型为：bool,", t)
	}
}

type WashingMachine interface {
	Dry()
	Wash()
}
type Dryer struct {
}
type Haier struct {
	Dryer
}

func (d Dryer) Dry() {
	fmt.Println("甩干功能")
}
func (h Haier) Wash() {
	fmt.Println("洗衣功能")
}
