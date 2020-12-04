package template

var (
	ImproveMain = `package main

import (
	"os"
	"os/signal"
	"syscall"
	"context"
	"sync"

	"{{.Dir}}/conf"
	"{{.Dir}}/dao"
	"{{.Dir}}/http"
	"{{.Dir}}/rpc"

	log "github.com/sirupsen/logrus"
)


func init() {

	conf.LoadConf()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	waiter := &sync.WaitGroup{}
	waiter.Add(1)

	go dao.InitDao()
	go http.InitHttpServer()
	go rpc.RPCServerInit(ctx,waiter)

	quitChan := make(chan os.Signal)
	signal.Notify(quitChan,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	log.Infof("server start end")
	sig := <-quitChan
	cancel()
	waiter.Wait()
	log.Infof("server close by sig %s", sig.String())
}
`
)
