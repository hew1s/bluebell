package models

// 定义请求参数结构体
// 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required"`
}

// Login登录请求参数
type ParamLogin struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
}
