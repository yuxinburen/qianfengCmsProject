package model

import (
	"time"
	"irisDemo/CmsProject/utils"
)

/**
 * 用户信息结构体,用于生成用户信息表
 */
type User struct {
	Id           int64     `xorm:"pk autoincr" json:"id"`        //主键 用户ID
	UserName     string    `xorm:"varchar(12)" json:"username"`  //用户名称
	RegisterTime time.Time `json:"register_time"`                //用户注册时间
	Mobile       string    `xorm:"varchar(11)" json:"mobile"`    //用户的移动手机号
	IsActive     int64     `json:"is_active"`                    //用户是否激活
	Balance      int64     `json:"balance"`                      //用户的账户余额（简单起见，使用int类型）
	Avatar       string    `xorm:"varchar(255)" json:"avatar"`   //用户的头像
	Pwd          string    `json:"password"`                     //用户的账户密码
	DelFlag      int64     `json:"del_flag"`                     //是否被删除的标志字段 软删除
	CityName     string    `xorm:"varchar(24)" json:"city_name"` //用户所在城市的名称
	City         *City     `xorm:"- <- ->"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (user *User) UserToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":           user.Id,
		"user_id":      user.Id,
		"username":     user.UserName,
		"city":         user.CityName,
		"registe_time": utils.FormatDatetime(user.RegisterTime),
		"mobile":       user.Mobile,
		"is_active":    user.IsActive,
		"balance":      user.Balance,
		"avatar":       user.Avatar,
	}
	return respInfo
}
