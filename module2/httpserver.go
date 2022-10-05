package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func requestAndResponse(response http.ResponseWriter, request *http.Request) {
	headers := request.Header
	for header := range headers {
		values := headers[header]
		for index := range values {
			// strings.TrimSpace(s string)会返回一个string类型的slice，将最前面和最后面的ASCII定义的空格去掉，遇到了\0等其他字符会认为是非空格
			values[index] = strings.TrimSpace(values[index])
		}
		println(header + "=" + strings.Join(values, ","))
		response.Header().Set(header, strings.Join(values, ","))
	}
	_, err := fmt.Fprintln(response, "header的数据:", headers)
	if err != nil {
		return
	}

	_, err = io.WriteString(response, "succeed")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/requestAndResponse", requestAndResponse)
	// http://127.0.0.1:81/requestAndResponse
	err := http.ListenAndServe(":81", nil)
	if nil != err {
		log.Fatal(err) //显示错误日志
	}
}
