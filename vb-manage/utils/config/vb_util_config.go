package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	fmt.Println("初始化配置中...")

	//start read config

	//prod
	//viper.SetConfigName("voluteblog.conf")

	//dev
	viper.SetConfigName("voluteblog-dev.conf")

	viper.AddConfigPath("./conf/")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath("./")
	viper.AddConfigPath("/home/server/main/conf/")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		panic("Viper 配置出错")
	}
}

//config
//server
func VBGetConfigServer() VBConfigServer {
	serverConfig := VBConfigServer{}
	if err := viper.UnmarshalKey("server", &serverConfig); err != nil {
		fmt.Println("解析服务器配置出错:")
		panic("Server 配置出错")
	}
	return serverConfig
}


//log
func VBGetConfigLog() VBConfigLog {
	logConfig := VBConfigLog{}
	if err := viper.UnmarshalKey("log", &logConfig); err != nil {
		panic("Log 配置出错")
	}
	return logConfig
}

//database
func VBGetConfigDBMysql() VBConfigDBMySQL {

	mysqlConfig := VBConfigDBMySQL{}
	if err := viper.UnmarshalKey("database", &mysqlConfig); err != nil {
		fmt.Println("解析MySQL数据库配置出错:")
		panic("MySQL 配置出错")
	}
	return mysqlConfig
}

func VBGetConfigDBRedis() VBConfigDBRedis {

	redisConfig := VBConfigDBRedis{}
	if err := viper.UnmarshalKey("redis", &redisConfig); err != nil {
		fmt.Println("解析redis内存数据库配置出错:")
		panic("redis 配置出错")
	}
	return redisConfig
}

func VBGetConfigDBOSS() VBConfigDBOSS {

	ossConfig:=VBConfigDBOSS{}
	if err:=viper.UnmarshalKey("oss",&ossConfig); err!=nil{
		fmt.Println("解析oss配置出错:")
		panic("oss 配置出错")
	}

	return ossConfig;
}