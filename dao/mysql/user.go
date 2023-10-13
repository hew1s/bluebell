package mysql

import (
	"bulebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

// 将数据库每一步操作封装成函数
// logic层需求调用

const secret = "liwenzhou.com"

// 检查用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return ErrorUserExist
	}
	if count > 0 {
		return err
	}
	return
}

// InsertUser插入用户
func InsertUser(u *models.User) (err error) {
	// 加密密码
	password := encryptPassword(u.Password)
	// 执行sql语句入库
	sqlStr := "insert into user(user_id,username,password) values(?,?,?)"
	_, err = db.Exec(sqlStr, u.UserID, u.Username, password)
	return
}

// md5加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := "select user_id,username,password from user where username =?"
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}


func GetUserByID(uid int64) (user *models.User,err error) {
	user = new(models.User)
	sqlStr := "select username from user where user_id = ?"
	err = db.Get(user,sqlStr,uid)
	return user,err
}