package controllers

import (
	"context"
	"fmt"
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

	//infinite processing will be end from here
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
func OrderById(c *gin.Context) {
	Id := c.Param("id")
	smt := "select id, user_id, product_id, quantity, order_date from orders where id =$1"
	result := database.DB.QueryRow(context.Background(), smt, Id)
	var order models.Order
	err := result.Scan(&order.Id, &order.User_Id, &order.Product_Id, &order.Quantity, &order.Order_Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, order)
}
func CreateOrder(c *gin.Context) {

	var order models.Order
	err := c.ShouldBindJSON(&order)

	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	smt := "insert into orders (id ,user_id ,product_id ,quantity,order_date ) values($1,$2,$3,$4,$5)"
	result, err := database.DB.Exec(context.Background(), smt, &order.Id, &order.User_Id, &order.Id, &order.Product_Id, &order.Quantity, &order.Order_Date)
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, result)

}
