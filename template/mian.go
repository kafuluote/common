package template

var (
	ImproveMain = `package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/kafuluote/common/xlog"
	"{{.Dir}}/conf"
	"{{.Dir}}/dao"
	"{{.Dir}}/http"
	"{{.Dir}}/rpc"

	log "github.com/sirupsen/logrus"
)

func init() {
	conf.LoadConf()

	path := os.Getenv("MICRO_LOG_PATH")
	if len(path) == 0 {
		panic("please config MICRO_LOG_PATH")
		return
	}
	level := os.Getenv("MICRO_LOG_LEVEL")
	if len(level) == 0 {
		panic("please config MICRO_LOG_LEVEL")
		return
	}

	xlog.InitLogger(path, "{{.Alias}}.log", level)
}

func main() {
	go dao.InitDao()
	go http.InitHttpServer()
	go rpc.RPCServerInit()

	quitChan := make(chan os.Signal)
	signal.Notify(quitChan,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	sig := <-quitChan
	log.Infof("server close by sig %s", sig.String())
}
`
)
