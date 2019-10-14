package template

var (
	ImproveDao = `package dao

var DB *Dao

type Dao struct {
	redis *RedisCli
	mysql *Mysql
}

func NewDao() (dao *Dao) {
	return &Dao{
		redis: NewRedisCli(),
		mysql: NewMysql(),
	}
}

func InitDao() {
	DB = NewDao()
}
`

	ImproveDaoMysql = `package dao

import (
	"github.com/go-xorm/xorm"
	"{{.Dir}}/conf"
	log "github.com/sirupsen/logrus"
)

type Mysql struct {
	im *xorm.Engine
}

func NewMysql() (mysql *Mysql) {

	engine, err := xorm.NewEngine("mysql", conf.Config.Mysql.DSN)
	if err != nil {
		log.Fatalf("db err is %s", err)
	}
	engine.ShowSQL(true)

	err = engine.Ping()
	if err != nil {
		log.Fatalf("db err is %s", err)
	}

	mysql = &Mysql{
		im: engine,
	}
	return mysql
}

func (s *Dao) GetMysqlConn() *xorm.Engine {
	return s.mysql.im
}
`
	ImproveDaoRedis = `package dao

import (
	"github.com/go-redis/redis"
	"{{.Dir}}/conf"
	log "github.com/sirupsen/logrus"
)

type RedisCli struct {
	con *redis.Client
}

func NewRedisCli() *RedisCli {

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Config.Redis.Addr,
		Password: conf.Config.Redis.Pwd, // no password set
		DB:       0,                     // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("redis connect faild ")
	}

	log.Infoln(pong)

	return &RedisCli{
		con: client,
	}
}

func (s *Dao) GetRedisConn() *redis.Client {
	return s.redis.con
}

`
)
