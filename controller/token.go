package controller

import (
	"dans-multi-pro-test/auth"
	"dans-multi-pro-test/database"
	"dans-multi-pro-test/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GenerateToken(c *gin.Context) {
	var request TokenRequest
	var user entity.User

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	row := database.DB.QueryRow("SELECT id, username, password FROM user WHERE username=?", user.Username)
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username or Password"})
		return
	}

	err = user.CheckPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		c.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
