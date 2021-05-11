package router

import (
	"github.com/eyotang/gin-vue-admin/server/api/v1"
	"github.com/eyotang/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

func InitEmailRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("email").Use(middleware.OperationRecord())
	{
		UserRouter.POST("emailTest", v1.EmailTest) // 发送测试邮件
	}
}
