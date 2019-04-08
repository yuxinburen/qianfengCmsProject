package model

/**
 * 商店实体结构体定义
 */
type Shop struct {
	Id                          int        `xorm:"pk autoincr" json:"item_id"`      //店铺Id
	Name                        string     `xorm:"varchar(32)" json:"name"`         //店铺名称
	Address                     string     `xorm:"varchar(128)" json:"address"`     //店铺地址
	Latitude                    float32    `json:"latitude"`                        //经度
	Longitude                   float32    `json:"longitude"`                       //纬度
	Description                 string     `xorm:"varchar(255)" json:"description"` //店铺简介
	Phone                       string     `json:"phone"`                           //店铺电话
	PromotionInfo               string     `json:"promotion_info"`                  //店铺标语
	FloatDeliveryFee            int        `json:"float_delivery_fee"`              //配送费
	FloatMinimumOrderAmount     int        `json:"float_minimum_order_amount"`      //起送价
	IsPremium                   bool       `json:"is_premium"`                      //品牌保障
	DeliveryMode                bool       `json:"delivery_mode"`                   //蜂鸟专送
	New                         bool       `json:"new"`                             //新开店铺
	Bao                         bool       `json:"bao"`                             //外卖保
	Zhun                        bool       `json:"zhun"`                            //准时达
	Piao                        bool       `json:"piao"`                            //开发票
	StartTime                   string     `json:"startTime"`                       //营业开始时间
	EndTime                     string     `json:"endTime"`                         //营业结束时间
	ImagePath                   string     `json:"image_path"`                      //店铺头像
	BusinessLicenseImage        string     `json:"business_license_image"`          //营业执照
	CateringServiceLicenseImage string     `json:"catering_service_license_image"`  //餐饮服务许可证
	Category                    string     `json:"category"`                        //店铺类型
	Status                      int        `json:"status"`                          //店铺状态
	RecentOrderNum              int        `json:"recent_order_num"`                //最近一个月的销量
	RatingCount                 int        `json:"rating_count"`                    //评分次数
	Rating                      int        `json:"rating"`                          //综合评分
	Dele                        int        `json:"dele"`                            //是否已经被删除 1表示已经删除 0表示未删除
	Activities                  []*Service `xorm:"-"`                               //商家提供的服务 结构体
}
