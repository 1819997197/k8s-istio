package main

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var log *logrus.Logger

func main() {
	log = logrus.New()
	log.Out = getWriter("./log/info.log") //os.Stdout
	log.Formatter = &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"}

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		resp := http_get()
		log.WithFields(logrus.Fields{"method": "/test"}).Info(resp)
		w.Write(resp)
	})
	http.HandleFunc("/will", func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(logrus.Fields{"method": "/will"}).Info("hello will!")
		w.Write([]byte("hello will!"))
	})

	log.Info("listen 8080")
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func http_get() []byte { //生成client 参数为默认
	client := &http.Client{}
	//生成要访问的url
	url := "http://order-svc:50050/test"
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

func getWriter(filename string) io.Writer {
	hook, err := rotatelogs.New(
		filename+".%Y-%m-%d",
		rotatelogs.WithLinkName(filename),         //为最新的日志建立软连接
		rotatelogs.WithMaxAge(time.Hour*24*30),    //保存30天
		rotatelogs.WithRotationTime(24*time.Hour), //切割频率 1分钟
	)
	if err != nil {
		log.Error("日志启动异常")
		panic(err)
	}
	return hook
}
