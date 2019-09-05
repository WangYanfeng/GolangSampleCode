package sample

/**
 * 1. 通道  make(chan interface{})
 * 2. 单向通道
 * 		chan<-  只能发送
 * 		<-chan  只能接收
 * 3. close(ch)
 *
 * 4. 当channel 是public时，需要考虑 “读取超时” 问题，防止死锁。
 * 5. 给被关闭通道发送数据将会触发 panic, 从已关闭的通道接收数据时将不会发生阻塞
 *
 * 6. 多核并行化
 *
 * */
import (
	"fmt"
	"time"
)

func channelTimeout() {
	ch := make(chan struct{})

	//fake input
	go func() {
		time.Sleep(time.Second * 13)
		ch <- struct{}{}
	}()

	select {
	case <-ch:
		fmt.Println("read channel success")
	case <-time.After(time.Second * 10):
		fmt.Println("read timeout!")

		// not add default
	}

	fmt.Println("......")
}

//ChannelTest : channel
func ChannelTest() {
	channelTimeout()
}
