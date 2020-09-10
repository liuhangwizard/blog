package VBSession

import "time"

type VBSessionData struct {
	SessionKey string `json:"key"`
	UserId uint32 `json:"userid"`
	UserName string	`json:"username"`
	RealName string `json:"realname"`
	LastDate time.Time	`json:"date"`
}