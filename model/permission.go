package model

/**
 * 管理员权限表格（基础表格）
 */
type Permission struct {
	PermissionId int64  `xorm:"pk autoincr"json:"id"`     //主键id
	Level        string `xorm:"varchar(32)" json:"level"` //权限级别
	Name         string `xorm:"varchar(32)" json:"name"`  //权限名称
}
