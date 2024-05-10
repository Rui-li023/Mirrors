// 自动生成模板Scanlog
package attack_awareness

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 扫描感知 结构体  Scanlog
type Scanlog struct {
 global.GVA_MODEL 
      Source_ip  string `json:"source_ip" form:"source_ip" gorm:"column:source_ip;comment:;size:20;"`  //源IP 
      Source_port  *int `json:"source_port" form:"source_port" gorm:"column:source_port;comment:;size:18;"`  //源端口 
      Protocol  string `json:"protocol" form:"protocol" gorm:"column:protocol;comment:;size:10;"`  //协议 
      Dest_ip  string `json:"dest_ip" form:"dest_ip" gorm:"column:dest_ip;comment:;size:20;"`  //目的IP 
      Dest_port  *int `json:"dest_port" form:"dest_port" gorm:"column:dest_port;comment:;size:18;"`  //目的端口 
      Detail  string `json:"detail" form:"detail" gorm:"column:detail;comment:;type:text;"`  //详细内容 
}


// TableName 扫描感知 Scanlog自定义表名 scanlog
func (Scanlog) TableName() string {
  return "scanlog"
}

