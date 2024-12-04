package main

import (
	config "w2/d3/config/database"
	cust_middleware "w2/d3/internal/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"w2/d3/internal/userHandler"
)

func main() {
	config.InitDB()
	defer config.CloseDB()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// public routes
	e.POST("/login", handler.Login)
	e.POST("/register", handler.Register)

	// protected routes
	u := e.Group("/users")
	u.Use(cust_middleware.JwtMiddleware)

	// u := e.Group("/users")
	// u.POST("/", handler.CreateUser)
	// u.GET("/", handler.Get)

	e.Logger.Fatal(e.Start(":8080"))

}
