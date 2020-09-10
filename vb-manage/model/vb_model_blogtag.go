package model

type VBModelBlogTag struct {
	Id     uint32 `json:"id"`
	BlogId uint32 `json:"blogid"`
	TagId  uint32 `json:"tagid"`
}

func NewModelBlogTag() *VBModelBlogTag {
	return &VBModelBlogTag{}
}
