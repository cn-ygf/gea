// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// SysDictType is the golang structure for table sys_dict_type.
type SysDictType struct {
    DictId     int64       `orm:"dict_id,primary"  json:"dict_id"`     // 字典主键             
    DictName   string      `orm:"dict_name"        json:"dict_name"`   // 字典名称             
    DictType   string      `orm:"dict_type,unique" json:"dict_type"`   // 字典类型             
    Status     string      `orm:"status"           json:"status"`      // 状态（0正常 1停用）  
    CreateBy   string      `orm:"create_by"        json:"create_by"`   // 创建者               
    CreateTime *gtime.Time `orm:"create_time"      json:"create_time"` // 创建时间             
    UpdateBy   string      `orm:"update_by"        json:"update_by"`   // 更新者               
    UpdateTime *gtime.Time `orm:"update_time"      json:"update_time"` // 更新时间             
    Remark     string      `orm:"remark"           json:"remark"`      // 备注                 
}