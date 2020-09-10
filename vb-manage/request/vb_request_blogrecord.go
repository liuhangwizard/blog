package request

type VBRBlogRecordNew struct {
	BlogId uint32 `json:"blog_id" binding:"required"`
	Ip     string `json:"ip"`
	City   string `json:"city"`
	CityId string `json:"city_id"`
	Client string `json:"client" binding:"required"`
}

type VBRBlogRecordQueryParam struct {
	VBRBasePage
	Title     string `json:"title"`
	City      string `json:"city"`
	Catagory  string `json:"catagory"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
