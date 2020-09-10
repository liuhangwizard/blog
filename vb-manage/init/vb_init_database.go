package init

import (
	VBMySQL "vb-server/database/mysql"
	VBRedis "vb-server/database/redis"
	VBOSS "vb-server/database/oss"

)

func initDatabase(){
	//读取错误会panic
	//mysql
	VBMySQL.InitData()

	//oss
	VBOSS.InitData()

	//redis
	VBRedis.InitData()
}
