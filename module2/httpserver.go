package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
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

	_, err = io.WriteString(response, "Get headers succeed.")
	if err != nil {
		return
	}
}

func getEnvVersion(response http.ResponseWriter, request *http.Request) {
	envVer := os.Getenv("VERSION")
	fmt.Printf("env version:%s\n", envVer)
	response.Header().Set("VERSION", envVer)
	_, err := io.WriteString(response, "get env version ")
	if err != nil {
		return
	}
}

func serverLog(response http.ResponseWriter, request *http.Request) {
	ipInfo := request.RemoteAddr
	ipStr := strings.Split(ipInfo, ":")
	fmt.Printf("client ip:%s\n", ipStr[0])
	httpCode := strconv.Itoa(http.StatusOK)
	fmt.Printf("http code:%s\n", httpCode)
	_, err := io.WriteString(response, "get client ip and http code")
	if err != nil {
		return
	}
}

func healthZ(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	_, err := io.WriteString(response, "visit healthZ")
	if err != nil {
		return
	}
}

func main() {
	// http://127.0.0.1:81/requestAndResponse
	http.HandleFunc("/requestAndResponse", requestAndResponse)
	// http://127.0.0.1:81/getEnvVersion
	http.HandleFunc("/getEnvVersion", getEnvVersion)
	// http://127.0.0.1:81/serverLog
	http.HandleFunc("/serverLog", serverLog)
	// http://127.0.0.1:81/healthZ
	http.HandleFunc("/healthZ", healthZ)
	err := http.ListenAndServe(":81", nil)
	if nil != err {
		log.Fatal(err) //显示错误日志
	}
}
