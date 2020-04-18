package controllers

import (
	"github.com/kataras/iris"
	"sync"
)

type healthController struct {
}

var healthControllerInstance *healthController
var healthControllerOnce sync.Once

func HealthController() *healthController {
	healthControllerOnce.Do(func() {
		healthControllerInstance = &healthController{}
		healthControllerInstance.initialize()
	})
	return healthControllerInstance
}

func (l *healthController) initialize() {

}

func (l *healthController) Get(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{"status": "ok"})
}
