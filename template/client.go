package template

var (
	ImproveClient=`package rpc

import "{{.Dir}}/rpc/client"

var InnerService *RPCClient

type RPCClient struct {
	{{title .Alias}}Con *client.{{title .Alias}}RPCCli
}

func NewRPCClient() (c *RPCClient) {
	return &RPCClient{
		{{title .Alias}}Con:client.New{{title .Alias}}RPCCli(),
	}
}

func InitInnerService() {
	InnerService = NewRPCClient()
}`

	ImproveIp=`package client

import (
	"{{.Dir}}/conf"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"time"
	{{.Alias}} "{{.Dir}}/proto/{{.Alias}}"
)

type {{title .Alias}}RPCCli struct {
	{{title .Alias}}Con {{.Alias}}.{{title .Alias}}Service
}

func New{{title .Alias}}RPCCli() (t *{{title .Alias}}RPCCli) {
	service := micro.NewService(
		micro.Name("go.micro.srv.{{.Alias}}"),
		micro.RegisterTTL(time.Duration(conf.Config.Registry.TTL)),
		micro.RegisterInterval(time.Duration(conf.Config.Registry.Interval)),
		micro.Registry(consul.NewRegistry(registry.Addrs(conf.Config.Registry.RegistryAddr))),
	)

	service.Init()
	return 	&{{title .Alias}}RPCCli{
		{{title .Alias}}Con: {{.Alias}}.New{{title .Alias}}Service("go.micro.srv.{{.Alias}}",service.Client()),
	}
}
`
)
