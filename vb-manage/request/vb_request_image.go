package request

//获取上传签名
type VBRImageNewSign struct {
	FileName string `json:"filename" binding:"required"`
}

//上传回调
type VBRImageNew struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
	Src  string `json:"src" binding:"required"`
}

//图片条件查询
type VBRImageQueryParam struct {
	VBRBasePage
	Name string `json:"name"`
	Date string `json:"date"`
	Type string `json:"type"`
}

//修改图片分类
type VBRImageUpdateType struct {
	Id   uint32 `json:"id" binding:"required"`
	Type string `json:"type" binding:"required"`
}
