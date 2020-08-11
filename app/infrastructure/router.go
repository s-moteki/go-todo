package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/s-moteki/go-todo/app/interfaces/controllers"
)

// Router ルータ
var Router *gin.Engine

func init() {
	router := gin.Default()

	todoController := controllers.NewTodoController(NewSQLHandler())
	router.GET("/todos", func(c *gin.Context) { todoController.Index(c) })
	router.GET("/todos/:id", func(c *gin.Context) { todoController.Show(c) })
	router.GET("/search", func(c *gin.Context) { todoController.Search(c) })
	router.POST("/todos", func(c *gin.Context) { todoController.Create(c) })
	router.PATCH("/todos/:id", func(c *gin.Context) { todoController.Edit(c) })

	Router = router
}
