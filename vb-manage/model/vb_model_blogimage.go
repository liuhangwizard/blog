package model

import (
	"strconv"
	"strings"
	. "vb-server/database/mysql"
	"vb-server/request"
	"vb-server/response"
)

type VBModelBlogImage struct {
	Id      uint32 `json:"id"`
	BlogId  uint32 `json:"blog_id"`
	ImageId uint32 `json:"image_id"`
}

func NewModelBlogImage() *VBModelBlogImage {
	return &VBModelBlogImage{}
}

func (m *VBModelBlogImage) DataBlogRecentImageList(query *request.VBRBlogMainRecentImageList) ([]response.VBResultBlogImage, error) {

	idList := make([]string, len(query.BlogIdList))
	for i, v := range query.BlogIdList {
		idList[i] = strconv.Itoa(int(v))
	}

	//id组成的字符串列表
	idListContent := strings.Join(idList, ",")

	//data
	data := Main().Table("vb_blog")

	//select
	//联合 标签 统计进行查询 去除重复
	data = data.Select([]string{
		"vb_blog.id",
		"vb_image.src as image",
	})

	//join
	data = data.Joins("LEFT JOIN vb_blog_image ON vb_blog_image.blog_id=vb_blog.id")
	data = data.Joins("LEFT JOIN vb_image ON vb_blog_image.image_id=vb_image.id")

	//where
	data = data.Where("FIND_IN_SET(vb_blog.id,?)", idListContent)

	//order
	data = data.Order("FIND_IN_SET(vb_blog.id,'" + idListContent + "')")

	//result
	result := []response.VBResultBlogImage{}
	err := data.Find(&result).Error

	return result, err
}

func (m *VBModelBlogImage) DataBlogImageInfo(query *request.VBRBlogInfoImage) (response.VBResultBlogImageInfo, error) {

	//data
	data := Main().Table("vb_blog_image")

	//select
	data = data.Select([]string{
		"vb_blog_image.image_id",
		"src AS image",
	})

	//join
	data = data.Joins("LEFT JOIN vb_image ON vb_blog_image.image_id = vb_image.id ")

	//where
	data = data.Where("vb_blog_image.blog_id = ?", query.BlogId)

	//result
	result := response.VBResultBlogImageInfo{}
	err := data.Find(&result).Error

	return result, err

}
