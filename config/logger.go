package config

import (
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"strings"
	"time"
)

func Logger() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context, log *log.Logger) {
		start := time.Now()

		log.SetPrefix("[" + start.Format("2006-01-02 15:04:05") + "] ")

		addr := req.Header.Get("X-Real-IP")
		if addr == "" {
			addr = req.Header.Get("X-Forwarded-For")
			if addr == "" {
				addr = req.RemoteAddr
			}
		}

		// 过滤掉 assets 下的所有日志输出
		if strings.HasPrefix(req.URL.Path, "/assets") {
			return
		}

		log.Printf("Started %s %s for %s", req.Method, req.URL.Path, addr)

		rw := res.(martini.ResponseWriter)
		c.Next()

		log.Printf("Completed %v %s in %v\n\n", rw.Status(), http.StatusText(rw.Status()), time.Since(start))
	}
}
