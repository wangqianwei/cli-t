package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Print(1)

	resp, err := http.Get("http://fanyi.youdao.com/openapi.do?keyfrom=<keyfrom>&key=G0gXmrRh2SlUcJf5w8L9uYaRyU51Ij6w&type=data&doctype=json&version=1.1&q=good")
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))
}
