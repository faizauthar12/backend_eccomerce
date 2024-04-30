package models

import (
	"github.com/faizauthar12/eccomerce/global-utils/model"
)

type CartItem struct {
	UUID     string  `json:"uuid" bson:"uuid"`
	Price    float64 `json:"price" bson:"price"`
	Quantity int64   `json:"quantity" bson:"quantity"`
}

type Cart struct {
	UUID      string     `json:"uuid" bson:"uuid"`
	UserUUID  string     `json:"user_uuid" bson:"user_uuid"`
	CartItems []CartItem `json:"cart_items" bson:"cart_items"`
	CreatedAt int64      `json:"created_at" bson:"created_at"`
	UpdatedAt int64      `json:"updated_at" bson:"updated_at"`
}

type CartChan struct {
	Cart     *Cart           `json:"cart" bson:"cart"`
	Total    int64           `json:"total" bson:"total"`
	Error    error           `json:"error" bson:"error"`
	ErrorLog *model.ErrorLog `json:"error_log" bson:"error_log"`
}

type CartUpdateRequest struct {
	UUID     string `json:"uuid" bson:"uuid"`
	UserUUID string `json:"user_uuid" bson:"user_uuid"`
}
