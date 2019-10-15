package template

var (
	ImproveConf = `package conf

import (
	"encoding/json"

	"github.com/kafuluote/common/conf"
	"github.com/zouyx/agollo"
)

var Config = NewAppConfig()

type AppConfig struct {
	Registry conf.ServiceDiscoveryServer "json:'registry'"
	Redis    conf.Redis                  "json:'redis'"
	Mysql    conf.MySQL                  "json:'mysql'"
	Trace    string  					 "json:'trace_address'"
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

func LoadConf() {
	agollo.InitCustomConfig(func() (config *agollo.AppConfig, e error) {
		return &agollo.AppConfig{
			AppId:         "10001",
			Cluster:       "default",
			NamespaceName: "TEST1.head",
			Ip:            "192.168.238.128:8080",
		}, nil
	})

	err := agollo.Start()
	if err != nil {
		panic(err.Error())
	}

	d := agollo.GetStringValue("registry", "")

	err = json.Unmarshal([]byte(d), &Config.Registry)
	if err != nil {
		panic(err.Error())
	}
	
	Config.Trace = agollo.GetStringValue("trace_address", "localhost:6831")

	go lookup()
}

func lookup() {
	event := agollo.ListenChangeEvent()
	go func() {
		for e := range event {
			if e.Namespace == "TEST1.head" {

			}
		}
	}()

}

`
)
