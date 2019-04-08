package model

import "time"

/**
 * 用户订单结构实体定义
 */
type UserOrder struct {
	Id            int64        `xorm:"pk autoincr" json:"id"` //主键
	SumMoney      int64        `xorm:"default 0" json:"sum_money"`
	Time          time.Time    `xorm:"DateTime" json:"time"`         //时间
	OrderTime     uint64       `json:"order_time"`                   //订单创建时间
	OrderStatusId int64        `xorm:"index" json:"order_status_id"` //订台状态id
	OrderStatus   *OrderStatus `xorm:"-"`                            //订单对象
	UserId        int64        `xorm:"index" json:"user_id"`         //用户编号Id
	User          *User        `xorm:"-"`                            //订单对应的账户，并不进行结构体字段映射
	ShopId        int64        `xorm:"index" json:"shop_id"`         //用户购买的商品编号
	Shop          *Shop        `xorm:"-"`                            //商品结构体，不进行映射
	AddressId     int64        `xorm:"index" json:"address_id"`      //地址结构体的Id
	Address       *Address     `xorm:"-"`                            //地址结构体，不进行映射
	DelFlag       int64        `xorm:"default 0" json:"del_flag"`    //删除标志 0为正常 1为已删除
}
