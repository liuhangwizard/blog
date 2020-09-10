package model

import (
	"github.com/jinzhu/gorm"
	"time"
	. "vb-server/database/mysql"
	"vb-server/request"
	"vb-server/response"
)

type VBModelBlogRecord struct {
	Id           uint32    `json:"id"`
	BlogId       uint32    `json:"blogid"`
	ReaderIp     string    `json:"readerip"`
	ReaderCity   string    `json:"readercity"`
	ReaderCityId string    `json:"readercityid"`
	ReaderClient string    `json:"readerclient"`
	ReaderDate   time.Time `json:"readerdate"`
}

func NewModelBlogRecord() *VBModelBlogRecord {
	return &VBModelBlogRecord{}
}


func (m *VBModelBlogRecord) DataRecordNew(info *request.VBRBlogRecordNew) error {

	return Main().Transaction(func(tx *gorm.DB) error {

		recordData:=VBModelBlogRecord{
			BlogId:       info.BlogId,
			ReaderIp:     info.Ip,
			ReaderCity:   info.City,
			ReaderCityId: info.CityId,
			ReaderClient: info.Client,
			ReaderDate:   time.Now(),
		}

		err := tx.Table("vb_blog_record").Create(&recordData).Error

		return err
	})

}

func (m *VBModelBlogRecord) DataRecordInfo() (response.VBResultBlogRecordInfo, error) {

	//result
	result := response.VBResultBlogRecordInfo{}

	//info
	now := time.Now()
	todayEnd := now
	todayStart := now.AddDate(0, 0, -1)
	yesterdayStart := now.AddDate(0, 0, -2)

	//point
	var point uint
	var pointToday uint
	if err := Main().Table("vb_blog_record").Count(&point).Error; err != nil {
		return result, err
	}

	//昨天 今天 前天
	if err := Main().Table("vb_blog_record").Where("reader_date BETWEEN ? AND ?", todayStart, todayEnd).Count(&pointToday).Error; err != nil {
		return result, err
	}

	//yesterday
	var pointYesterday uint
	if err := Main().Table("vb_blog_record").Where("reader_date BETWEEN ? AND ?", yesterdayStart, todayStart).Count(&pointYesterday).Error; err != nil {
		return result, err
	}

	toast := "持平"
	yesterdayCount := int(pointToday) - int(pointYesterday)
	if yesterdayCount > 0 {
		toast = "上升"
	}
	if yesterdayCount < 0 {
		toast = "下降"
	}

	result.Point = point
	result.PointToday = pointToday
	result.Yesterday = toast
	result.YesterdayCount = yesterdayCount

	//city
	city := Main().Table("vb_blog_record").Select([]string{
		"reader_city as city",
		"COUNT(reader_city) as city_count",
	})

	if err := city.Group("reader_city").Order("city_count DESC").Limit(1).Find(&result).Error; err != nil {
		return result, err
	}

	//catagory
	catagory := Main().Table("vb_blog").Select([]string{
		"catagory",
		"COUNT( catagory ) AS catagory_count",
	})

	catagory = catagory.Joins("LEFT JOIN vb_blog_record ON vb_blog.id = vb_blog_record.id")
	if err := catagory.Group("catagory").Order("catagory_count DESC").Limit(1).Find(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

func (m *VBModelBlogRecord) DataChartBlogBasic() ([]response.VBResultBlogRecordChartBlogBasic, error) {

	//获取图表的基本信息
	data := Main().Table("vb_blog")

	//select
	data = data.Select([]string{
		"vb_blog.id",
		"title",
		"catagory",
		"LENGTH(content) AS length",
		"COUNT( vb_blog_record.id ) AS record",
	})

	//joins
	data = data.Joins("LEFT JOIN vb_content ON vb_blog.id = vb_content.blog_id")
	data = data.Joins("LEFT JOIN vb_blog_record ON vb_blog.id = vb_blog_record.blog_id")

	//group order
	data = data.Group("vb_blog.id,title,catagory,length").Order("vb_blog.id")

	//result
	result := []response.VBResultBlogRecordChartBlogBasic{}
	err := data.Find(&result).Error

	return result, err
}

func (m *VBModelBlogRecord) DataChartBlogHot() ([]response.VBResultBlogRecordChartBlogHot, error) {

	//获取图表的热度信息
	data := Main().Table("vb_blog")

	//select
	data = data.Select([]string{
		"vb_blog.id",
		"title",
		"COUNT(vb_tag_record.id) as hot",
	})

	//joins
	data = data.Joins("LEFT JOIN vb_blog_tag ON vb_blog.id = vb_blog_tag.blog_id")
	data = data.Joins("LEFT JOIN vb_tag_record ON  vb_blog_tag.tag_id=vb_tag_record.tag_id")

	//group order
	data = data.Group("vb_blog.id,title").Order("vb_blog.id")

	//result
	result := []response.VBResultBlogRecordChartBlogHot{}
	err := data.Find(&result).Error

	return result, err
}

func (m *VBModelBlogRecord) DataQueryParam(param *request.VBRBlogRecordQueryParam) ([]response.VBResultBlogRecordQuery, uint, error) {

	data := Main().Table("vb_blog_record")

	//select
	data = data.Select([]string{
		"reader_date",
		"reader_ip",
		"reader_city_id",
		"reader_city",
		"catagory",
		"title",
	})

	//join
	data = data.Joins("LEFT JOIN vb_blog ON vb_blog.id = vb_blog_record.blog_id")

	//title
	if param.Title != "" {
		data = data.Where("title LIKE ?", "%"+param.Title+"%")
	}

	//catagory
	if param.Catagory != "" {
		data = data.Where("catagory = ?", param.Catagory)
	}

	//city
	if param.City != "" {
		data = data.Where("reader_city LIKE ?", "%"+param.City+"%")
	}

	//date
	if param.StartDate != "" && param.EndDate != "" {
		startDate := param.StartDate + " 00:00:00"
		endDate := param.EndDate + " 23:59:59"
		data = data.Where("reader_date BETWEEN ? AND ?", startDate, endDate)
	}

	//order
	data = data.Order("reader_date DESC")

	//page
	page := param.Page
	pageSize := param.PageSize

	if page > 0 && pageSize > 0 {
		data = data.Limit(pageSize).Offset((page - 1) * pageSize)
	}

	//result
	var count uint
	result := []response.VBResultBlogRecordQuery{}
	err := data.Find(&result).Error
	errCount := data.Limit(-1).Offset(-1).Count(&count).Error
	if errCount != nil {
		return []response.VBResultBlogRecordQuery{}, 0, errCount
	}
	return result, count, err
}

func (m *VBModelBlogRecord) DataBlogRecord(blogId uint32) (response.VBResultBlogInfoRecord, error) {

	//data
	data := Main().Table("vb_blog_record")

	//select
	data = data.Select([]string{
		"blog_id as id",
		"COUNT(blog_id) as record",
	})

	//where
	data = data.Where("vb_blog_record.blog_id = ?", blogId)

	//group
	data = data.Group("blog_id")

	//result
	result := response.VBResultBlogInfoRecord{}
	err := data.Find(&result).Error
	return result, err
}

func (m *VBModelBlogRecord) DataMainCount() response.VBResultBlogMainCount {

	//result
	result := response.VBResultBlogMainCount{}

	//tag
	Main().Table("vb_tag").Count(&result.Tag)
	Main().Table("vb_blog").Count(&result.Article)
	Main().Table("vb_blog_record").Count(&result.Record)

	result.Avatar = "123.png"
	result.Back = "234.png"

	return result
}
