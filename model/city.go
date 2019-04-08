package model

/*
 * 城市基础结构体字段(基础表)
 */
type City struct {
	CityId    int64   `xorm:"pk autoincr" json:"id"`       //主键
	CityName  string  `xorm:"varchar(12)" json:"name"`     //城市名称
	PinYin    string  `xorm:"varchar(32)" json:"pin_yin"`  //城市拼音
	Longitude float32 `xorm:"default 0" json:"longitude"`  //城市经度
	Latitude  float32 `xorm:"default 0" json:"latitude"`   //城市纬度
	AreaCode  string  `xorm:"varchar(6)" json:"area_code"` //城市的地区编码
	Abbr      string  `xorm:"varchar(12)" json:"abbr"`     //城市的拼音缩写
}
