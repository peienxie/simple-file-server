package main

import (
	"fmt"
	"log"
	"net/http"

	flag "github.com/spf13/pflag"
)

var port = flag.IntP("port", "p", 8080, "listening port")
var root = flag.StringP("root", "r", "./", "root directory of server.")

func main() {
	flag.Parse()

	addr := fmt.Sprintf(":%d", *port)
	handler := http.FileServer(http.Dir(*root))

	log.Printf("HTTP server is listening on port: %d", *port)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal("listen err:", err)
	}
}
