package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var loggedInUser *User

func main() {
	router := gin.Default()

	// Login route
	router.POST("/login", loginHandler)

	// Logout route
	router.POST("/logout", logoutHandler)

	// Account route
	router.GET("/account", accountHandler)

	router.Run(":8080")
}

func loginHandler(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Here you would typically check the credentials against a database
	// For demonstration purposes, we'll just check if the username and password are "admin"
	if user.Username == "admin" && user.Password == "admin" {
		loggedInUser = &user
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

func logoutHandler(c *gin.Context) {
	loggedInUser = nil
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func accountHandler(c *gin.Context) {
	if loggedInUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not logged in"})
		return
	}

	// For demonstration purposes, we'll just return the logged-in user's details
	c.JSON(http.StatusOK, loggedInUser)
}
