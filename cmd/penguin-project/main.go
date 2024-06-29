package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

const (
	defaultAddress string = "127.0.0.1"
	defaultPort    string = "80"
)

var (
	address string
	port    string
)

func main() {
	flag.StringVar(&address, "addr", defaultAddress, "Server address")
	flag.StringVar(&port, "port", defaultPort, "Server port")
	flag.Parse()

	fmt.Printf("Server running on %v port %v", address, port)
	runAPI()
}

func runAPI() {
	fullAddress := address + ":" + port
	http.HandleFunc("/test", TestHandle)
	http.ListenAndServe(fullAddress, nil)
}

func TestHandle(wr http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(wr, "Hello Világ\n")
	fmt.Fprintf(wr, "Az összeg %v\n", getSum(req))
}

func getSum(req *http.Request) int {
	params := req.URL.Query()
	a := params.Get("a")
	b := params.Get("b")

	aNumber, _ := strconv.Atoi(a)
	bNumber, _ := strconv.Atoi(b)

	return (aNumber + bNumber)
}
