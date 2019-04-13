package service

import (
	"irisDemo/CmsProject/model"
	"github.com/go-xorm/xorm"
)

/**
 * 管理员服务
 * 标准的开发模式将每个实体的提供的功能以接口标准的形式定义,供控制层进行调用。
 *
 */
type AdminService interface {
	//通过管理员用户名+密码 获取管理员实体 如果查询到，返回管理员实体，并返回true
	//否则 返回 nil ，false
	GetByAdminNameAndPassword(username, password string) (model.Admin, bool)

	//获取管理员总数
	GetAdminCount() (int64, error)
}

func NewAdminService(db *xorm.Engine) AdminService {
	return &adminSevice{
		engine: db,
	}
}

/**
 * 管理员的服务实现结构体
 */
type adminSevice struct {
	engine *xorm.Engine
}

/**
 * 查询管理员总数
 */
func (ac *adminSevice) GetAdminCount() (int64, error) {
	count, err := ac.engine.Count(new(model.Admin))

	if err != nil {
		panic(err.Error())
		return 0, err
	}
	return count, nil
}

/**
 * 通过用户名和密码查询管理员
 */
func (ac *adminSevice) GetByAdminNameAndPassword(username, password string) (model.Admin, bool) {
	var admin model.Admin

	ac.engine.Where(" admin_name = ? and pwd = ? ", username, password).Get(&admin)

	return admin, admin.AdminId != 0
}
