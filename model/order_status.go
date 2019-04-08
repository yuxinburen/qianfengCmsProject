package model

/**
 * 订单状态结构体定义
 */
type OrderStatus struct {
	Id         int    `xorm:"pk antoincr" json:"id"` //主键
	StatusId   int                                   //订单状态编号
	StatusDesc string `xorm:"varchar(255)"`          // 订单状态描述

}
