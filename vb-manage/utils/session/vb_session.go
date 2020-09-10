package VBSession

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/satori/go.uuid"
	VBRedis "vb-server/database/redis"
	VBLog "vb-server/utils/log"
)

type VBSession struct {

}

func (s *VBSession) NewId() string{
	id:=uuid.NewV4()
	return id.String()
}

func (s *VBSession) GetSession(sessionId string) (VBSessionData,error) {

	//pool
	conn:=VBRedis.Get()
	defer conn.Close()

	sessionData:=VBSessionData{}
	data,err:=redis.String(conn.Do("GET",sessionId))
	if err!=nil{
		//data为空
		return sessionData,nil
	}

	err=json.Unmarshal([]byte(data),&sessionData)
	if err!=nil{
		VBLog.Print.Info("[session string转json出错]")
	}

	return sessionData,nil
}

func (s *VBSession) SetSession(sessionId string,sessionData VBSessionData) error {

	//pool
	conn:=VBRedis.Get()
	defer conn.Close()


	data,err:=json.Marshal(sessionData)
	if err!=nil{
		VBLog.Print.Info("[session json转string出错]")
	}

	_,err=conn.Do("SET",sessionId,string(data))
	if err!=nil{
		VBLog.Print.Info("[session 将session设置到redis失败]")
	}

	return nil
}

func (s *VBSession) UpdateSessionTime (sessionId string,second int) error {

	//pool
	conn:=VBRedis.Get()
	defer conn.Close()


	_,err:=conn.Do("EXPIRE",sessionId,second)
	if err!=nil{
		VBLog.Print.Info("[session 设置key时间失败] "+string(second))
	}

	return nil
}


func (s *VBSession) RemoveSession(sessionId string) error {

	//pool
	conn:=VBRedis.Get()
	defer conn.Close()


	_,err:=conn.Do("DEL",sessionId)
	if err!=nil{
		VBLog.Print.Info("[session 移除session失败] "+sessionId)
	}

	return nil
}






