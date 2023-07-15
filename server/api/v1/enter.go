package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/NestAirlinePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/NestExecRecordPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/NestInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Nestrolepkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/UserTeemlinkPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup            system.ApiGroup
	ExampleApiGroup           example.ApiGroup
	NestInfoApiGroup          NestInfo.ApiGroup
	NestrolepkgApiGroup       Nestrolepkg.ApiGroup
	NestAirlinePkgApiGroup    NestAirlinePkg.ApiGroup
	NestExecRecordPkgApiGroup NestExecRecordPkg.ApiGroup
	UserTeemlinkPkgApiGroup   UserTeemlinkPkg.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
