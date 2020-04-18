package app

import (
	"fmt"
	"getzoop/zec-companion-app/controllers"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
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
	app.iris.Logger().SetLevel("debug")
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
			l := fmt.Sprintf("->%s [%s] %s - %s %s",
				now.Format("2006/01/02 - 15:04:05"),
				ip,
				status,
				method,
				path)
			app.iris.Logger().Info(l)
		},
		Skippers: nil,
	}))

	app.createRoutes()
}

func (app *application) RunServer(port int) error {
	return app.iris.Run(iris.Addr(fmt.Sprintf(":%d", port)), iris.WithoutServerError(iris.ErrServerClosed))
}

func (app *application) createRoutes() {
	healthController := controllers.HealthController()
	app.iris.Get("/health", healthController.Get)

	app.Logger().Infof("%+V", app)
	app.Logger().Debugf("%+V", app)

	app.Logger().Info("qweqwe")
	app.Logger().Debug("qweqwe")

	app.Logger().Info(app.iris)
	app.Logger().Debug(app.iris)
}

func (app *application) Logger() *golog.Logger {
	return app.iris.Logger()
}

func Logger() *golog.Logger {
	return App().Logger()
}
