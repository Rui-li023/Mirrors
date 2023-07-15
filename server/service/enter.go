package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/NestAirlinePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/NestExecRecordPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/NestInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Nestrolepkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/UserTeemlinkPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup            system.ServiceGroup
	ExampleServiceGroup           example.ServiceGroup
	NestInfoServiceGroup          NestInfo.ServiceGroup
	NestrolepkgServiceGroup       Nestrolepkg.ServiceGroup
	NestAirlinePkgServiceGroup    NestAirlinePkg.ServiceGroup
	NestExecRecordPkgServiceGroup NestExecRecordPkg.ServiceGroup
	UserTeemlinkPkgServiceGroup   UserTeemlinkPkg.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
