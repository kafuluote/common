package template

var (
	ImproveHttp = `package http

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"
	log "github.com/sirupsen/logrus"

	"{{.Dir}}/conf"
	"{{.Dir}}/http/controller"
)

func initRouter() *gin.Engine {
	r := gin.Default()

	new(controller.ActionGroup).Router(r)

	return r
}

func InitHttpServer() {

	service := web.NewService(
		web.Name("go.micro.api.{{.Alias}}"),
		web.Registry(consul.NewRegistry(registry.Addrs(conf.Config.Registry.RegistryAddr))),
	)

	service.Init()

	r := initRouter()

	service.Handle("/", r)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

`


	ImproveController=`package controller

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
