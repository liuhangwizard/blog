package model

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
	. "vb-server/database/mysql"
	"vb-server/request"
	"vb-server/response"
)

type VBModelBlog struct {
	Id         uint32          `json:"id"`
	UserId     uint32          `json:"userid"`
	Title      string          `json:"title"`
	Summary    string          `json:"summary"`
	Catagory   string          `json:"catagory"`
	Directory  string          `json:"directory"`
	Config     json.RawMessage `json:"config"`
	Status     string          `json:"status"`
	CreateDate time.Time       `json:"createdate"`
	UpdateDate time.Time       `json:"updatedate"`
}

func NewModelBlog() *VBModelBlog {
	return &VBModelBlog{}
}

//server
func (m *VBModelBlog) DataNewBlog(article *request.VBRBlogNew, userId uint32) error {

	//开始事务
	return Main().Transaction(func(tx *gorm.DB) error {
		//创建文章记录
		date, err := time.ParseInLocation("2006-01-02 15:04:05", article.Date+" 23:59:59", time.Local)
		if err != nil {
			return err
		}

		blogData := VBModelBlog{
			UserId:     userId,
			Title:      article.Title,
			Catagory:   article.Catagory,
			Summary:    article.Summary,
			Directory:  article.Directory,
			Config:     nil,
			Status:     "1",
			CreateDate: date,
			UpdateDate: time.Now(),
		}

		if err := tx.Table("vb_blog").Create(&blogData).Error; err != nil {
			return err
		}

		//new id
		newBlogId := blogData.Id

		//关联标签
		tagIdList := article.Tags
		for _, id := range tagIdList {

			blogTag := VBModelBlogTag{
				BlogId: newBlogId,
				TagId:  uint32(id),
			}
			if count := tx.Table("vb_blog_tag").Create(&blogTag).RowsAffected; count == 0 {
				return errors.New("connect tag blog fail")
			}
		}

		//关联主图
		blogImageData := VBModelBlogImage{
			BlogId:  newBlogId,
			ImageId: article.ImageId,
		}

		if err := tx.Table("vb_blog_image").Create(&blogImageData).Error; err != nil {
			return err
		}

		//关联内容
		contentData := VBModelContent{
			BlogId:  newBlogId,
			Content: article.Content,
			Html:    article.Html,
		}

		if err := tx.Table("vb_content").Create(&contentData).Error; err != nil {
			return err
		}

		//commit
		return nil
	})

}

func (m *VBModelBlog) DataUpdateStatus(article *request.VBRBlogUpdateStatus) error {

	//开始事务
	return Main().Transaction(func(tx *gorm.DB) error {

		if count := tx.Table("vb_blog").Select("status").Where("id = ?", article.Id).Update(map[string]interface{}{"status": article.Status}).RowsAffected; count == 0 {
			return errors.New("update status fail")
		}

		return nil
	})
}

func (m *VBModelBlog) DataUpdateBlog(article *request.VBRBlogUpdate) error {

	//开始事务
	return Main().Transaction(func(tx *gorm.DB) error {

		//id
		blogId := article.Id

		//删除旧的tag关联
		if err := tx.Table("vb_blog_tag").Where("blog_id = ?", blogId).Delete(VBModelBlogTag{}).Error; err != nil {
			return err
		}

		//关联新tag
		tagIdList := article.Tags
		for _, id := range tagIdList {

			blogTag := VBModelBlogTag{
				BlogId: blogId,
				TagId:  uint32(id),
			}
			if err := tx.Table("vb_blog_tag").Create(&blogTag).Error; err !=nil {
				return err
			}
		}


		//删除旧的image关联
		if err:=tx.Table("vb_blog_image").Where("blog_id = ?",blogId).Delete(VBModelBlogImage{}).Error;err!=nil{
			return err
		}

		//关联新的image
		blogImage:=VBModelBlogImage{
			BlogId:  blogId,
			ImageId: article.ImageId,
		}
		if err := tx.Table("vb_blog_image").Create(&blogImage).Error; err != nil {
			return err
		}


		//更新content
		if err := tx.Table("vb_content").Where("blog_id = ?", blogId).Update(map[string]interface{}{
			"content": article.Content,
			"html":    article.Html,
		}).Error; err != nil {
			return err
		}

		//更新basic date
		cDate, err := time.ParseInLocation("2006-01-02 15:04:05", article.Date+" 23:59:59", time.Local)
		if err != nil {
			return err
		}

		//更新basic
		newBlog := VBModelBlog{
			Title:      article.Title,
			Summary:    article.Summary,
			Catagory:   article.Catagory,
			Directory:  article.Directory,
			CreateDate: cDate,
			UpdateDate: time.Now(),
		}
		if err := tx.Table("vb_blog").Where("id =?", blogId).Updates(newBlog).Error; err != nil {
			return err
		}

		return nil
	})
}

