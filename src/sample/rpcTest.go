package sample

/**
* RPC 远程过程调用，C/S架构。RPC协议基于TCP UDP或者HTTP上，开发者调用另一台计算机上的程序，而无须了解底层网络通信应用程序协议细节。
*
* RPC 客户端程序，通过网络或者其他I/O连接调用一个远程对象的公开方法（大写字母开头）
* RPC 服务端，可将一个对象注册为可访问的服务。服务进程保持睡眠状态知道客户端的调用请求。
*
* NOTICE: 一个RPC服务端可以注册多个不同类型的对象，但是不允许注册同一类型的多个对象。必须满足：
*    =>  func (t *T) MethodName(argType T1, replyType *T2) error
* 1. 方法首字母大写
* 2. 两个参数， T1 和 T2 默认会使用 Go 内置的 encoding/gob 包进行编码解码
* 3. 第一个参数表示由 RPC 客户端传入的参数，第二个参数表示要返回给 RPC 客户端的结果
* 4. 该方法最后返回一个 error 类型的值
*
*
* Server
*		rpc.ServeConn
*
* Client
* 		1. Call() 同步处理
* 		2. Go() 异步处理
* Call()/Go() 都必须指定要调用的服务及其方法名称，以及一个客户端传入参数的引用，还有一个用于接收处理结果参数的指针
* 如果没有明确指定 RPC 传输过程中使用何种编码“解码器”，默认将使用 Go 标准库提供的 encoding/gob 包进行数据传输
*
*
* rpc.Register()
* rpc.HandleHTTP()
*
* rpc.DialHTTP()
*
* Gob 是Go语言自己的以二进制形式序列化和反序列化数据格式。用于rpc、以及应用程序返回的数据通信。
* */

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"sync"
	"time"
)

// Args :参数结构
type Args struct {
	A, B int
}

// Quotient :返回值结构
type Quotient struct {
	Quo, Rem int
}

// Arith : RPC一个类型
type Arith int

// Multiply : RPC call, Arith:Multiply
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Divide : RPC call, Arith:Divide
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func rpcServer() {
	arith := new(Arith) // Arith类型的对象

	rpc.Register(arith) // 注册RPC
	rpc.HandleHTTP()    // rpc 作为 http 的handler

	l, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
	fmt.Println("RPC server listen on tcp:1234")
}

//--------------------------------------

func rpcClient() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}

// RPCTest : test rpc package. fake in one machine
func RPCTest() {
	var wg sync.WaitGroup
	wg.Add(2)

	rpcServer()
	time.Sleep(time.Second * 3)
	go rpcClient()

	wg.Wait()
}
