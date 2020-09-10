package response

import "time"

//图片条件查询的响应
type VBResultImageQueryParam struct {
	Id         uint32    `json:"id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Src        string    `json:"src"`
	CreateDate time.Time `json:"create_date"`
}
