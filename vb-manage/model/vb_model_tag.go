package model

import (
	"github.com/jinzhu/gorm"
	"time"
	. "vb-server/database/mysql"
	"vb-server/request"
	. "vb-server/response"
)

type VBModelTag struct {
	Id         uint32    `json:"id"`
	Name       string    `json:"name"`
	CreateDate time.Time `json:"-"`
}

func NewModelTag() *VBModelTag {
	return &VBModelTag{}
}

func (m *VBModelTag) TableName() string {
	return "vb_tag"
}

func (m *VBModelTag) DataGetTagList() ([]VBModelTag, error) {

	tagList := []VBModelTag{}
	err := Main().Find(&tagList).Error
	return tagList, err
}

func (m *VBModelTag) DataTagBlogList(info *request.VBRTagBlogList) ([]VBResultTagBlogList, uint, error) {

	//data
	data := Main().Table("vb_blog").Select([]string{
		"vb_blog.id",
		"title",
		"catagory",
		"create_date",
		"GROUP_CONCAT(vb_blog_tag.tag_id) as tags",
	})

	data = data.Joins("LEFT JOIN vb_blog_tag on vb_blog.id=vb_blog_tag.blog_id")
	data = data.Group("vb_blog.id").Having("FIND_IN_SET(?,tags)", info.Id)

	//page
	page := info.Page
	pageSize := info.PageSize

	if page > 0 && pageSize > 0 {
		data = data.Limit(pageSize).Offset((page - 1) * pageSize)
	}

	//result
	var count uint
	result := []VBResultTagBlogList{}
	err := data.Find(&result).Error
	errCount := data.Limit(-1).Offset(-1).Count(&count).Error
	if errCount != nil {
		return []VBResultTagBlogList{}, 0, errCount
	}

	return result, count, err
}

func (m *VBModelTag) New(info *request.VBRTagNew) error {
	//开始事务
	return Main().Transaction(func(tx *gorm.DB) error {

		tag := VBModelTag{
			Name:       info.Name,
			CreateDate: time.Now(),
		}

		err := tx.Create(&tag).Error

		return err
	})
}

func (m *VBModelTag) Delete(info *request.VBRTagDelete) error {

	return Main().Transaction(func(tx *gorm.DB) error {

		//删除标签记录
		if err := tx.Table("vb_tag_record").Where("tag_id = ?", info.Id).Delete(VBModelTagRecord{}).Error; err != nil {
			return err
		}

		//删除标签文章关联
		if err := tx.Table("vb_blog_tag").Where("tag_id = ?", info.Id).Delete(VBModelBlogTag{}).Error; err != nil {
			return err
		}

		//删除标签
		if err := tx.Table("vb_tag").Where("id = ?", info.Id).Delete(VBModelTag{}).Error; err != nil {
			return err
		}

		return nil
	})
}
