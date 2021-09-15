package controller

import (
	"fmt"
	"logic"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const CodeServerBusy = "服务器繁忙"

// GetPostListHandler 获取帖子列表的处理函数
func GetPostListHandler(c *gin.Context) {
	// 获取分页参数
	page, size := getPageInfo(c)
	// 获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		fmt.Printf("original err:%T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n %+v\n", err)
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
	// 返回响应
}
