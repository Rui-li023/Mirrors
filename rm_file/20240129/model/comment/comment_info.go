// 自动生成模板CommentInfo
package comment

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 评论信息 结构体  CommentInfo
type CommentInfo struct {
      global.GVA_MODEL
      CommentUserId  *int `json:"commentUserId" form:"commentUserId" gorm:"column:comment_user_id;comment:评论人id;"`  //评论人id 
      CommentText  string `json:"commentText" form:"commentText" gorm:"column:comment_text;comment:评论信息;size:300;"`  //评论信息 
      CommentCommentId  *int `json:"commentCommentId" form:"commentCommentId" gorm:"column:comment_comment_id;comment:评论的评论id;"`  //评论的评论id 
      CommentLikeNumber  *int `json:"commentLikeNumber" form:"commentLikeNumber" gorm:"column:comment_like_number;comment:评论的评论id;"`  //评论点赞 
}


// TableName 评论信息 CommentInfo自定义表名 commentInfo
func (CommentInfo) TableName() string {
  return "commentInfo"
}

