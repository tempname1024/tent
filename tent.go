package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var path string

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	p := filepath.Join(path, strings.Replace(r.URL.Path, "..", "", -1))
	http.ServeFile(w, r, p)
}

func main() {
	var host string
	var port uint64

	flag.StringVar(&host, "host", "127.0.0.1", "IP address to listen on")
	flag.Uint64Var(&port, "port", 0, "Port to listen on (default random)")
	flag.StringVar(&path, "path", "./", "Absolute or relative path to serve")
	flag.Parse()

	if _, err := os.Stat(path); err != nil {
		panic(err)
	}

	if port == 0 {
		rand.Seed(time.Now().UnixNano())
		for {
			port = uint64(rand.Intn(65535-1024) + 1024)
			conn, _ := net.DialTimeout("tcp", net.JoinHostPort(host,
				strconv.Itoa(int(port))), time.Second)
			if conn != nil {
				conn.Close()
			} else {
				break
			}
		}
	}

	http.HandleFunc("/", RequestHandler)
	fmt.Printf("Listening on %v port %v (http://%v:%v/)\n", host, port, host, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil))
}
