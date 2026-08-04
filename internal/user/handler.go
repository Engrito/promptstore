package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	handlers Service
}

func NewUserHandlers(service Service) *UserHandlers {
	return &UserHandlers{
		handlers: service,
	}
}

func (h *UserHandlers) Singup(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.handlers.Add(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Id = id
	ctx.JSON(http.StatusCreated, user)
}

func (h *UserHandlers) Login(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.handlers.GetById(ctx, user.Id)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	ctx.JSON(http.StatusOK, u)
}

func (h *UserHandlers) GetById(ctx *gin.Context) {
	id := ID(ctx.Query("id"))
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	user, err := h.handlers.GetById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandlers) GetAll(ctx *gin.Context) {
	users, err := h.handlers.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (h *UserHandlers) UpdateByID(ctx *gin.Context) {
	id := ID(ctx.Query("id"))
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.handlers.UpdateByID(ctx, id, user); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user updated"})
}

func (h *UserHandlers) DeleteByID(ctx *gin.Context) {
	id := ID(ctx.Query("id"))
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	if err := h.handlers.DeleteByID(ctx, id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
