package handler

import (
	"net/http"

	todo "github.com/al1enn/go_todo_app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createTodoItem(ctx *gin.Context) {
	_, ok := ctx.Get(userCtx)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return
	}
	var input todo.TodoItem
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	//call services
}

func (h *Handler) getTodoItem(ctx *gin.Context) {
	_, ok := ctx.Get(userCtx)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return
	}
}
