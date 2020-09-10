package request

//blog
//创建新文章
type VBRBlogNew struct {
	Title     string `json:"title" binding:"required"`
	Date      string `json:"date" binding:"required"`
	Catagory  string `json:"catagory" binding:"required"`
	Summary   string `json:"summary" binding:"required"`
	ImageId   uint32 `json:"image_id" binding:"required"`
	Tags      []int  `json:"tags" binding:"required"`
	Directory string `json:"directory"`
	Content   string `json:"content" binding:"required"`
	Html      string `json:"html" binding:"required"`
}

//更新文章
type VBRBlogUpdate struct {
	VBRBlogNew
	Id uint32 `json:"id" binding:"required"`
}

//修改文章状态
type VBRBlogUpdateStatus struct {
	Id     uint32 `json:"id" binding:"required"`
	Status string `json:"status" binding:"required"`
}

//client 获取文章所有信息
type VBRBlogInfo struct {
	Id uint32 `json:"id" binding:"required"`
}

//client 根据文章分类获取文章相关推荐信息 排除当前id的文章
type VBRBlogInfoPropose struct {
	CurrentId uint32 `json:"current_id" binding:"required"`
	Catagory  string `json:"catagory" binding:"required"`
	Order     string `json:"order"`
	Max       uint   `json:"max"`
}

//blog page
//文章条件查询
type VBRBlogQueryParam struct {
	VBRBasePage
	Title    string `json:"title"`
	Date     string `json:"date"`
	Catagory string `json:"catagory"`
}

//client
//文章主页分类查询
type VBRBlogMainRecent struct {
	VBRBasePage
	Catagory string `json:"catagory"`
	Tag      string `json:"tag"`
	Title    string `json:"title"`
}
