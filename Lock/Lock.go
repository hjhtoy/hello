package Lock

import (
	"fmt"
	"sync"
	"time"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex
var lockRW sync.RWMutex
var loadOnce sync.Once
var icons map[string]string

func Test() {
	//wg.Add(2)
	//go add()
	//go add()
	//wg.Wait()
	//fmt.Println(x)
	//
	////读写锁
	//start := time.Now()
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go write()
	//}
	//
	//for i := 0; i < 1000; i++ {
	//	wg.Add(1)
	//	go read()
	//}
	//
	//wg.Wait()
	//end := time.Now()
	//fmt.Println(end.Sub(start))

	//Sync.Once
	Icon("left")

	//并发安全的单例模式
	var ins = GetInstance()
	ins.Hello()
}
type singleton struct {
}
var instance *singleton
var singleOnce sync.Once
func GetInstance() *singleton {
	singleOnce.Do(func() {
		instance = &singleton{}
	})
	return instance
}
func(s *singleton) Hello(){
	fmt.Println("singleton hello")
}
// Icon 是并发安全的
func Icon(name string) string {
	x := 10
	//funcWithParams := func(){
	//	loadIcons(x)
	//}

	loadOnce.Do(func(){
		loadIcons(x)
	})
	return icons[name]
}

func loadIcons(x int) {
	fmt.Println("闭包传过来的参数：", x)
	x = x * 2
	icons = map[string]string{
		"left":  "left.png",
		"up":    "up.png",
		"right": "right.png",
		"down":  "down.png",
	}
}

func write()  {
	//加写锁:
	//当一个goroutine获取-*写锁*-之后，其他的goroutine无论是获取读锁还是写锁都会等待。
	lockRW.Lock()
	time.Sleep(time.Second)
	//解写锁
	lockRW.Unlock()
	wg.Done()
}
func read()  {
	//加读锁：
	//当一个goroutine获取-*读锁*-之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待
	lockRW.RLock()
	time.Sleep(time.Second)
	//解读锁
	lockRW.RUnlock()
	wg.Done()
}
func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}
