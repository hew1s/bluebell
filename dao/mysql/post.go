package mysql

import "bulebell/models"

func CreatePost(post *models.Post)(err error) {
	sqlStr := `insert into post
	(post_id,author_id,community_id,title,content)
	value(?,?,?,?,?)`
	_,err = db.Exec(sqlStr,post.ID,post.AuthorID,post.CommunityID,post.Title,post.Content)
	return
}