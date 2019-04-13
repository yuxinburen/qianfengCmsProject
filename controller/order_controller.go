package controller

import (
	"context"
	"irisDemo/CmsProject/service"
)

type OrderController struct {
	Ctx context.Context
	Service service.OrderService
}
