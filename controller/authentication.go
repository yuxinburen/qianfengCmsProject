package controller

import "github.com/kataras/iris"

/**
 * 认证
 */
func Authentication(context iris.Context) {
	context.Next()
}
