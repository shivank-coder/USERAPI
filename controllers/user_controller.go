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

// delete user by its id
func DeleteById(c *gin.Context) {
	Id := c.Param("id")

	// Check for related orders
	var count int
	err := database.DB.QueryRow(context.Background(), "SELECT COUNT(*) FROM orders WHERE user_id=$1", Id).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check related orders"})
		return
	}
	//means altleast some orders are connected with orders id
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "cannot delete user with associated orders"})
		return
	}

	// Proceed with deletion
	smt := "DELETE from users where id=$1"
	result, err := database.DB.Exec(context.Background(), smt, Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while deleting the user"})
		return
	}
	//check if there is rows effected or not by this operation
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no user found with the given id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully", "rows_affected": rowsAffected})
}
