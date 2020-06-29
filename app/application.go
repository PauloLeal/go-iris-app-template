package app

import (
	"fmt"
	"github.com/PauloLeal/go-iris-app-template/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var appInstance *application
var appOnce sync.Once

type application struct {
	iris *iris.Application
}

func App() *application {
	appOnce.Do(func() {
		appInstance = &application{iris: iris.New()}
		appInstance.initialize()
	})
	return appInstance
}

func (app *application) initialize() {
	app.iris.Logger().SetLevel("info")
	app.iris.Use(recover.New())
	app.iris.Use(logger.New(logger.Config{
		Status:             true,
		IP:                 true,
		Method:             true,
		Path:               true,
		Query:              true,
		Columns:            false,
		MessageContextKeys: nil,
		MessageHeaderKeys:  nil,
		LogFunc: func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
			l := fmt.Sprintf("ACCESS [%s] %s - %s %s",
				ip,
				status,
				method,
				path)
			logrus.Info(l)
		},
		Skippers: nil,
	}))

	app.createRoutes()
}

func (app *application) RunServer(port int) error {
	return app.iris.Run(iris.Addr(fmt.Sprintf(":%d", port)), iris.WithoutServerError(iris.ErrServerClosed))
}

func (app *application) createRoutes() {
	healthController := controllers.NewHealthController()
	app.iris.Get("/health", healthController.Get)

	app.iris.Handle("ALL", "/*", func(ctx iris.Context) {
		ctx.StatusCode(404)
	})

}
