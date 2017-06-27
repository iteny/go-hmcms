package main

import (
	"flag"
	"go-hmcms/controllers/intendant"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/iteny/hmgo/router"
)

var dir string
var port int
var staticHandler http.Handler

// 初始化参数
func init() {
	dir = path.Dir(os.Args[0])
	flag.IntVar(&port, "port", 80, "服务器端口")
	flag.Parse()
	staticHandler = http.FileServer(http.Dir(dir))
}

func main() {
	router := router.New()
	router.GET("/admin", intendant.Login)
	router.POST("/admin/login", intendant.LoginGo)
	// router.NotFound = router.GET("/error", intendant.Error)
	router.GET("/static/*filepath", StaticServer)

	s := &http.Server{
		Addr:           ":" + strconv.Itoa(port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
func StaticServer(w http.ResponseWriter, req *http.Request, ps router.Params) {
	if req.URL.Path != "/" {
		staticHandler.ServeHTTP(w, req)
		return
	}
	io.WriteString(w, "hello, world!\n")
}
