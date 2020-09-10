package response

type VBResultBlogImage struct {
	Id    uint32 `json:"id"` //这个是文章的id
	Image string `json:"image"`
}

type VBResultBlogImageInfo struct {
	ImageId uint32 `json:"id"` //这个是image的id
	Image   string `json:"image"`
}
