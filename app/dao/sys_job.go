// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"gea/app/dao/internal"
)

// sysJobDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type sysJobDao struct {
	*internal.SysJobDao
}

var (
	// SysJob is globally public accessible object for table sys_job operations.
	SysJob = &sysJobDao{
		internal.SysJob,
	}
)

// Fill with you ideas below.