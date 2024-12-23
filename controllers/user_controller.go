package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"go-practice-app/database"
	"go-practice-app/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	rows, err := database.DB.Query(context.Background(), "SELECT id, name, email, created_at FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		users = append(users, user)
	}
	c.JSON(http.StatusOK, users)

}
func CreateUser(c *gin.Context) {
	// Binding incoming JSON to a User struct
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		// Invalid input
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// SQL query to insert user into the database
	smt := "INSERT INTO users (name, email, created_at) VALUES ($1, $2, $3)"

	// Execute the query with context and the user's details
	_, err := database.DB.Exec(context.Background(), smt, user.Name, user.Email, user.CreatedAt)
	if err != nil {
		// If there's an error, respond with a failure message
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		log.Printf("Error executing query: %v", err)
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "User is created"})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	smt := "SELECT id, name, email, created_at FROM users WHERE id = $1"

	row := database.DB.QueryRow(context.Background(), smt, id)
	var user models.User

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to retrieve user: %v", err)})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}
