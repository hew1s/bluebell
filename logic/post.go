package logic

import (
	"bulebell/dao/mysql"
	"bulebell/models"
	"bulebell/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error){
	// 1.生成post_id
	p.ID = int64(snowflake.GetID())
	// 2.保存数据库
	return mysql.CreatePost(p)
}