package VBRedis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
	"vb-server/utils/config"
)

var (
	Pool *redis.Pool
)

func InitData(){

	//redis config
	info:=config.VBGetConfigDBRedis()
	infoDail:=info.Host+":"+info.Port
	infoTimeOut:=time.Duration(info.TimeOut)*time.Second

	//redis pool
	Pool = &redis.Pool{
		MaxIdle:     info.PoolSize,
		MaxActive:   info.PoolActive,
		IdleTimeout: infoTimeOut,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			//connet
			conn, err := redis.Dial("tcp", infoDail,
				redis.DialPassword(info.Password),
				redis.DialDatabase(info.DB),
				redis.DialConnectTimeout(infoTimeOut),
				redis.DialReadTimeout(infoTimeOut),
				redis.DialWriteTimeout(infoTimeOut))
			if err != nil {
				return nil, err
			}

			//select DB
			_,err=conn.Do("SELECT",info.DB)
			if err!=nil{
				return nil, err
			}

			return conn, nil
		},
	}
	//db log
	fmt.Println("数据库[Redis]参数:")
	fmt.Println(infoDail)
}

func Get() redis.Conn {
	return Pool.Get()
}