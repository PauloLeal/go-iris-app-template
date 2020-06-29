package controllers

import (
	"github.com/kataras/iris"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (ctrl *HealthController) Get(ctx iris.Context) {
	ctx.StatusCode(200)
	_, _ = ctx.JSON(iris.Map{"status": "ok"})
}
