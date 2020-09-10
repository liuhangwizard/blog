package VBMySQL

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"vb-server/utils/config"
)

type VBDataBase=*gorm.DB

// var Main *gorm.DB

var main *gorm.DB


func InitData(){

	//get config
	dbInfo:=config.VBGetConfigDBMysql()

	//db info
	dialect:=dbInfo.Type;
	content:=dbInfo.Username+":"+dbInfo.Password+"@("+dbInfo.Host+":"+dbInfo.Port+")/"+dbInfo.Name+dbInfo.Args;


	//db args
	fmt.Println("数据库[MySQL]参数:")
	fmt.Println(content)


	//open data base
	var err error;
	main,err=gorm.Open(dialect, content)
	if err!= nil{
		panic(err)
	}

	//conn pool
	//main.DB().SetMaxIdleConns(dbInfo.Maxidleconns)
	//main.DB().SetMaxOpenConns(dbInfo.Maxopenconns)

	//db mode
	debugMode:=false;
	if dbInfo.Mode=="debug"{
		debugMode=true
	}
	main.LogMode(debugMode)

}

func Main() *gorm.DB{
	return main.Debug()
}