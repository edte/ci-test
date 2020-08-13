// @program: ci-test
// @author: edte
// @create: 2020-08-13 20:58
package main

import (
	"net/http"
)

type A struct {
}

func (a A) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello world"))
}

func main() {
	http.ListenAndServe(":8080", A{})
}
