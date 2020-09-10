package request

//admin
type VBRAdminLogin struct {
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}