func (m *VBModelBlog) DataQueryParam(param *request.VBRBlogQueryParam, userId uint32) ([]response.VBResultBlogQueryParam, uint, error) {

	data := Main().Table("vb_blog")

	//select
	//data = data.Select([]string{"id", "title", "status", "create_date", "catagory"})
	data = data.Select([]string{"vb_blog.id", "title", "catagory", "status", "create_date", "COUNT(vb_blog_record.blog_id) AS record"})

	//join
	data = data.Joins("LEFT JOIN vb_blog_record ON vb_blog.id=vb_blog_record.blog_id")

	//title
	if param.Title != "" {
		data = data.Where("title LIKE ?", param.Title+"%")
	}

	//catagory
	if param.Catagory != "" {
		data = data.Where("catagory = ?", param.Catagory)
	}

	//end date
	if param.Date != "" {
		finalDate := param.Date + " 23:59:59"
		data = data.Where("create_date <= ?", finalDate)
	}

	//group
	data = data.Group("vb_blog.id")

	//order
	data = data.Order("vb_blog.create_date DESC")

	//page
	page := param.Page
	pageSize := param.PageSize

	if page > 0 && pageSize > 0 {
		data = data.Limit(pageSize).Offset((page - 1) * pageSize)
	}

	//result
	var count uint
	result := []response.VBResultBlogQueryParam{}
	err := data.Find(&result).Error
	errCount := data.Limit(-1).Offset(-1).Count(&count).Error
	if errCount != nil {
		return []response.VBResultBlogQueryParam{}, 0, errCount
	}
	return result, count, err
}

func (m *VBModelBlog) DataManageInfo(article *request.VBRBlogInfo) (response.VBResultBlogInfoManage, error) {

	//"vb_blog,vb_blog_tag,vb_content"
	data := Main().Table("vb_blog")

	//select
	data = data.Select([]string{"vb_blog.id", "title", "status", "summary", "catagory", "directory", "create_date", "content", "GROUP_CONCAT(vb_blog_tag.tag_id) as tags"})

	//join
	data = data.Joins("LEFT JOIN vb_blog_tag ON vb_blog.id = vb_blog_tag.blog_id")
	data = data.Joins("LEFT JOIN vb_content ON vb_blog.id=vb_content.blog_id")

	//where
	data = data.Where("vb_blog.id = ?", article.Id)

	//group
	data = data.Group("vb_blog.id,vb_content.content")

	//result
	result := response.VBResultBlogInfoManage{}
	err := data.Find(&result).Error
	return result, err

}

//client article
func (m *VBModelBlog) DataClientInfo(article *request.VBRBlogInfo) (response.VBResultBlogInfoClient, error) {

	//这里的sql还可以再优化

	//result
	result := response.VBResultBlogInfoClient{}

	//data
	dataTags := Main().Table("vb_blog")

	//select
	dataTags = dataTags.Select([]string{
		"vb_blog.id",
		"title",
		"status",
		"summary",
		"catagory",
		"directory",
		"create_date",
		"GROUP_CONCAT(vb_blog_tag.tag_id) as tags",
	})

	dataTags = dataTags.Joins("LEFT JOIN vb_blog_tag ON vb_blog.id = vb_blog_tag.blog_id")

	//where
	dataTags = dataTags.Where("vb_blog.id = ?", article.Id).Group("vb_blog.id")

	if err := dataTags.Find(&result).Error; err != nil {
		return result, err
	}

	//先联合查询tags 单独查询 record html
	modelContent := NewModelContent()
	modelBlogRecord := NewModelBlogRecord()

	dataHTML, err := modelContent.DataBlogHTML(article.Id)
	if err != nil {
		return result, err
	}

	dataRecord, _ := modelBlogRecord.DataBlogRecord(article.Id)
	result.Html = dataHTML.Html
	result.Record = dataRecord.Record

	return result, nil
}

