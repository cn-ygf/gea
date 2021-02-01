// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"gea/app/dao/internal"
)

// sysRoleDeptDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type sysRoleDeptDao struct {
	*internal.SysRoleDeptDao
}

var (
	// SysRoleDept is globally public accessible object for table sys_role_dept operations.
	SysRoleDept = &sysRoleDeptDao{
		internal.SysRoleDept,
	}
)

// Fill with you ideas below.