package request

//文章主页图片查询 根据文章的id来查询 是一个list
type VBRBlogMainRecentImageList struct {
	BlogIdList []uint32 `json:"blog_id_list"`
}

//单个文章的图片信息
type VBRBlogInfoImage struct {
	BlogId uint32 `json:"blog_id"`
}
