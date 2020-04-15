package main

import (
    "net"
    "net/http"
    "net/http/fcgi"
    "os"
    "strconv"
    "time"
)

type FastCGIServer struct{}

// 暂停1s， 打印标识的进程id
func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
    time.Sleep(time.Second)
    resp.Write([]byte("ProcessId: " + strconv.Itoa(os.Getpid()) + "\n"))
}

func main() {
    listener, _ := net.Listen("tcp", "127.0.0.1:9001")
    srv := new(FastCGIServer)
    fcgi.Serve(listener, srv)
}
