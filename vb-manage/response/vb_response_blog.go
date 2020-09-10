package response

import (
	"time"
)

type VBResultBlogQueryParam struct {
	Id         uint32    `json:"id"`
	Title      string    `json:"title"`
	Status     string    `json:"status"`
	Record     uint      `json:"record"`
	Catagory   string    `json:"catagory"`
	CreateDate time.Time `json:"create_date"`
}

type VBResultBlogInfoMini struct {
	Id         uint32    `json:"id"`
	Title      string    `json:"title"`
	Catagory   string    `json:"catagory"`
	CreateDate time.Time `json:"create_date"`
}

type VBResultBlogInfo struct {
	Id         uint32    `json:"id"`
	Title      string    `json:"title"`
	Status     string    `json:"status"`
	Summary    string    `json:"summary"`
	Catagory   string    `json:"catagory"`
	Directory  string    `json:"directory"`
	CreateDate time.Time `json:"create_date"`
}

type VBResultBlogInfoHTML struct {
	Id   uint32 `json:"id"`
	Html string `json:"html"`
}

type VBResultBlogInfoRecord struct {
	Id     uint32 `json:"id"`
	Record uint   `json:"record"`
}

type VBResultBlogInfoTags struct {
	Id   uint32 `json:"id"`
	Tags string `json:"tags"`
}

//manage 更新文章信息的信息返回
type VBResultBlogInfoManage struct {
	VBResultBlogInfo
	Tags    string `json:"tags"`
	ImageId   uint32 `json:"image_id"`
	Content string `json:"content"`
}

//client
//article
type VBResultBlogInfoClient struct {
	VBResultBlogInfo
	Tags   string `json:"tags"`
	Record uint   `json:"record"`
	Html   string `json:"html"`
}

type VBResultBlogInfoPropose struct {
	VBResultBlogInfoMini
	Record uint `json:"record"`
}

//main
type VBResultBlogMainCount struct {
	Tag     uint   `json:"tag"`
	Record  uint   `json:"record"`
	Article uint   `json:"article"`
	Back    string `json:"back"`
	Avatar  string `json:"avatar"`
}

type VBResultBlogMainRecent struct {
	VBResultBlogInfo
	Tags   string `json:"tags"`
	Image  string `json:"image"`
	Record uint   `json:"record"`
}
