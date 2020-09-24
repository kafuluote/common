package template

var (
	ImproveRpcServer2 = `package rpc

import (
	"fmt"
	"time"

	"{{.Dir}}/conf"
	"{{.Dir}}/rpc/handler"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/broker/nsq/v2"
	"github.com/micro/go-plugins/registry/consul/v2"

	"github.com/Allenxuxu/microservices/lib/tracer"

	proto "{{.Dir}}/proto/{{.Alias}}"
	log "github.com/sirupsen/logrus"

	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	opentracing "github.com/opentracing/opentracing-go"
)

const name = "go.micro.srv.hello"

func RPCServerInit() {
	t, io, err := tracer.NewTracer(name, conf.Config.Trace)
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)


	service := micro.NewService(
		micro.Name("go.micro.srv.{{.Alias}}"),
		micro.RegisterTTL(time.Duration(conf.Config.Registry.TTL)),
		micro.RegisterInterval(time.Duration(conf.Config.Registry.Interval)),
		micro.Registry(consul.NewRegistry(registry.Addrs(conf.Config.Registry.RegistryAddr))),
		micro.Version("latest"),
		micro.Broker(nsq.NewBroker(broker.Addrs("my-nsq:4150"))),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	service.Init()

	proto.Register{{title .Alias}}Handler(service.Server(), new(handler.{{title .Alias}}))

	if err := service.Run(); err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}

}

`

	ImproveRpcHandler2 = `package handler

import (
	"context"

	log "github.com/sirupsen/logrus"

	{{.Alias}} "{{.Dir}}/proto/{{.Alias}}"
)

type {{title .Alias}} struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *{{title .Alias}}) Call(ctx context.Context, req *{{.Alias}}.Request, rsp *{{.Alias}}.Response) error {
	log.Info("Received {{title .Alias}}.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *{{title .Alias}}) Stream(ctx context.Context, req *{{.Alias}}.StreamingRequest, stream {{.Alias}}.{{title .Alias}}_StreamStream) error {
	log.Infof("Received {{title .Alias}}.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&{{.Alias}}.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *{{title .Alias}}) PingPong(ctx context.Context, stream {{.Alias}}.{{title .Alias}}_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&{{.Alias}}.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
`
)
