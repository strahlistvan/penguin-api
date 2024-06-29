package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Server running on localhost port 8082")
	runAPI()
}

func runAPI() {
	http.HandleFunc("/test", TestHandle)
	http.ListenAndServe(":8082", nil)
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
