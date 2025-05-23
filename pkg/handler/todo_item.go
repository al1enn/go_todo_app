package handler

import (
	"net/http"
	"strconv"

	todo "github.com/al1enn/go_todo_app"
	"github.com/gin-gonic/gin"
)

type getAllTodoItemsResponse struct {
	Data  []todo.TodoItem `json:"data"`
	Total int             `json:"total"`
}

func (h *Handler) createTodoItem(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	categoryId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}
	var input todo.TodoItem
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	h.services.TodoItem.Create(userId, categoryId, input)
	//call services
}

func (h *Handler) getAllTodoItems(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	todoItems, err := h.services.TodoItem.GetAll(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, getAllTodoItemsResponse{
		Data:  todoItems,
		Total: len(todoItems),
	})
}

// func (h *Handler) getTodoItemById(ctx *gin.Context) {
// 	userId, err := getUserId(ctx)
// 	if err != nil {
// 		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	itemId, err := strconv.Atoi(ctx.Param("id"))
// 	if err != nil {
// 		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
// 		return
// 	}
// 	todoItem, err := h.services.TodoItem.GetById(userId, itemId)
// 	if err != nil {
// 		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, map[string]interface{}{
// 		"data": todoItem,
// 	})

// }
