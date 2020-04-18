package controllers

import (
	"getzoop/zec-companion-app/models"
	"github.com/kataras/iris"
)

func sendStatus(ctx iris.Context, apiError models.ApiStatus) {
	ctx.StatusCode(apiError.Status().Code)
	_, _ = ctx.JSON(iris.Map{"status": apiError.Status().Code, "message": apiError.Status().Message})
}
