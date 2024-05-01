package models

import "github.com/faizauthar12/eccomerce/global-utils/model"

type OrderItem struct {
	ProductUUID string  `json:"product_uuid" bson:"product_uuid"`
	Price       float64 `json:"price" bson:"price"`
	Quantity    int64   `json:"quantity" bson:"quantity"`
	Total       float64 `json:"total" bson:"total"`
}

type Order struct {
	UUID      string      `json:"uuid" bson:"uuid"`
	UserUUID  string      `json:"user_uuid" bson:"user_uuid"`
	OrderItem []OrderItem `json:"order_items" bson:"order_items"`
	Total     float64     `json:"total" bson:"total"`
	Address   string      `json:"address" bson:"address"`
	CreatedAt int64       `json:"created_at" bson:"created_at"`
	UpdatedAt int64       `json:"updated_at" bson:"updated_at"`
}

type OrderChan struct {
	Order      *Order          `json:"order" bson:"order"`
	Total      int64           `json:"total" bson:"total"`
	Error      error           `json:"error" bson:"error"`
	ErrorLog   *model.ErrorLog `json:"error_log" bson:"error_log"`
	UUID       string          `json:"uuid" bson:"uuid"`
	StatusCode int             `json:"status_code" bson:"status_code"`
}

type OrdersChan struct {
	Orders     []*Order        `json:"orders" bson:"orders"`
	Total      int64           `json:"total" bson:"total"`
	Error      error           `json:"error" bson:"error"`
	ErrorLog   *model.ErrorLog `json:"error_log" bson:"error_log"`
	UUID       string          `json:"uuid" bson:"uuid"`
	StatusCode int             `json:"status_code" bson:"status_code"`
}

type OrderRequest struct {
	UserUUID string `json:"user_uuid" bson:"user_uuid"`
	NumItems int64  `json:"num_items" bson:"num_items"`
	Pages    int64  `json:"pages" bson:"pages"`
}

type Orders struct {
	Orders []*Order `json:"orders" bson:"orders"`
	Total  int64    `json:"total" bson:"total"`
}
