package Struct

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	Age   int
	Name  string
	Color string
	Hobby string
	home  string
}

func (c *Cat) Home() string {
	return c.home
}

func (c *Cat) SetHome(home string) {
	c.home = home
}

func NewCat(name string, age int, color string, hobby string) *Cat {
	return &Cat{Name: name, Age: age, Color: color, Hobby: hobby}
}

func init(){
	fmt.Println("这里是结构体测试包......")
}
func Test() {
	//var cat1 = Cat{
	//	Name:  "小白",
	//	Age:   10,
	//	Color: "白色",
	//	Hobby: "吃小鱼干",
	//}
	//fmt.Println(cat1)
	//fmt.Printf("Cat1地址：%p\n", &cat1)
	//
	//p1 := person{
	//	name: "李四",
	//	city: "南宁",
	//	age:  20,
	//}
	//
	//p2 := new(person)
	//p2.name = "张三"
	//p2.city = "北京"
	//p2.age = 20
	//
	//p3 := &person{
	//	name: "王五",
	//	city: "杭州",
	//	age: 20,
	//}
	//p4:=person{
	//	"赵六",
	//	"上海",
	//	20,
	//}
	//
	//addAge1(p1)
	//fmt.Println("p1:", p1)
	//addAge2(p2)
	//fmt.Println("p2", p2)
	//addAge2(p3)
	//fmt.Println("p3", p3)
	//addAge1(p4)
	//fmt.Println("p4", p4)

	////面试题
	//m := make(map[string]*student)
	//stus := []student{
	//	{name: "小王子", age: 18},
	//	{name: "娜扎", age: 23},
	//	{name: "大王八", age: 9000},
	//}
	//for _, stu := range stus {
	//	////根本原因在于for-range会使用同一块内存去接收循环中的值。
	//	//m[stu.name] = &stu
	//
	//	newStu := stu
	//	//fmt.Printf("stu Type: %T\n", stu)
	//	fmt.Printf("stu adrr: %p\n", &stu)
	//	//fmt.Printf("newStu Type: %T\n", newStu)
	//	fmt.Printf("newStu adrr: %p\n", &newStu)
	//	m[stu.name] = &newStu
	//}
	//
	//for k, v := range m {
	//	fmt.Println(k, "=>", v.name)
	//}
	//
	////方法
	//newP := newPerson("张三", "北京", 23)
	//newP.Dream()

	//JSON序列化 (可系列化的字段必须是public)
	c1 := make([]*class, 0, 10)
	c1 = append(c1, &class{
		title: "一班1",
		students: []*student{
			&student{
				name: "张三1",
				age:  12,
			}, &student{
				name: "李四1",
				age:  20,
			}, &student{
				name: "王五1",
				age:  18,
			},
		},
	},
		&class{
			title: "二班2",
			students: []*student{
				&student{
					name: "张三2",
					age:  12,
				}, &student{
					name: "李四2",
					age:  20,
				}, &student{
					name: "王五2",
					age:  18,
				},
			},
		},
		&class{
			title: "三班3",
			students: []*student{
				&student{
					name: "张三3",
					age:  12,
				}, &student{
					name: "李四3",
					age:  20,
				}, &student{
					name: "王五3",
					age:  18,
				},
			},
		},
	)

	fmt.Println("--C1:--", *c1[0].students[0], *c1[1], *c1[2])
	fmt.Println(json.Marshal(c1))
	data, err := json.Marshal(c1[0])
	if err != nil {
		fmt.Println("C1 to json error:", err)
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("new json:%s\n", data)

	pp := personNew{
		Name: "张三",
		City: "北京",
		age:  15,
	}
	pData, _ := json.Marshal(pp)
	fmt.Println(pp)
	fmt.Printf("new person json:%s\n", pData)
}

type personNew struct {
	Name string		`json:"FullName"`
	City string		`json:"LivingCity"`
	age  int		//私有字段不能被json包访问
}
type class struct {
	title    string
	students []*student
}

type student struct {
	name string
	age  int
}

type person struct {
	name string
	city string
	age  int
}

//person的构造函数
func newPerson(name, city string, age int) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

//person类型的方法
func (p *person) Dream() {
	fmt.Println(p.name, "的梦想是一夜暴富!")
}

func addAge1(p person) {
	p.age += 10
}
func addAge2(p *person) {
	p.age += 12
}

func AddNew(a, b int) int {
	return a + b
}
