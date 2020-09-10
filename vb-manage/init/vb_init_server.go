package init

import (
	VBRouter "vb-server/router"
	VBServer "vb-server/server"
)

func initServer(){
	//读取错误会panic
	VBServer.InitServer()
	VBRouter.InitRouter()

}
