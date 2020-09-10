package model

import (
	. "vb-server/database/mysql"
	"vb-server/response"
)

type VBModelContent struct {
	Id      uint32 `json:"id"`
	BlogId  uint32 `json:"blogid"`
	Content string `json:"content"`
	Html    string `json:"html"`
}

func NewModelContent() *VBModelContent {
	return &VBModelContent{}
}

func (m *VBModelContent) TableName() string {
	return "vb_content"
}


func (m *VBModelContent) DataBlogHTML(blogId uint32) (response.VBResultBlogInfoHTML,error){

	//data
	data := Main().Table("vb_content")

	//select
	data=data.Select([]string{
		"blog_id as id",
		"html",
	})

	//where
	data=data.Where("vb_content.blog_id = ?",blogId)

	//result
	result:=response.VBResultBlogInfoHTML{}
	err:=data.Find(&result).Error
	return result,err
	
}
