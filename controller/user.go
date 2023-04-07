package controller

import (
	"dans-multi-pro-test/auth"
	"dans-multi-pro-test/database"
	"dans-multi-pro-test/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user entity.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	_, err := database.DB.Exec("INSERT INTO user (username, password) VALUES (?, ?);", user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"userId": user.ID, "username": user.Username})
}

func Login(c *gin.Context) {
	var input entity.User
	var user entity.User

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	row := database.DB.QueryRow("SELECT id, username, password FROM user WHERE username=?", input.Username)
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username"})
		return
	}

	err = user.CheckPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		c.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error when generate jwt"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"userId": user.ID, "username": user.Username, "tokenString": tokenString})
}
