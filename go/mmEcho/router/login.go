package router

import (
	"app/index/indexAdmin"

	"github.com/labstack/echo/v4"
)

func LoginRouter(e *echo.Echo) {
	e.GET("/admin/login/index", indexAdmin.LoginIndex)
	e.POST("/admin/login/loginsave", indexAdmin.LoginSave)
}
