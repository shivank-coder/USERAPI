package models

import "time"

type Order struct {
	Id         int       `json:"id"`
	User_Id    int       `json:"user_id"`
	Product_Id int       `json:"product_id"`
	Quantity   int       ` json:"quantity"`
	Order_Date time.Time `json:" order_date"`
}
