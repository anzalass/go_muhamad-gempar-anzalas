package controller

import (
	"net/http"
	"strconv"
	"testinggg/model"

	"github.com/labstack/echo/v4"
)

type UserControllerInterface interface {
	CreateUser() echo.HandlerFunc
	LoginUser() echo.HandlerFunc
	GetAllUser() echo.HandlerFunc
	GetUserByID() echo.HandlerFunc
	DeleteUserByID() echo.HandlerFunc
	UpdateUserByID() echo.HandlerFunc
}

type UserController struct {
	model model.UserInterface
}

func NewUserModelControllerInterface(m model.UserInterface) UserControllerInterface {
	return &UserController{
		model: m,
	}
}

func (uc *UserController) InitUserController(um model.UserModel) {
	uc.model = &um
}

func (uc *UserController) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = model.User{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "Invalid user input",
				"msg":     "Invalid user input",
			})
		}
		var res = uc.model.CreateUser(input)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "Error creating user",
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "creating user",
			"data":    res,
		})

	}

}

func (uc *UserController) LoginUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = model.User{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "invalid user",
			})
		}

		var res = uc.model.LoginUser(input)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "cant process login",
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "sukses",
			"data":    res,
		})
	}
}

func (uc *UserController) GetAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var res = uc.model.GetAllUser()
		return c.JSON(http.StatusOK, map[string]any{
			"message": "sukses",
			"data":    res,
		})

	}
}

func (uc *UserController) GetUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")
		cnv, err := strconv.Atoi(paramId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": err,
			})
		}

		var res = uc.model.GetUserByID(uint64(cnv))
		if res == nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "error getting user",
			})
		}
		if res.Name == "" {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "Not found user",
			})
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uc *UserController) DeleteUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")
		cnv, err := strconv.Atoi(paramId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "id not found",
			})
		}
		uc.model.DeleteUserByID(uint64(cnv))

		return c.JSON(http.StatusOK, map[string]any{
			"message": "Delete User",
		})
	}
}

func (uc *UserController) UpdateUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")
		cnv, err := strconv.Atoi(paramId)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "id invalid",
			})
		}
		var input = model.User{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "update invalid",
			})
		}
		input.ID = uint64(cnv)
		var res = uc.model.UpdateUserByID(input)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "invalid",
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "sukses update",
			"data":    res,
		})
	}

}
