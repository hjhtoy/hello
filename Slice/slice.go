package Slice

import (
	"fmt"
	"sort"
)

func Test() {
	a := [7]int{1, 2, 3, 4, 5, 6, 7}
	s1 := a[0:3]

	fmt.Println(len(s1), cap(s1))

	b := make([]int, 2, 10)
	b[0] = 1
	b[1] = 2
	fmt.Println("b:", b, "len:", len(b), "cap:", cap(b))
	c := b
	c[0] = 11
	fmt.Println("b:", b, "len:", len(b), "cap:", cap(b))
	var d []int
	d = append(d, b...)
	d = append(d, c...)
	fmt.Println("-------", d)
	//删除
	d = append(d[:2], d[3:]...)
	fmt.Println("---删除11----", d)
	var e []int
	e = append(e, 1, 2, 3, 4, 5, 6)
	fmt.Println("e:", e)
	//复制
	f := make([]int, 10, 20)
	copy(f, e)
	fmt.Println("----f:", f)

	//数组排序
	var arrSort = [...]int{3, 7, 8, 9, 1}
	var sliceSort = arrSort[:]
	sort.Ints(sliceSort)
	//for index, value := range sliceSort {
	//	arrSort[index] = value
	//}
	fmt.Println("数组排序：", arrSort)
}
