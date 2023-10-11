package logic

import (
	"bulebell/dao/mysql"
	"bulebell/dao/redis"
	"bulebell/models"
	"bulebell/pkg/jwt"
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
	userID := snowflake.GetID()
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
func Login(p *models.ParamLogin)(token string,err error){
	// 
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递出来的是一个user指针，包含user信息
	if err :=  mysql.Login(user);err != nil{
		return "",err
	}
	
	// 1.将用户的userID与token进行对应存入redis
	token ,_= jwt.GetToken(user.UserID,user.Username)
	redis.Login(user.UserID,token)
	// 2.当重新登陆时就删除旧token插入新token
	// 3.在token检验中对相应userID的token进行检测
	// user.UserID 生成JWT
	return jwt.GetToken(user.UserID,user.Username)
}