package model

/**
 * 管理员的权限明细表（业务结构表）
 */
type AdminPermission struct {
	Admin      *Admin      `xorm:"extends"` //不需要映射admin结构体
	Permission *Permission `xorm:"extends"` //不需要映射permission结构体
}
