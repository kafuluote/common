package conf

import (
	"github.com/kafuluote/common/xtime"
)

type CommConf struct {
	Ver     string
	LogPath string
}

// =================================== HTTP ==================================
// HTTPServer http server settings.
type HTTPServer struct {
	Addrs        []string
	MaxListen    int32
	ReadTimeout  xtime.Duration
	WriteTimeout xtime.Duration
}

// HTTPClient http client settings.
type HTTPClient struct {
	Dial      xtime.Duration
	Timeout   xtime.Duration
	KeepAlive xtime.Duration
	Timer     int
}

// MultiHttp outer/inner/local http server settings.
type MultiHTTP struct {
	Outer *HTTPServer
	Inner *HTTPServer
	Local *HTTPServer
}

type Server struct {
	Proto string
	Addr  string
}

type RPCServer struct {
	Proto string
	Addr  string
}

type ConfDiscovery struct {
	Role     string
	Interval xtime.Duration
}

type ServiceDiscoveryServer struct {
	ServiceName string `json:"service_name"`
	WebName     string `json:"web_name"`
	//RPCAddr     string         `json:"rpc_addr"`
	RegistryAddr string `json:"registry_addr"`
	Interval     int64  `json:"interval"`
	TTL          int64  `json:"ttl"`
}

type ServiceDiscoveryClient struct {
	ServiceName string
	EtcdAddr    string
	Balancer    string
}

type Etcd struct {
	Name    string
	Root    string
	Addrs   []string
	Timeout xtime.Duration
}

type Zookeeper struct {
	Root    string
	Addrs   []string
	Timeout xtime.Duration
}

// Redis client settings.
/*
type Redis struct {
	Name         string         `json:"name"`
	Proto        string         `json:"proto"`
	Addr         string         `json:"addr"`
	Active       int            `json:"active"`
	Idle         int            `json:"idle"`
	DialTimeout  xtime.Duration `json:"dial_timeout"`
	ReadTimeout  xtime.Duration `json:"read_timeout"`
	WriteTimeout xtime.Duration `json:"write_timeout"`
	IdleTimeout  xtime.Duration `json:"idle_timeout"`
	Pwd string `json:"pwd"`
	Num int `json:"num"`
}
*/
type Redis struct {
	Name         string `json:"name"`
	Proto        string `json:"proto"`
	Addr         string `json:"addr"`
	Ttl          int32 `json:"ttl"`
	Active       int    `json:"active"`
	Idle         int    `json:"idle"`
	DialTimeout  int64  `json:"dial_timeout"`
	ReadTimeout  int64  `json:"read_timeout"`
	WriteTimeout int64  `json:"write_timeout"`
	IdleTimeout  int64  `json:"idle_timeout"`
	Pwd          string `json:"pwd"`
	Num          int    `json:"num"`
}

// KafkaProducer kafka producer settings.
type KafkaProducer struct {
	Zookeeper *Zookeeper
	Brokers   []string
	Sync      bool // true: sync, false: async
}

// KafkaConsumer kafka client settings.
type KafkaConsumer struct {
	Group     string
	Topics    []string
	Offset    bool // true: new, false: old
	Zookeeper *Zookeeper
}

type MySQL struct {
	Name   string `json:"name"`   // for trace
	DSN    string `json:"dsn"`    // data source name
	Active int    `json:"active"` // pool
	Idle   int    `json:"idle"`   // pool
}

type MongoDB struct {
	Addrs       string
	DB          string
	DialTimeout xtime.Duration
}

type ES struct {
	Addrs string
}

type LogData struct {
	Path  string
	Name  string
	Level string
}

type GameType struct {
	Rtype int32
	Eth   string
	Count int32
}

//短信
type Sms struct {
	Account string
	Pwd string
	SmsUrl  string
	Title string
}

//邮箱
type Email struct {
	AppKey string
	SecretKey string
	Title string
	Subject string
	Alias string
}

//极验
type Gree struct {
	Key string
	Id string
}

//实名认证
type Auth struct {
	Id string
	Key string
}

type Log struct {
	LogPath  string
	LogLevel string
	LogName  string
}