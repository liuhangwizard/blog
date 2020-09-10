package model

import (
	"time"
	. "vb-server/database/mysql"
)

type VBModelPermission struct {
	Id         uint32    `json:"id"`
	UserId     string    `json:"userid"`
	Config     string    `json:"config"`
	UpdateDate time.Time `json:"updatedate"`
}

func NewModelPermission() *VBModelPermission {
	return &VBModelPermission{}
}

func (m *VBModelPermission) TableName() string {
	return "vb_permission"
}

func (m *VBModelPermission) DataGetUserPermission(userid uint32) (*VBModelPermission, error) {

	permission := NewModelPermission()
	err := Main().Where(map[string]interface{}{
		"user_id": userid,
	}).Find(&permission).Error

	return permission, err
}