func (m *VBModelBlog) DataClientInfoPropose(article *request.VBRBlogInfoPropose) ([]response.VBResultBlogInfoPropose, error) {

	//data
	data := Main().Table("vb_blog")

	//select
	data = data.Select([]string{
		"vb_blog.id",
		"title",
		"catagory",
		"create_date",
		"COUNT(vb_blog_record.id) AS record",
	})

	//join
	data = data.Joins("LEFT JOIN vb_blog_record ON vb_blog.id=vb_blog_record.blog_id")

	//where
	if article.Catagory != "" && article.CurrentId != 0 {
		data = data.Where("catagory = ? AND vb_blog.id <> ?", article.Catagory, article.CurrentId)
	}

	//group
	data = data.Group("vb_blog.id")

	//oreder
	if article.Order != "" {
		data = data.Order(article.Order)
	}

	//limit
	if article.Max > 0 {
		data = data.Limit(article.Max)
	}

	//result
	result := []response.VBResultBlogInfoPropose{}
	err := data.Find(&result).Error

	return result, err
}

//client main
func (m *VBModelBlog) DataClientMainArchive() ([]response.VBResultBlogInfoPropose, error) {

	//data
	data := Main().Table("vb_blog")

	//select
	data = data.Select([]string{
		"vb_blog.id",
		"title",
		"catagory",
		"create_date",
		"COUNT(vb_blog_record.id) AS record",
	})

	//join
	data = data.Joins("LEFT JOIN vb_blog_record ON vb_blog.id=vb_blog_record.blog_id")

	//group
	data = data.Group("vb_blog.id").Order("create_date DESC")

	//result
	result := []response.VBResultBlogInfoPropose{}
	err := data.Find(&result).Error

	return result, err
}

func (m *VBModelBlog) DataClientMainRecent(article *request.VBRBlogMainRecent) ([]response.VBResultBlogMainRecent, uint, error) {

	//data
	data := Main().Table("vb_blog")

	//select
	//联合 标签 统计进行查询 去除重复
	data = data.Select([]string{
		"vb_blog.id",
		"title",
		"catagory",
		"summary",
		"vb_blog.create_date",
		"GROUP_CONCAT( DISTINCT vb_tag.name ) AS tags",
		"COUNT(DISTINCT vb_blog_record.id) as record",
	})

	//join
	data = data.Joins("LEFT JOIN vb_blog_record ON vb_blog.id=vb_blog_record.blog_id")
	data = data.Joins("LEFT JOIN vb_blog_tag ON vb_blog.id = vb_blog_tag.blog_id")
	data = data.Joins("LEFT JOIN vb_tag ON vb_blog_tag.tag_id=vb_tag.id")

	//where
	if article.Catagory != "" {
		data = data.Where("catagory = ?", article.Catagory)
	}

	if article.Title != "" {
		data = data.Where("title like ?", "%"+article.Title+"%")
	}

	//having
	if article.Tag != "" {
		data = data.Having("FIND_IN_SET(?,tags)", article.Tag)
	}

	//group
	data = data.Group("vb_blog.id").Order("create_date DESC")

	page := article.Page
	pageSize := article.PageSize

	if page > 0 && pageSize > 0 {
		data = data.Limit(pageSize).Offset((page - 1) * pageSize)
	}

	//result
	result := []response.VBResultBlogMainRecent{}
	var countBlog uint
	errBlog := data.Find(&result).Error
	errBlogCount := data.Limit(-1).Offset(-1).Count(&countBlog).Error
	if errBlogCount != nil {
		return result, countBlog, errBlogCount
	}

	return result, countBlog, errBlog
}
