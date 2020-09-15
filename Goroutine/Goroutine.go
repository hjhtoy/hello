package Goroutine

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Test() {
	for i := 0; i <= 2; i++ {
		wg.Add(1)
		go hello(i)
	}

	wg.Wait()

	//无缓冲的通道
	//使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道。
	c := make(chan int)
	go RecChan(c)
	c <- 10

	//有缓存的通道
	c = make(chan int, 10)

	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}

	//单向通道
	fmt.Println("-----单向通道-----")
	in := make(chan int, 100)
	for i := 0; i < 10; i++ {
		in <- i + 1
	}
	close(in)

	out := make(chan int, 100)
	ReSetChan(out, in)
	fmt.Println("转移后的Out通道")
	for o := range out {
		fmt.Println(o)
	}

	//worker pool（goroutine池）
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	for i := 0; i < 3; i++ {
		go worker(i, jobs, result)
	}

	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 5; i++ {
		<-result
	}

	//通道在接收数据时，如果没有数据可以接收将会发生阻塞。
	//select多路复用
	//select的使用类似于switch语句，它有一系列case分支和一个默认的分支。
	//可处理一个或多个channel的发送/接收操作。
	//如果多个case同时满足，select会随机选择一个。
	//对于没有case的select{}会一直等待，可用于阻塞main函数。
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
		case x := <-ch:
			fmt.Println(x)
		default:
			fmt.Println("none...")
		}
	}
}

func worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Println("worker:", id, ". start job:", job)
		time.Sleep(time.Second)
		fmt.Println("worker:", id, ". end job:", job)
		result <- job * 2
	}
}

func ReSetChan(out chan<- int, in <-chan int) {
	for c := range in {
		out <- c * 10
	}
	close(out)
}

func RecChan(c chan int) {
	x := <-c
	fmt.Println(x)
}

func hello(i int) {
	defer wg.Done()
	fmt.Println("Hello Goroutine!", i)
}
