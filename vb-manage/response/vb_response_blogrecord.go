package response

import (
	"time"
)

//总面板的响应
type VBResultBlogRecordInfo struct {
	Point          uint   `json:"point"`           //访问总数
	PointToday     uint   `json:"point_today"`     //今日访问数
	Yesterday      string `json:"yesterday"`       //访问趋势 add sub equal
	YesterdayCount int    `json:"yesterday_count"` //趋势数量 正负数
	City           string `json:"city"`            //最多点击城市区域
	CityCount      uint   `json:"city_count"`      //城市区域计数
	Catagory       string `json:"catagory"`        //最多点击分类
	CatagoryCount  uint   `json:"catagory_count"`  //点击分类计数
}

//访问记录的统计
type VBResultBlogRecordQuery struct {
	ReaderDate   time.Time `json:"date"`
	ReaderIp     string    `json:"ip"`
	ReaderCityId string    `json:"code"`
	ReaderCity   string    `json:"city"`
	Catagory     string    `json:"catagory"`
	Title        string    `json:"title"`
}

//图表
//图表 blog
//blog 基本信息 包括字数统计 访问量
type VBResultBlogRecordChartBlogBasic struct {
	Id       uint32 `json:"id"`
	Title    string `json:"title"`
	Catagory string `json:"catagory"`
	Length   uint   `json:"length"`
	Record   uint   `json:"record"`
}

//blog 热度信息 包括热度
type VBResultBlogRecordChartBlogHot struct {
	Id    uint32 `json:"id"`
	Title string `json:"title"`
	Hot   uint   `json:"hot"`
}

//blog 基本信息 热度信心的整合
type VBResultBlogRecordChartBlog struct {
	Id       uint32 `json:"id"`
	Title    string `json:"title"`
	Catagory string `json:"catagory"`
	Length   uint   `json:"length"`
	Record   uint   `json:"record"`
	Hot      uint   `json:"hot"`
}
