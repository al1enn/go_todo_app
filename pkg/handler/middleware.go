package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	ctx.Set(userCtx, userId)
}

func getUserId(ctx *gin.Context) (int, error) {
	userId, ok := ctx.Get(userCtx)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}
	idInt, ok := userId.(int)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user id is not an integer")
		return 0, errors.New("user id is not an integer")
	}
	return idInt, nil
}
