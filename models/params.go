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
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 投票数据
type ParamVoteData struct {
	PostID    string `json:"post_id" binding:"required"`
	Direction int8  `json:"direction,string" binding:"required,oneof=1 0 -1"` //赞成（1）或者反对（-1）
}
