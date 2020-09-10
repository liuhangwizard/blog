package request

type VBRCatagoryNew struct {
	Content string `json:"content" binding:"required"`
}

type VBRCatagoryDelete struct {
	Id	uint32 `json:"id" binding:"required"`
}