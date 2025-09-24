package httpadapter

import (
	"net/http"

	"github.com/FatwahFir/xpanca-be/internal/dto"
	"github.com/FatwahFir/xpanca-be/internal/usecase"
	"github.com/FatwahFir/xpanca-be/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{ uc *usecase.AuthUsecase }

func NewAuthHandler(r gin.IRouter, uc *usecase.AuthUsecase) {
	h := &AuthHandler{uc: uc}
	r.POST("/auth/login", h.Login)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Err(c, http.StatusBadRequest, response.CodeBadRequest, "invalid payload", nil)
		return
	}
	token, u, err := h.uc.Login(c, req.Username, req.Password)
	if err != nil {
		if err == usecase.ErrInvalidCredentials {
			response.Err(c, http.StatusUnauthorized, response.CodeUnauthorized, "invalid credentials", nil)
			return
		}
		response.Err(c, http.StatusInternalServerError, response.CodeServerError, "cannot issue token", nil)
		return
	}
	response.OK(c, gin.H{"token": token, "user": u})
}
