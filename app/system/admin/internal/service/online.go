package service

import (
	"gea/app/dao"
	"gea/app/system/admin/internal/define"
	"gea/app/utils/page"
	"gea/app/utils/token"
)

var Online = &onlineService{}

type onlineService struct{}

//根据条件分页查询数据
func (s *onlineService) GetList(param *define.OnlineApiSelectPageReq) *define.OnlineServiceList {
	m := dao.SysUserOnline.As("t")
	if param != nil {

		if param.Token != "" {
			m = m.Where("t.token = ?", param.Token)
		}

		if param.LoginName != "" {
			m = m.Where("t.login_name like ?", "%"+param.LoginName+"%")
		}

		if param.DeptName != "" {
			m = m.Where("t.dept_name like ?", "%"+param.DeptName+"%")
		}

		if param.Ipaddr != "" {
			m = m.Where("t.ipaddr = ?", param.Ipaddr)
		}

		if param.LoginLocation != "" {
			m = m.Where("t.login_location = ?", param.LoginLocation)
		}

		if param.Browser != "" {
			m = m.Where("t.browser = ?", param.Browser)
		}

		if param.Os != "" {
			m = m.Where("t.os = ?", param.Os)
		}

		if param.Status != "" {
			m = m.Where("t.status = ?", param.Status)
		}

		if param.BeginTime != "" {
			m = m.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			m = m.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	total, err := m.Count()
	if err != nil {
		return nil
	}
	page := page.CreatePaging(param.PageNum, param.PageSize, total)
	m = m.Limit(page.StartNum, page.Pagesize)
	if param.OrderByColumn != "" {
		m = m.Order(param.OrderByColumn + " " + param.IsAsc)
	}
	result := &define.OnlineServiceList{
		Page:  page.PageNum,
		Size:  page.Pagesize,
		Total: page.Total,
	}
	if err = m.Structs(&result.List); err != nil {
		return nil
	}
	return result
}

// 强退
func (s *onlineService)ForceLogout(tokenStr string) {
	token.RemoveCache(tokenStr)
	// 修改状态
	dao.SysUserOnline.Delete(dao.SysUserOnline.Columns.Token,tokenStr)
}
