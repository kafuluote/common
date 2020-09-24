package template

var (
	ImproveHttp2 = `package http

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-micro/v2/registry"
	"github.com/opentracing/opentracing-go"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/Allenxuxu/microservices/lib/tracer"
	"github.com/Allenxuxu/microservices/lib/wrapper/tracer/opentracing/gin2micro"
	log "github.com/sirupsen/logrus"

	"{{.Dir}}/conf"
	"{{.Dir}}/http/controller"
)

const name = "go.micro.api.{{.Alias}}"

func initRouter() *gin.Engine {
	r := gin.Default()

	new(controller.ActionGroup).Router(r)

	return r
}

func InitHttpServer() {
	gin2micro.SetSamplingFrequency(50)
	t, io, err := tracer.NewTracer(name, conf.Config.Trace)
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()

	opentracing.SetGlobalTracer(t)


	service := web.NewService(
		web.Name("go.micro.api.{{.Alias}}"),
		web.Registry(consul.NewRegistry(registry.Addrs(conf.Config.Registry.RegistryAddr))),
		web.RegisterTTL(time.Second*15),
		web.RegisterInterval(time.Second*10),
	)

	service.Init()

	r := initRouter()

	service.Handle("/", r)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

`
)
