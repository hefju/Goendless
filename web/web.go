//本想着找个golang的gui写小程序的, 但golang的gui基本上都关掉.
//还是用html+css吧. 以前运行golang的httpserver,然后打开浏览器输入ip和端口,
//我想更进一步,双击就打开可操作的界面. 问题是无法关闭啊,关掉浏览器也不会关闭httpserver
package web

import (
	"log"
	"net/http"
	"github.com/hefju/Goendless/explorer"
//	"fmt"
//	"net"
//	"time"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello, world!"))
}
func CloseServer(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello, world!"))
}

func Run() {


//	server := &http.Server{
//		Addr:         ":12345",
//		ReadTimeout:  10 * time.Second,
//		WriteTimeout: 10 * time.Second,
//		Handler:      nil,
//	}
//	ln, err := net.Listen("tcp", ":12345")
//	if err != nil {
//		log.Fatal("ListenAndServe: ", err)
//		//panic("listen:12345 fail:"+err.Error())
//	}
//	if err = server.Serve(ln); err != nil {
//		fmt.Println("start http server fail:", err)
//	}
//	ln.Close()

		explorer.OpenExplorer("127.0.0.1:12345/hello")
		http.HandleFunc("/hello", HelloServer)
		err := http.ListenAndServe(":12345", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}

}
