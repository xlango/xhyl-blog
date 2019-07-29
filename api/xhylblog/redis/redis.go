package redis

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
	"time"
)


var Client *redis.Client

type RedisClient struct {
}

func initRedis(db []int) {
	dbCode := 0
	if len(db) > 0 {
		dbCode = db[0]
	}

	//读取配置文件
	var addr string
	var poolSize int
	var readTimeout int
	var writeTimeout int
	var idleTimeout int //空闲连接释放时间

	addr=beego.AppConfig.String("redis_addr")
	if addr=="" {
		addr="127.0.0.1:6379"
	}
	poolSize,errPL:=beego.AppConfig.Int("pool_size")
	if errPL!=nil {
		poolSize=10
	}
	readTimeout,errRT:=beego.AppConfig.Int("read_timeout")
	if errRT!=nil {
		readTimeout=100
	}
	writeTimeout,errWT:=beego.AppConfig.Int("write_timeout")
	if errWT!=nil {
		writeTimeout=100
	}
	idleTimeout,errIT:=beego.AppConfig.Int("idle_timeout")
	if errIT!=nil {
		idleTimeout=60
	}

	Client = redis.NewClient(&redis.Options{
		Addr:         addr,
		PoolSize:     poolSize,
		ReadTimeout:  time.Millisecond * time.Duration(readTimeout),
		WriteTimeout: time.Millisecond * time.Duration(writeTimeout),
		IdleTimeout:  time.Second * time.Duration(idleTimeout),
		DB:           dbCode,
	})

	_, err := Client.Ping().Result()
	if err != nil {
		logs.Error(err.Error())
		panic("init redis error")
	}
}

func (this *RedisClient) Get(key string, db ...int) (string, bool) {
	initRedis(db)
	r, err := Client.Get(key).Result()
	if err != nil {
		return "", false
	}
	return r, true
}

func (this *RedisClient) SetExpTime(key string, val interface{}, expTime int32, db ...int) {
	initRedis(db)
	Client.Set(key, val, time.Duration(expTime)*time.Second)
}