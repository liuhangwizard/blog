package model

import (
	"encoding/json"
	"time"
	. "vb-server/database/mysql"
)

type VBModelAdmin struct {
	Id         uint32          `json:"id"`
	Username   string          `json:"username"`
	Password   string          `json:"password"`
	Realname   string          `json:"realname"`
	Gender     int             `json:"gender"`
	Level      int             `json:"level"`
	Info       json.RawMessage `json:"info"`
	CreateDate time.Time       `json:"createdate"`
}

func NewModelAdmin() *VBModelAdmin {
	return &VBModelAdmin{}
}

func (m *VBModelAdmin) TableName() string {
	return "vb_admin"
}

func (m *VBModelAdmin) DataGetAdmin(username string, password string) (*VBModelAdmin, error) {

	admin := NewModelAdmin()

	err := Main().Where(map[string]interface{}{
		"username": username,
		"password": password,
	}).Find(&admin).Error

	return admin, err
}
