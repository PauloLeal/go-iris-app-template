package app

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
	"testing"
)

func TestCreateServer(t *testing.T) {
	app := App()
	e := httptest.New(t, app.iris)

	e.GET("/health").Expect().Status(httptest.StatusOK).JSON().Equal(iris.Map{"status": "ok"})
}
