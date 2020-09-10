package constant

const (
	//status
	CODE_STATUS_SUCCESS = 20000 //成功状态
	CODE_STATUS_ERROR   = 40000 //失败状态

	//common
	CODE_COMMOM_PARAM_LENTH_ERROR=45000 //参数长度错误

	//admin
	CODE_AMDIN_DONOT_LOGIN      = 40001 //用户未登录
	CODE_ADMIN_TOKEN_INVALID = 40010 //令牌失效
)
