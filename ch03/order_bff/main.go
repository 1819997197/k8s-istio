package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		resp := http_get()
		w.Write(resp)
	})
	http.HandleFunc("/will", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello will!"))
	})

	fmt.Println("listen 8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func http_get() []byte { //生成client 参数为默认
	client := &http.Client{}
	//生成要访问的url
	url := "http://localhost:50050/test"
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("http_get http.NewRequest: ", err)
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(reqest)
	defer response.Body.Close()

	b, _ := ioutil.ReadAll(response.Body)
	fmt.Println("http_get response: ", string(b))
	return b
}
