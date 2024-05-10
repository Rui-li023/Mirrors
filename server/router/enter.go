package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/attack_awareness"
	"github.com/flipped-aurora/gin-vue-admin/server/router/environment"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System           system.RouterGroup
	Example          example.RouterGroup
	Attack_awareness attack_awareness.RouterGroup
	Environment      environment.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
