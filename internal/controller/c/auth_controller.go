package controller

import (
	dto "SJTU-Canteen-Community/internal/dto/c"
	"SJTU-Canteen-Community/internal/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *service.AuthService
}

func NewAuthController(AuthService *service.AuthService) *AuthController {
	return &AuthController{AuthService: AuthService}
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginReq dto.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID, err := ac.AuthService.Login(loginReq)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", userID)
	if err := session.Save(); err != nil {
		c.JSON(500, gin.H{"error": "Failed to save session"})
		return
	}

	c.JSON(200, gin.H{"message": "Login successful"})

}

func (ac *AuthController) TestLoginStatus(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.JSON(401, gin.H{"error": "Not logged in"})
		return
	}
	c.JSON(200, gin.H{"message": "Logged in", "user_id": userID})
}

func (ac *AuthController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		c.JSON(500, gin.H{"error": "Failed to clear session"})
		return
	}
	c.JSON(200, gin.H{"message": "Logout successful"})

}
