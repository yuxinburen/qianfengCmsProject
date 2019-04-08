package service

import "github.com/go-xorm/xorm"

/**
 * 商店Shop的服务
 */
type ShopService interface {
	//查询商店总数，并返回
	GetShopCount() int64
}

type shopService struct {
	Engine *xorm.Engine
}

/**
 * 新实例化一个商店模块服务对象结构体
 */
func NewShopService(engine *xorm.Engine) ShopService {
	return &shopService{Engine: engine}
}

/**
 * 查询商店的总数然后返回
 */
func (ss *shopService) GetShopCount() int64 {
	return 0
}
