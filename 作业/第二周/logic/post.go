package logic

import "github.com/pkg/errors"
import mysql

func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, errors.Wrap(err, "mysql getPostList failed")
	}

	//其他处理
	return
}
