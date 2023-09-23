package controller

import (
	"echodatabase/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	err := model.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success created",
		"user":    user,
	})

}

func GetAllUsers(c echo.Context) error {
	var users []model.User

	err := model.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  users,
		"total": len(users),
	})

}

func GetUserById(c echo.Context) error {
	id := c.Param("id")
	idv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	var user model.User
	res := model.DB.First(&user, idv).Error
	if res != nil {
		return c.JSON(http.StatusNoContent, map[string]interface{}{
			"message": "gak ada",
		})
	}

	return c.JSON(http.StatusOK, user)
}

func GetUserNBookById(c echo.Context) error {
	id := c.Param("id")
	idv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	rows, err := model.DB.Table("users").Select("users.name, books.judul").Joins("left join books on books.penulis = users.id").Where("users.id =?", idv).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	defer rows.Close()

	var resultss []struct {
		Name  string
		Judul string
	}

	var penulis string

	for rows.Next() {
		var result struct {
			Name  string
			Judul string
		}

		err := rows.Scan(&result.Name, &result.Judul)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err,
			})
		}
		penulis = result.Name
		resultss = append(resultss, result)
	}

	if len(resultss) == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"penulis": penulis,
		"data":    resultss,
	})

}
func GetUserNBlogById(c echo.Context) error {
	id := c.Param("id")
	idv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	rows, err := model.DB.Table("users").Select("users.name, blogs.judul_blog").Joins("left join blogs on blogs.id_user = users.id").Where("users.id =?", idv).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	defer rows.Close()

	var resultss []struct {
		Name  string
		Judul string
	}

	var penulis string

	for rows.Next() {
		var result struct {
			Name  string
			Judul string
		}

		err := rows.Scan(&result.Name, &result.Judul)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err,
			})
		}
		penulis = result.Name
		resultss = append(resultss, result)
	}

	if len(resultss) == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"penulis": penulis,
		"data":    resultss,
	})

}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	idv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	user := model.User{}
	c.Bind(&user)

	// user.ID = uint64(idv)
	err = model.DB.Model(&user).Where("id = ?", idv).Updates(map[string]interface{}{"name": user.Name, "email": user.Email}).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Updated user",
		"data":    user,
	})
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	idv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	user := model.User{}
	err = model.DB.Delete(&user, idv).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Deleted successfully",
	})
}
