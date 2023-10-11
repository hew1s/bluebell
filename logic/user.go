package logic

import (
	"bulebell/dao/mysql"
	"bulebell/models"
	"bulebell/pkg/snowflake"
)

// 存放业务逻辑的代码

// 注册业务
func SignUp(p *models.ParamSignUp)(err error) {
	// 判断用户是否存在
	if err := mysql.CheckUserExist(p.Username);err!=nil{
		return err
	}
	// 2.生成uid
	userID := snowflake.GenID()
	// 构造User实例
	u := &models.User{
		UserID: userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 3.密码加密
	// 4.保存进入数据库
	return mysql.InsertUser(u)
}

// 登录业务
func Login(p *models.ParamLogin)error{
	// 
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)
}