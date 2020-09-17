package GoroutineExercise

import (
	"fmt"
	"math/rand"
	"time"
)

func Test()  {
	jobChan := make(chan int, 100)
	resultChan := make(chan int, 100)
	go RandomNum(10, jobChan)
	for i:=0;i<24;i++{
		go Process(jobChan, resultChan)
	}
	//close(resultChan)
	//Process(jobChan, resultChan)
	i:=1
	for result := range resultChan{
		fmt.Println("ResulChan", i, ":", result)
		i++
	}

}

func RandomNum(radNum int, ch chan <- int)  {
	total := 0
	for i := 0; i< radNum;i++ {
		rand.Seed(time.Now().UnixNano())
		x :=  rand.Intn(100)
		total += x
		ch <- x
	}
	fmt.Println("生产的随机数总和：", total)
	close(ch)
}

func Process(jobChan <- chan int, resultChan chan <- int){
	for ch := range jobChan{
		resultChan <- ch
	}

	close(resultChan)
}