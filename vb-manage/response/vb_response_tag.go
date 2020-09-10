package response

import "time"

type VBResultTagBlogList struct {
	Id         uint32    `json:"id"`
	Title      string    `json:"title"`
	Catagory   string    `json:"catagory"`
	Tags       string    `json:"tags"`
	CreateDate time.Time `json:"create_date"`
}
