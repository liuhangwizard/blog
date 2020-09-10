package request

type VBRTagNew struct {
	Name string `json:"name" binding:"required"`
}

type VBRTagDelete struct {
	Id	uint32 `json:"id" binding:"required"`
}

type VBRTagBlogList struct {
	VBRBasePage
	Id	uint32 `json:"id" binding:"required"`
}