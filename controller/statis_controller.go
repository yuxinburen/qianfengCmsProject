package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"irisDemo/CmsProject/service"
	"strings"
	"github.com/kataras/iris/mvc"
	"irisDemo/CmsProject/utils"
)

/**
 * 统计功能控制者
 */
type StatisController struct {
	//上下文环境对象
	Ctx iris.Context

	//统计功能的服务实现接口
	Service service.StatisService

	//session
	Session *sessions.Session
}

/**
 * 解析统计功能路由请求
 */
func (sc *StatisController) GetCount() mvc.Result {
	// /statis/user/2019-03-10/count
	path := sc.Ctx.Path()

	var pathSlice []string
	if path != "" {
		pathSlice = strings.Split(path, "/")
	}

	//不符合请求格式
	if len(pathSlice) != 5 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": utils.RECODE_FAIL,
				"count":  0,
			},
		}
	}

	//将最前面的去掉
	pathSlice = pathSlice[1:]
	model := pathSlice[1]
	date := pathSlice[2]
	var result int64
	switch model {
	case "user":
		iris.New().Logger().Error(date) //时间
		result = sc.Service.GetUserDailyCount(date)
	case "order":
		result = sc.Service.GetOrderDailyCount(date)
	case "admin":
		result = sc.Service.GetAdminDailyCount(date)
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"count":  result,
		},
	}
}
