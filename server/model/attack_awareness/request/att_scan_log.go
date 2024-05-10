package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type ScanlogSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      Source_ip  string `json:"source_ip" form:"source_ip" `
                      Source_port  *int `json:"source_port" form:"source_port" `
                      Protocol  string `json:"protocol" form:"protocol" `
                      Dest_ip  string `json:"dest_ip" form:"dest_ip" `
                      Dest_port  *int `json:"dest_port" form:"dest_port" `
    request.PageInfo
}
