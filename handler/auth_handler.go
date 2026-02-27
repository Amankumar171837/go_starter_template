package handler

import (
	"net/http"
	"strconv"

	"go_starter_template/domain"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service domain.AuthService
}

func NewAuthHandler(service domain.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) setAuthHeaders(c *gin.Context, tokens *domain.TokenResponse) {
	c.Header("access-token", tokens.AccessToken)
	c.Header("refresh-token", tokens.RefreshToken)
	c.Header("access-expire", strconv.FormatInt(tokens.AccessExpiry, 10))
	c.Header("refresh-expire", strconv.FormatInt(tokens.RefreshExpiry, 10))
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req domain.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.service.Register(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.setAuthHeaders(c, tokens)
	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.service.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	h.setAuthHeaders(c, tokens)
	c.JSON(http.StatusOK, gin.H{"Status": "200"})
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	var req domain.RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.service.Refresh(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	h.setAuthHeaders(c, tokens)
	c.JSON(http.StatusOK, gin.H{"message": "Refresh successful"})
}
