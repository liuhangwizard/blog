package constant

const (

	//common
	COMMON_SUCCESS           = "success"
	COMMON_ERROR             = "error"
	COMMOM_ERROR_PARAM_LENTH = "参数长度错误"
	COMMOM_ERROR_PARAM_BIND  = "参数绑定失败"

	//login
	//login success
	ADMIN_SUCCESS_LOGIN = "登录成功"

	//login error
	ADMIN_ERROR_LOGIN_PARAM   = "错误的登录请求参数"
	ADMIN_ERROR_LOGIN_INVALID = "用户名或密码错误"

	//login content
	ADMIN_ERROR_DONOT_LOGIN  = "用户未登录"
	ADMIN_ERROR_SESSION_NONE = "用户会话已经过期"

	//permission

	//blog
	BLOG_SUCCESS_NEW         = "创建文章成功"
	BLOG_ERROR_NEW           = "创建文章失败"
	BLOG_SUCCESS_PARAM_QUERY = "文章参数查询成功"
	BLOG_ERROR_PARAM_QUERY   = "文章参数查询失败"
	//database
	DATABASE_ERROR_QUERY_FAIL = "查询数据库失败"

	//image
	IMAGE_SUCCESS_NEW="新增图片成功"
	IMAGE_ERROR_FILENAME  = "错误的文件名"
	IMAGE_ERROR_FILEBACK  ="错误的文件返回参数"
	IMAGE_ERROR_SIGN_FAIL = "签名失败"
	IMAGE_ERROR_TYPE_FAIL ="修改图片类型失败"
	IMAGE_ERROR_TYPE_PARAM="错误的图片类型修改参数"
)
