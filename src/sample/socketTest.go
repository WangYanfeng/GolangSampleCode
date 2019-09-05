package sample

/**
 * ParseCIDR() / ParseIP() / ParseMAC()
 *
 * 1. Serer
 * 		Listen() / ListenTCP() / ListenIP() ...
 * 		Accept() / Addr()
 * 2. client
 * 		Dial(net, "host:port") / DialTCP(net.TCPAddr) / DialIP() / DialUDP()
 *		例子：Dial("ip4:icmp", "www.baidu.com")
 * 3. net.TCPConn
 * 		Read() / Write()
 * 		Close()
 * 		SetDeadline() / SetReadDeadline()
 * */
import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

// HOST : socket host
const HOST = "localhost"

// PORT : socket port
const PORT = 6666
const clientTimeout = 30
const serverTimeout = 60

const serverCmdReady = "READY"
const clientConnect = "CONNECT"
const clientQuit = "QUIT"
const serverQuit = "SERVER_QUIT"

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func checkTimeout(heartbeat <-chan string, timeout chan<- struct{}) {
	for {
		select {
		case hb := <-heartbeat:
			if hb == clientQuit {
				return
			}
		case <-time.After(time.Second * clientTimeout):
			fmt.Printf("/******    client timeout    ******/\n")
			timeout <- struct{}{}
		}
	}
}

func serverSockHandler(conn net.Conn, mgmtChan chan string) error {
	defer conn.Close()
	buf := make([]byte, 1024)

	for {
		if err := conn.SetReadDeadline(time.Now().Add(time.Second * clientTimeout)); err != nil {
			return err
		}
		c, err := conn.Read(buf)
		if err != nil {
			if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
				fmt.Printf("/******    client timeout    ******/\n")
				mgmtChan <- clientQuit
			}
			return err
		}

		conn.Write(bytes.TrimSpace(bytes.ToUpper(buf[:c])))

		if bytes.Equal(bytes.ToUpper(bytes.TrimSpace(buf[:c])), []byte("BYE")) {
			fmt.Printf("/******    client leave    ******/\n")
			mgmtChan <- clientQuit
			break
		}
	}
	return nil
}

func serverSock(mgmtChan chan string) {
	listener, err := net.Listen("tcp", HOST+":"+strconv.Itoa(PORT))
	checkError(err)
	mgmtChan <- serverCmdReady
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		checkError(err)
		mgmtChan <- clientConnect

		// handle new client connection
		go serverSockHandler(conn, mgmtChan)
	}
}

func clientSock() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", HOST+":"+strconv.Itoa(PORT))
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	defer conn.Close()
	fmt.Printf("/******    client connect    ******/\n")

	buf := make([]byte, 1024)
	for {
		inputIO := bufio.NewReader(os.Stdin)
		fmt.Printf("C: ")
		input, err := inputIO.ReadBytes('\n')
		checkError(err)

		//fmt.Printf("%q len:%d\n", input, len(input))  // '\r\n' or '\n'
		if len(bytes.TrimSpace(input)) == 0 {
			continue
		}

		_, err = conn.Write(input)
		checkError(err)

		c, err := conn.Read(buf[0:])
		checkError(err)
		fmt.Printf("Server: %q [%d]\n", bytes.TrimSpace(buf[:c]), c)
		isBye := bytes.Equal(bytes.ToUpper(bytes.TrimSpace(buf[:c])), []byte("BYE"))
		if isBye {
			conn.Close()
			break
		}
	}
}

// SocketTest : is fake
func SocketTest() {
	var timeout *time.Timer
	mgmtChan := make(chan string, 1)
	clientCount := 0

	go serverSock(mgmtChan)

	for {
		cmd := <-mgmtChan
		if cmd == serverCmdReady {
			fmt.Println("Server Ready. Listen on ", HOST+":"+strconv.Itoa(PORT))
			go clientSock()
		} else if cmd == clientConnect {
			clientCount++
			if timeout != nil {
				timeout.Stop()
				timeout = nil
			}
		} else if cmd == clientQuit {
			clientCount--
		} else if cmd == serverQuit {
			break
		}

		if clientCount == 0 && timeout == nil {
			timeout = time.AfterFunc(time.Second*serverTimeout, func() {
				fmt.Printf("/******    server timeout    ******/\n")
				mgmtChan <- serverQuit
			})
		}
	}
}

// FormatIP : use ParseIP
func FormatIP() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	str := os.Args[1]
	ip := net.ParseIP(str)
	if ip == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Printf("The address is: %s\n", ip.String())
	}
	os.Exit(0)
}
