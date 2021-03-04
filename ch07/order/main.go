package main

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

var log *logrus.Logger

func main() {
	log = logrus.New()
	log.Out = getWriter("./log/info.log") //os.Stdout
	log.Formatter = &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"}

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		s := "hello world!"
		log.WithFields(logrus.Fields{"method": "/test"}).Info("hello world!")
		w.Write([]byte(s))
	})

	log.Info("listen 50050")
	http.ListenAndServe("0.0.0.0:50050", nil)
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
