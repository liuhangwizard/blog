package model

import (
	"github.com/jinzhu/gorm"
	"time"
	. "vb-server/database/mysql"
	"vb-server/request"
)

type VBModelCatagory struct {
	Id         uint32    `json:"id"`
	Content    string    `json:"content"`
	CreateDate time.Time `json:"-"`
}

func NewModelCatagory() *VBModelCatagory {
	return &VBModelCatagory{}
}

func (m *VBModelCatagory) TableName() string {
	return "vb_catagory"
}

func (m *VBModelCatagory) DataGetCatagoryList() ([]VBModelCatagory, error) {

	catagoryList := []VBModelCatagory{}
	err := Main().Find(&catagoryList).Error
	return catagoryList, err
}

func (m *VBModelCatagory) New(info *request.VBRCatagoryNew) error {

	return Main().Transaction(func(tx *gorm.DB) error {

		catagory := VBModelCatagory{
			Content:    info.Content,
			CreateDate: time.Now(),
		}

		err := tx.Create(&catagory).Error
		return err
	})
}

func (m *VBModelCatagory) Delete(info *request.VBRCatagoryDelete) error {

	return Main().Transaction(func(tx *gorm.DB) error {

		if err := tx.Where("id = ?", info.Id).Delete(VBModelCatagory{}).Error; err!=nil{
			return err
		}

		return nil
	})
}
