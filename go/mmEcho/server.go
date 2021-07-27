package main

import (
	"app/router"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	router.LoginRouter(e)
	router.IndexIndexRouter(e)
	router.IndexAdminRouter(e)
	router.ForumIndexRouter(e)

	router.ForumAdminRouter(e)
	/*
		g := e.Group("/guest")
		g.GET("/index", control.GuestIndex)
		g.GET("/show", control.GuestShow)
		u := e.Group("/user")
		u.GET("/index", control.UserIndex)
		u.GET("/login", control.UserLogin)
		u.GET("/register", control.UserReg)
	*/
	e.Logger.Fatal(e.Start(":1323"))
}

func Filter(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 路由拦截 - 登录身份、资源权限判断等
		println("Api路由拦截：", c.Path())
		return next(c)
	}
}
