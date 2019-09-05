package sample

/**
 * sync.WaitGroup 用于等待一组 goroutine 结束
 * 		Add(int) / Done() / Wait()
 * sync.Mutex 互斥锁
 * 		Lock() / Unlock()
 * sync.RWMutex 读写锁
 * 		RLock() / RUnlock()
 * 		Lock() / Unlock()
 * sync.Pool 池
 * 		构造函数
 * 		Put() / Get()
 * 		当GC（garbage collector）时会回收pool中未使用的对象
 * sync.Once
 * 		Do()
 * sync.Cond 条件等待
 * 		NewCond()
 * 		Wait() / Broadcast() / Signal()
 *
 * atomic 原子操作
 * 		AddInt32()
 *
 * bytePool.Get().(*[]byte) 强制类型转换
 *
 * 死锁 / 活锁 / 饥饿
 * */

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func testWaitGroup() {
	var wg sync.WaitGroup

	seconds := []int{1, 2, 3, 4, 5}
	for i, s := range seconds {
		// 计数加 1
		wg.Add(1)
		go func(i, s int) {
			// 计数减 1
			time.Sleep(time.Second * time.Duration(rand.Intn(5)))
			fmt.Printf("goroutine%d 结束\n", i)
			defer wg.Done()
		}(i, s)
	}

	// 等待执行结束
	wg.Wait()
	fmt.Println("所有 goroutine 执行结束")
}

// Mutex  一个互斥锁只能同时被一个 goroutine 锁定，其它 goroutine 将阻塞直到互斥锁被解锁（重新争抢对互斥锁的锁定）
func testMutex() {
	ch := make(chan struct{}, 2)

	var l sync.Mutex
	go func() {
		l.Lock()
		defer l.Unlock()
		t := rand.Intn(5)
		fmt.Printf("goroutine1: 我锁定了, 会锁定 %ds\n", t)
		time.Sleep(time.Second * time.Duration(t))
		fmt.Println("goroutine1: 我解锁了，你们去抢吧")
		ch <- struct{}{}
	}()

	go func() {
		fmt.Println("groutine2: 等待解锁")
		l.Lock()
		defer l.Unlock()
		t := rand.Intn(5)
		fmt.Printf("goroutine2: 我锁定了, 会锁定 %ds\n", t)
		time.Sleep(time.Second * time.Duration(t))
		fmt.Println("goroutine2: 我解锁了，你们去抢吧")
		ch <- struct{}{}
	}()

	// 等待 goroutine 执行结束
	for i := 0; i < 2; i++ {
		<-ch
	}
}

// RWMutex: 读锁定（RLock） 读解锁（RUnlock） -  写锁定（Lock）写解锁（Unlock）
// 当有一个 goroutine 获得写锁定，其它无论是读锁定还是写锁定都将阻塞直到写解锁
// 当有一个 goroutine 获得读锁定，其它读锁定仍然可以继续；当有一个或任意多个读锁定，写锁定将等待所有读锁定解锁之后才能够进行写锁定。
func testRWMutex() {
	var rwMutex *sync.RWMutex
	rwMutex = new(sync.RWMutex)
	//var rwMutex sync.RWMutex
	var wg sync.WaitGroup

	rFn := func(i int) {
		rwMutex.RLock()
		defer rwMutex.RUnlock()

		t := rand.Intn(5)
		fmt.Printf("%d is reading %d seconds\n", i, t)
		time.Sleep(time.Second * time.Duration(t))
		fmt.Printf("%d read done\n", i)

		wg.Done()
	}

	wFn := func(i int) {
		rwMutex.Lock()
		defer rwMutex.Unlock()

		t := rand.Intn(5)
		fmt.Printf("%d is writing %d seconds\n", i, t)
		time.Sleep(time.Second * time.Duration(t))
		fmt.Printf("%d write done\n", i)

		wg.Done()
	}

	wg.Add(4)

	go rFn(1)
	go rFn(2)
	go wFn(3)
	go wFn(4)

	wg.Wait()
}

// Once 使用 sync.Once 对象可以使得函数多次调用只执行一次
func testOnce() {
	var once sync.Once
	onceFunc := func() {
		fmt.Println("Only all this func once time")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceFunc)
			fmt.Println("in call outer func")
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
		fmt.Println(i)
	}
}

// Pool
func testPool() {
	var bytePool = &sync.Pool{
		New: func() interface{} {
			b := make([]byte, 1024)
			return &b
		},
	}

	a := time.Now().Unix()
	// 不使用对象池
	for i := 0; i < 5000000; i++ {
		obj := make([]byte, 1024)
		// fmt.Printf("address: %p\n", obj)
		_ = len(obj)
	}
	b := time.Now().Unix()

	// 使用对象池
	for i := 0; i < 5000000; i++ {
		obj := bytePool.Get().(*[]byte)
		// fmt.Printf("address: %p\n", obj)
		_ = obj
		bytePool.Put(obj)
	}
	c := time.Now().Unix()
	fmt.Println("without pool ", b-a, "s")
	fmt.Println("with    pool ", c-b, "s")
}

func testAtomic() {
	//增减操作
	var a int32
	fmt.Println("a : ", a)
	//函数名以Add为前缀，加具体类型名
	//参数一，是指针类型
	//参数二，与参数一类型总是相同
	//增操作
	newA := atomic.AddInt32(&a, 3)
	fmt.Println("new_a : ", newA)
	//减操作
	newA = atomic.AddInt32(&a, -2)
	fmt.Println("new_a : ", newA)
}

func testCrond() {
	condition := false // 条件不满足
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	// 让例程去创造条件
	go func() {
		mu.Lock()
		condition = true // 更改条件
		cond.Signal()    // 发送通知：条件已经满足
		mu.Unlock()
	}()
	mu.Lock()
	// 检查条件是否满足，避免虚假通知，同时避免 Signal 提前于 Wait 执行。
	for !condition {
		// 等待条件满足的通知，如果收到虚假通知，则循环继续等待。
		cond.Wait() // 等待时 mu 处于解锁状态，唤醒时重新锁定。
	}
	fmt.Println("条件满足，开始后续动作...")
	mu.Unlock()
}

//SyncTest : sync package
func SyncTest() {
	rand.Seed(time.Now().UnixNano())

	// testWaitGroup()

	// testMutex()

	// testRWMutex()

	//testOnce()

	// testPool()

	testCrond()
}
