package mysql

import (
	"bulebell/models"
)

func CreatePost(post *models.Post)(err error) {
	sqlStr := `insert into post
	(post_id,author_id,community_id,title,content)
	value(?,?,?,?,?)`
	_,err = db.Exec(sqlStr,post.ID,post.AuthorID,post.CommunityID,post.Title,post.Content)
	return
}

func GetPostByID(pid int64) (data *models.Post,err error) {
	data = new(models.Post)
	sqlStr := `select post_id,title,content,author_id,community_id,create_time from post where post_id = ?`
	err = db.Get(data,sqlStr,pid)
	return
}

func GetPostList(pageNum,pageSize int64)(post []*models.Post,err error){
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
	from post
	ORDER BY create_time
	DESC
	limit ?,?`
	post= make([]*models.Post, 0,2)
	err = db.Select(&post,sqlStr,(pageNum-1)*pageSize,pageSize)
	return
}