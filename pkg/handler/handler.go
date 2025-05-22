package handler

import (
	"net/http"

	"github.com/al1enn/go_todo_app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/ping", h.ping)
			todo_categories := v1.Group("/todos")
			{
				todo_categories.POST("/category", h.createTodoCategory)
				todo_categories.GET("/category", h.getAllTodoCategories)
				todo_categories.GET("/category/:id", h.getTodoCategoryById)
				todo_categories.DELETE("/category/:id", h.deleteTodoCategory)
				todo_categories.PUT("/category/:id", h.updateTodoCategory)

			}
			todo_items := v1.Group("/items")
			{
				todo_items.POST("/", h.createTodoItem)
				// todo_items.GET("/", h.getAllTodoItems)
				// todo_items.GET("/:id", h.getTodoItemById)
				// todo_items.DELETE("/:id", h.deleteTodoItem)
				// todo_items.PUT("/:id", h.updateTodoItem)
			}
		}
	}

	return router
}

func (h *Handler) ping(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "pong",
	})
}
