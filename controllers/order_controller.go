package controllers

import (
	"context"
	"go-practice-app/database"
	"go-practice-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	smt := "SELECT * FROM orders"

	// Execute the query
	rows, err := database.DB.Query(context.Background(), smt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch orders"})
		return
	}
	defer rows.Close()

	var orders []models.Order

	// Iterate through the rows
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.Id, &order.User_Id, &order.Product_Id, &order.Quantity, &order.Order_Date); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to scan order"})
			return
		}
		orders = append(orders, order)
	}

	// Check for errors after iteration
	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred during rows iteration"})
		return
	}

	// Return the orders
	c.JSON(http.StatusOK, orders)
}
