package routes

import (
	"testinggg/controller"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uc controller.UserController) {
	e.POST("/users", uc.CreateUser())
	e.POST("/login", uc.LoginUser())
	e.GET("/users/:id", uc.GetUserByID())
	e.PUT("/users/:id", uc.UpdateUserByID())
	e.DELETE("/users/:id", uc.DeleteUserByID())
}
