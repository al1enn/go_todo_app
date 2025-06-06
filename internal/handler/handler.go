package handler

import (
	_ "github.com/al1enn/go_todo_app/docs"
	"github.com/al1enn/go_todo_app/internal/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	auth := router.Group("/auth")

	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		v1 := api.Group("/v1")
		{
			todo := v1.Group("/todo")
			todoCategories := todo.Group("/category")
			{
				todoCategories.POST("/", h.createTodoCategory)
				todoCategories.GET("/", h.getAllTodoCategories)
				todoCategories.GET("/:id", h.getTodoCategoryById)
				todoCategories.DELETE("/:id", h.deleteTodoCategory)
				todoCategories.PUT("/:id", h.updateTodoCategory)
				todoItems := todoCategories.Group("/:id/item")
				{
					todoItems.POST("", h.createTodoItem)
				}
			}
			todo_items := todo.Group("/item")
			{
				todo_items.GET("", h.getAllTodoItems)
				todo_items.GET("/:id", h.getTodoItemById)
				todo_items.DELETE("/:id", h.deleteTodoItem)
				todo_items.PUT("/:id", h.updateTodoItem)
			}
		}
	}

	return router
}
