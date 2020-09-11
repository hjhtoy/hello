package Map

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func Test() {
	map1 := make(map[string]int, 8)
	map1["张三"] = 28
	map1["李四"] = 19
	fmt.Println("map1：", map1)

	map2 := map[string]int{
		"王五": 33,
		"赵六": 66,
		"娜扎": 22,
	}
	fmt.Println("map2:", map2)

	//判断key是否存在
	v, ok := map2["王五"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Nothing!")
	}
	for k, v := range map2 {
		fmt.Println("Key:", k, "Value:", v)
	}

	fmt.Println("删除赵六")
	delete(map2, "赵六")
	for k, v := range map2 {
		fmt.Println("Key:", k, "Value:", v)
	}

	//map排序输出
	SortMap()

	//统计一个字符串中每个单词出现的次数
	WordCounter()

	//代码分析
	AboutMap()
}

//按照指定顺序遍历map
func SortMap() {
	//初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	var scoreMap = make(map[string]int, 100)
	for i := 0; i < 100; i++ {
		//生成stu开头的字符串
		key := fmt.Sprintf("stu%02d", i)
		//生成0~99的随机整数
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	fmt.Println("未排序的scoreMap:", scoreMap)

	//定义切片变量存放map的keys
	sortKeys := make([]string, 0, 200)
	for key := range scoreMap {
		sortKeys = append(sortKeys, key)
	}
	//对切片进行排序
	sort.Strings(sortKeys)
	//循环排序的keys，输出map
	fmt.Println("排序后的scoreMap:", sortKeys)
	for _, key := range sortKeys {
		//if key == "" {
		//	continue
		//}
		fmt.Println("key:", key, "value:", scoreMap[key])
	}
}

//作业
func WordCounter() {
	sentence := "how do you do"
	words := strings.Split(sentence, " ")
	fmt.Println("words:", words)
	//map:key-单词，value-单词出现的次数
	wordMap := make(map[string]int, 10)
	for _, word := range words {
		v, ok := wordMap[word]
		if ok {
			wordMap[word] = v + 1
		} else {
			wordMap[word] = 1
		}
	}

	fmt.Println("在字符串", sentence, "中：")
	for k, v := range wordMap {
		fmt.Println(k, "=", v)
	}
}

func AboutMap() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s) //1,2,3
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)         //1,3
	fmt.Printf("%+v\n", m["q1mi"]) //1,3
	fmt.Println(m)

	/*
		打印结果：
		[1 2 3]
		[1 3]
		[1 3 3]
		说明：
		0.切片：包含指针，长度和容量
		1.m["q1mi"] 和 s 指向同一个数组：[1 2 3]
		2.s 切片删除了下标为 1 的元素，直接在共享数组修改，修改后数组为 [1 3 3]
		3.此时 s 的指针，指向底层共享数组，长度为 2，容量为 4（容量为 4，是因为扩容一次）
		4.m["q1mi"] 的指针，也指向底层共享数组，长度为 3，容量为 4
	*/
}
