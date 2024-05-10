package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/attack_awareness"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/environment"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup           system.ApiGroup
	ExampleApiGroup          example.ApiGroup
	Attack_awarenessApiGroup attack_awareness.ApiGroup
	EnvironmentApiGroup      environment.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
