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
			todos := v1.Group("/todos")
			{
				todos.POST("/category", h.createTodoCategory)
				todos.GET("/category", h.getAllTodoCategories)
				todos.GET("/category/:id", h.getTodoCategoryById)
				todos.DELETE("/category/:id", h.deleteTodoCategory)
				todos.PUT("/category/:id", h.updateTodoCategory)
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
