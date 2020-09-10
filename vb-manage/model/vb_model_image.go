package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
	. "vb-server/database/mysql"
	"vb-server/request"
	"vb-server/response"
)

type VBModelImage struct {
	Id         uint32    `json:"id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Src        string    `json:"src"`
	CreateDate time.Time `json:"createdate"`
	UpdateDate time.Time `json:"updatedate"`
}

func NewModelImage() *VBModelImage {
	return &VBModelImage{}
}

func (m *VBModelImage) TableName() string {
	return "vb_image"
}

func (m *VBModelImage) DataNewImage(info *request.VBRImageNew) error {

	//oss上传成功 准备插入图片
	return Main().Transaction(func(tx *gorm.DB) error {

		//image
		image := VBModelImage{
			Name:       info.Name,
			Type:       info.Type,
			Src:        info.Src,
			CreateDate: time.Now(),
			UpdateDate: time.Now(),
		}

		err := tx.Create(&image).Error
		return err

	})
}

func (m *VBModelImage) DataQueryParam(param *request.VBRImageQueryParam) ([]response.VBResultImageQueryParam,uint,error) {

	//oss上传成功 准备插入图片
	data := Main().Table("vb_image")

	//select
	data = data.Select([]string{
		"id",
		"name",
		"type",
		"src",
		"create_date",
	})

	if param.Type != "" {
		data = data.Where("type = ?", param.Type)
	}

	if param.Date != "" {
		data = data.Where("create_date <= ?", param.Date)
	}

	if param.Name != "" {
		data = data.Where("name LIKE ?", param.Name+"%")
	}

	//page
	page := param.Page
	pageSize := param.PageSize

	if page > 0 && pageSize > 0 {
		data = data.Limit(pageSize).Offset((page - 1) * pageSize)
	}

	//result
	var count uint
	result := []response.VBResultImageQueryParam{}
	err := data.Find(&result).Error
	errCount := data.Limit(-1).Offset(-1).Count(&count).Error
	if errCount != nil {
		return []response.VBResultImageQueryParam{}, 0, errCount
	}

	return result,count,err

}

func (m *VBModelImage) DataUpateType(info *request.VBRImageUpdateType) error{
	//开始事务
	return Main().Transaction(func(tx *gorm.DB) error {

		if count := tx.Table("vb_image").Select("type").Where("id = ?", info.Id).Update(map[string]interface{}{"type": info.Type}).RowsAffected; count == 0 {
			return errors.New("update status fail")
		}

		return nil
	})
}