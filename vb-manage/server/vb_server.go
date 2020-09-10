package VBServer

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"vb-server/utils/config"
)

//type
type JSON=map[string]interface{}

var Main *gin.Engine;

func Cros(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Token, Language, From")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

func InitServer(){
	//set server
	serverInfo:=config.VBGetConfigServer()

	//mode
	Main=gin.Default()

	gin.SetMode(serverInfo.Mode)


	//gzip
	if serverInfo.Gzip=="1"{
		Main.Use(gzip.Gzip(gzip.DefaultCompression))
	}


	//cros
	//Main.Use(Cros)


}

func RunServer(){


	serverInfo:=config.VBGetConfigServer();
	fmt.Println("监听:"+serverInfo.Port)
	err := Main.Run(":"+serverInfo.Port)
	if err != nil {
		panic(err)
	}
}


