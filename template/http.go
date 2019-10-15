package template

var (
	ImproveHttp = `package http

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-micro/registry"
	"github.com/opentracing/opentracing-go"
	"github.com/micro/go-micro/registry/consul"
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

	ImproveController = `package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kafuluote/common/errs"
)

type ActionGroup struct{}

func (s *ActionGroup) Router(r *gin.Engine) {
	r.POST("/hello", s.Hello)
}

func (s *ActionGroup) Hello(c *gin.Context) {
	ret := errs.NewPublciError()
	defer func() {
		c.JSON(http.StatusOK, ret.GetResult())
	}()

	param := &struct {
		Name string
	}{}

	if err := c.ShouldBind(&param); err != nil {
		ret.SetErrCode(errs.ERRCODE_PARAM, err.Error())
		return
	}
	ret.SetDataValue(fmt.Sprintf("well %s", param.Name))
}

`
)
