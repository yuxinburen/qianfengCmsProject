package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"irisDemo/CmsProject/service"
	"github.com/kataras/iris/sessions"
	"encoding/json"
	"irisDemo/CmsProject/utils"
	"irisDemo/CmsProject/model"
)

/**
 * 管理员控制器
 */
type AdminController struct {
	//iris框架自动为每个请求都绑定上下文对象
	Ctx iris.Context

	//admin功能实体
	Service service.AdminService

	//session对象
	Session *sessions.Session
}

const (
	ADMINTABLENAME = "admin"
	ADMIN          = "admin"
)

type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

/**
 * 管理员退出功能
 * 请求类型：Get
 * 请求url：admin/singout
 */
func (ac *AdminController) GetSingout() mvc.Result {

	//删除session，下次需要从新登录
	ac.Session.Delete(ADMIN);
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"success": utils.Recode2Text(utils.RESPMSG_SIGNOUT),
		},
	}
}

/**
 * 处理获取管理员总数的路由请求
 * 请求类型：Get
 * 请求Url：admin/count
 */
func (ac *AdminController) GetCount() mvc.Result {

	count, err := ac.Service.GetAdminCount()
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"message": utils.Recode2Text(utils.RESPMSG_ERRORADMINCOUNT),
				"count":   0,
			},
		}
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"count":  count,
		},
	}
}

/**
 * 获取管理员信息接口
 * 请求类型：Get
 * 请求url：/admin/info
 */
func (ac *AdminController) GetInfo() mvc.Result {

	//从session中获取信息
	userByte := ac.Session.Get(ADMIN)

	//session为空
	if userByte == nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.EEROR_UNLOGIN,
				"message": utils.Recode2Text(utils.EEROR_UNLOGIN),
			},
		}
	}

	//解析数据到admin数据结构
	var admin model.Admin
	err := json.Unmarshal(userByte.([]byte), &admin)

	//解析失败
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.EEROR_UNLOGIN,
				"message": utils.Recode2Text(utils.EEROR_UNLOGIN),
			},
		}
	}

	//解析成功
	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"data":   admin.AdminToRespDesc(),
		},
	}
}

/**
 * 管理员登录功能
 * 接口：/admin/login
 */
func (ac *AdminController) PostLogin(context iris.Context) mvc.Result {

	iris.New().Logger().Info(" admin login ")

	var adminLogin AdminLogin
	ac.Ctx.ReadJSON(&adminLogin)

	//数据参数检验
	if adminLogin.UserName == "" || adminLogin.Password == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "0",
				"success": "登录失败",
				"message": "用户名或密码为空,请重新填写后尝试登录",
			},
		}
	}

	//根据用户名、密码到数据库中查询对应的管理信息
	admin, exist := ac.Service.GetByAdminNameAndPassword(adminLogin.UserName, adminLogin.Password)

	//管理员不存在
	if !exist {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "0",
				"success": "登录失败",
				"message": "用户名或者密码错误,请重新登录",
			},
		}
	}

	//管理员存在 设置session
	userByte, _ := json.Marshal(admin)
	ac.Session.Set(ADMIN, userByte)

	return mvc.Response{
		Object: map[string]interface{}{
			"status":  "1",
			"success": "登录成功",
			"message": "管理员登录成功",
		},
	}

}
