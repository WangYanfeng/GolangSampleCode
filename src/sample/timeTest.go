package sample

/**
 * time.Now() 返回Time类型 t
 * 		t.Unix() 时间->时间戳
 * 		t.Year() / t.Month() ...
 * 		t.Format() 格式化字符串
 * 		t.Sub(t)
 * 		t.Add(duration)
 *
 * time.Unix() 时间戳 -> 时间
 * time.Parse()
 * time.ParseInLocation()
 * time.LoadLocation() 设置时区
 *
 * time.Duration 时间间隔类型，包含time.Second / time.Minute / time.Hour
 * 		d.Seconds() / d.Minutes()
 * time.Duration()
 * time.Sleep()
 *
 * time.After(t Duration)  多少时间之后，但在取出管道内容前不阻塞
 * time.AfterFunc(time.Duration, func())  多少时间之后在goroutine line执行函数,不阻塞
 *
 * time.NewTimer() 定时器timer
 * 		timer.C 通道
 * 		timer.Reset()
 * 		timer.Stop()
 *
 * time.Tick(time.Duration) 每隔多少时间后，其他与After一致, 不阻塞
 * 		返回一个time.C管道，每隔1秒(time.Second)后会在此管道中放入一个时间点. 时间点记录的是放入管道那一刻的时间
 *
 * */
import (
	"fmt"
	"time"
)

func blockSleep() {
	fmt.Println("start block sleeping 2 seconds...")
	time.Sleep(time.Second * 2)
	fmt.Println("end block sleep.")
}

//time.After(time.Duration) 多少时间之后，但在取出管道内容前不阻塞
func noBlockSleep() {
	fmt.Println("in not block sleep. 2 second")
	//返回一个time.C这个管道，2秒后会在此管道中放入一个时间点(time.Now())
	timeChannel := time.After(time.Second * 2)

	fmt.Println("not block")
	fmt.Println("not block")
	fmt.Println("not block")
	//阻塞中，直到取出tc管道里的数据
	fmt.Println("block in channel read")
	<-timeChannel
	fmt.Println("the end")
}

//time.AfterFunc(time.Duration,func())  多少时间之后在goroutine line执行函数,不阻塞
func timeRun() {
	f := func() {
		fmt.Println("Timeout")
	}
	fmt.Println("start ....")
	time.AfterFunc(2*time.Second, f)
	time.Sleep(4 * time.Second)
}

//time.Tick(time.Duration) 每隔多少时间后，其他与After一致, 不阻塞
func tick() {
	timeChannel := time.Tick(time.Second) //返回一个time.C这个管道，每隔1秒(time.Second)后会在此管道中放入一个时间点. 时间点记录的是放入管道那一刻的时间
	for i := 1; i <= 5; i++ {
		t := <-timeChannel
		fmt.Println("get tick from channel: ", t)
	}
}

func timerEvent() {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("Timer has expired.")
	}()

	fmt.Println("Timer reset")
	timer.Reset(1 * time.Second)

	time.Sleep(10 * time.Second)
}

// TimeTest : time pkg
func TimeTest() {
	// timerEvent()
	// tick()
	// timeRun()
	// blockSleep()
	noBlockSleep()
}
