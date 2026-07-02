package user

import (
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
	panic("yet to implement")
}

func (h *UserHandlers) Login(ctx *gin.Context) {
	panic("yet to implement")
}

func (h *UserHandlers) GetById(ctx *gin.Context) {
	panic("yet to implement")
}

func (h *UserHandlers) GetAll(ctx *gin.Context) {
	panic("yet to implement")
}

func (h *UserHandlers) UpdateByID(ctx *gin.Context) {
	panic("yet to implement")
}

func (h *UserHandlers) DeleteByID(ctx *gin.Context) {
	panic("yet to implement")
}
