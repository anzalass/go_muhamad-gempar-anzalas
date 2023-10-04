package main

import (
	"testinggg/controller"
	"testinggg/model"
	"testinggg/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := model.InitModel()
	model.Migrate(db)

	userModel := model.UserModel{}
	userModel.Init(db)
	userController := controller.UserController{}
	userController.InitUserController(userModel)

	routes.RouteUser(e, userController)

	e.Logger.Fatal(e.Start(":8000").Error())

}
