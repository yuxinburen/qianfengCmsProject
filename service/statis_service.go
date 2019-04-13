package service

import (
	"github.com/go-xorm/xorm"
	"irisDemo/CmsProject/model"
	"time"
)

/**
 * 统计功能模块接口标准
 */
type StatisService interface {
	//查询某一天的用户的增长数量
	GetUserDailyCount(date string) int64
	GetOrderDailyCount(date string) int64
	GetAdminDailyCount(date string) int64
}

/**
 * 统计功能服务实现结构体
 */
type statisService struct {
	Engine *xorm.Engine
}

/**
 * 新建统计模块功能服务对象
 */
func NewStatisService(engine *xorm.Engine) StatisService {
	return &statisService{
		Engine: engine,
	}
}

/**
 * 查询某一日管理员的增长数量
 */
func (ss *statisService) GetAdminDailyCount(date string) int64 {
	//查询如期date格式解析
	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}

	endDate := startDate.AddDate(0, 0, 1)
	result, err := ss.Engine.Where(" create_time between ? and ? and status = 0 ", startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05")).Count(model.Admin{})
	if err != nil {
		return 0
	}

	return result
}

/**
 * 查询某一日订单的单日增长数量
 */
func (ss *statisService) GetOrderDailyCount(date string) int64 {

	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}

	endDate := startDate.AddDate(0, 0, 1)
	result, err := ss.Engine.Where(" time between ? and ? and del_flag = 0 ", startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05")).Count(model.UserOrder{})
	if err != nil {
		return 0
	}

	return result
}

/**
 * 查询某一日用户的单日增长数量
 */
func (ss *statisService) GetUserDailyCount(date string) int64 {

	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}

	endDate := startDate.AddDate(0, 0, 1)
	result, err := ss.Engine.Where(" register_time between ? and ? and del_flag = 0 ", startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05")).Count(model.User{})
	if err != nil {
		return 0
	}

	return result
}
