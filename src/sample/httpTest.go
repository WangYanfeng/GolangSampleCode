package sample

/** net/http package
 * 1 http.    客户端                              // 是http.DefaultClient基础上调用
 * 		Get() / Post() / PostForm() / Head()
 * 		NewRequest()
 *
 * 2. http.Response类
 * 		resp.Body.close()
 * 		resp.StatusCode
 *
 * 3. http.Header
 * 		header.Add()
 *
 * 4. http.    服务端                             // 一切的基础：ServeMux和 Handler。
 *                                               // 1. ServeMux是多路复用器。1. 默认路由（router） 即 Multiplexer器: http.DefaultServeMux。 2. 也可以自己构建
 * 		ListenAndServe(string, handler)          // 2.  handler 是处理器。对象只要有ServeHTTP（）方法（即，满足http.Handler接口）都可以作为处理器。 ServeMux也实现了ServeHttp，所以也可以做参数。
 * 		Handle() / HandleFunc()
 *
 * 		http.NewServeMux()
 * 		http.FileServer() / http.RedirectHandler() / NoFoundHandler()
 *
 * 5. TLS
 * 		http.ListenAndServeTLS()
 *
 * --------------- 高级封装 -------------------------------
 * 1. http.Client类
 * 		Get() / Post() / PostForm() / Head()
 * 		Do()
 * 		实现 Transport 接口的 RoundTripper 变量
 * 2. http.Sever类
 * 		{
 * 			Addr: ":8080",
 * 			Handler: "",
 * 			ReadTimeout: 10*time.Second,
 * 			WriteTimeout: 10*time.Second,
 * 			MaxHeaderBytes: 1<<20,
 * 		}
 *
 * 3. http.Transport       // “传输层” 为 “业务层” 屏蔽了细节：底层传输、代理、gzip、连接池管理、ssl认证等
 *
 * 4. 使用中间件
 * */

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"time"
)

func httpServer() {
	// 注册路由，多路复用器（Multiplexor）上将路径（url）关联到处理器（Handler）上
	// func是ServeHttp()
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello")) //fmt.Fprintln(w, "Hello World")
	})
	http.ListenAndServe(":8001", nil)
	fmt.Println("Server listen on :8001")

	/*
		mux := http.NewServeMux()

		rh := http.RedirectHandler("http://10.206.66.73/", 307)
		mux.Handle("/hello", rh)
		http.ListenAndServe(":8001", mux)
	*/
}

// 自定义Handler
type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	tm := time.Now().Format(th.format)
	
	w.Header().Set("name","For Test")
	w.WriteHeader(500)

	w.Write([]byte("The time is: " + tm))
}

func httpServer2() {
	mux := http.NewServeMux()

	th := &timeHandler{format: time.RFC1123}
	mux.Handle("/hello", th)

	http.ListenAndServe(":8001", mux)
	fmt.Println("Server 2 listen on :8001")
}

//中间件
func hello(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello"))
}
func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		// next handler
		next.ServeHTTP(wr, r)
		timeElapsed := time.Since(timeStart)
		fmt.Println(timeElapsed)
	})
}

func httpServer3() {
	hdl := http.HandlerFunc(hello)
	http.Handle("/hello", timeMiddleware(hdl))
	http.ListenAndServe(":8001", nil)
	fmt.Println("Server 3 listen on :8001")
}

func httpClient() {
	resp, err := http.Get("http://localhost:8001/hello")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	headers := resp.Header // resp.Proto
	for k, v := range headers {
		fmt.Printf("\t Header %v, v=%v\n", k, v)
	}

	fmt.Printf("resp content length: %d\n", resp.ContentLength)

	fmt.Printf("resp uncompressed: %t\n", resp.Uncompressed)
	fmt.Println(reflect.TypeOf(resp.Body))

	//io.Copy(os.Stdout, resp.Body)
	buf := bytes.NewBuffer(make([]byte, 0, 512))
	length, _ := buf.ReadFrom(resp.Body)

	fmt.Printf("Read buf : %d\n", length) // len(buf.Bytes()
	fmt.Println("content: ", string(buf.Bytes()))
}

func httpClient2() {
	// http.PostForm("url", url.Value{"":"","":""})
	// http.Post(url, content-type, io.Reader)
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://10.206.66.73/module/doLogin.php", strings.NewReader("username=admin@edwin.com&password=111111"))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", string(content))

}

// HTTPTest : use net/http package sample
func HTTPTest() {
	// httpClient2()

	var wg sync.WaitGroup
	wg.Add(2)

	// go httpServer()
	go httpServer2()
	time.Sleep(time.Second * 10)
	go httpClient()

	wg.Wait()
}
