package routes

import (
	"echodatabase/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/users", controller.CreateUser)
	e.GET("/users/:id", controller.GetUserById)
	e.GET("/booksuser/:id", controller.GetUserNBookById)
	e.GET("/blogsuser/:id", controller.GetUserNBlogById)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)
	e.GET("/users", controller.GetAllUsers)

	e.POST("/books", controller.PublishBook)
	e.GET("/books", controller.GetAllBooks)
	e.PUT("/books/:id", controller.UpdateBookById)
	e.DELETE("/books/:id", controller.DeleteBooksById)
	e.GET("/books/:id", controller.GetBookById)

	e.POST("/blogs", controller.PublishArtikel)
	e.GET("/blogs", controller.GetAllBlog)
	e.PUT("/blogs/:id", controller.UpdateBlogById)
	e.DELETE("/blogs/:id", controller.DeleteBlogById)
	e.GET("/blogs/:id", controller.GetBlogById)

	return e
}
