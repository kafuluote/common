package template

var (
	ImproveIp2 = `package client

import (
	"{{.Dir}}/conf"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"time"
	{{.Alias}} "{{.Dir}}/proto/{{.Alias}}"
)

type {{title .Alias}}RPCCli struct {
	{{title .Alias}}Con {{.Alias}}.{{title .Alias}}Service
}

func New{{title .Alias}}RPCCli() (t *{{title .Alias}}RPCCli) {
	service := micro.NewService(
		micro.Name("go.micro.cli.{{.Alias}}"),
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
