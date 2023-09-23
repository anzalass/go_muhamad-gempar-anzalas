package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------
// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserControllerById(c echo.Context) error {
	var id = c.Param("id")
	idv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Not Found data with id" + string(id),
		})
	}

	var user User
	found := false
	for _, val := range users {
		if val.Id == idv {
			user = val
			found = true
			break
		}

	}

	if !found {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Not Found data with id" + string(id),
		})
	}

	return c.JSON(http.StatusOK, user)
}

// // delete user by id
func DeleteUserController(c echo.Context) error {
	var id = c.Param("id")
	idv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Not Found data with id" + string(id),
		})
	}

	// Check if the user exists
	var user User
	found := false
	for _, u := range users {
		if u.Id == idv {
			user = u
			found = true
			break
		}
	}

	if !found {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message": "User with ID " + string(id) + " not found",
		})
	}

	// Delete the user
	for i, u := range users {
		if u.Id == idv {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}

	// Return a success message
	return c.JSON(http.StatusOK, map[string]any{
		"message": "User with ID " + string(id) + " deleted successfully",
		"user":    user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	var id = c.Param("id")
	idv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Not Found data with id" + string(id),
		})
	}

	var inputUpdate = User{}
	c.Bind(&inputUpdate)

	users[idv].Name = inputUpdate.Name
	users[idv].Email = inputUpdate.Email
	users[idv].Password = inputUpdate.Password

	return c.JSON(http.StatusOK, map[string]any{
		"message": "User with ID " + string(id) + " Updated successfully",
		"user":    users,
	})

}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)
	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()

	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserControllerById)
	e.GET("/users", GetUsersController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.Logger.Fatal(e.Start(":8000"))
}
