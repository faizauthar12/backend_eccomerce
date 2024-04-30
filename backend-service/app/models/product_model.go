package models

import (
	"github.com/faizauthar12/eccomerce/global-utils/model"
)

type Product struct {
	UUID      string `json:"uuid" bson:"uuid"`
	Name      string `json:"name" bson:"name"  binding:"required"`
	Price     int64  `json:"price" bson:"price"`
	Stock     int64  `json:"stock" bson:"stock" binding:"required"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at"`
}

type ProductChan struct {
	Product    *Product        `json:"product"`
	Total      int64           `json:"total"`
	Error      error           `json:"error"`
	ErrorLog   *model.ErrorLog `json:"error_log"`
	UUID       string          `json:"uuid"`
	StatusCode int             `json:"status_code"`
}

type ProductsChan struct {
	Products   []*Product      `json:"products"`
	Total      int64           `json:"total"`
	Error      error           `json:"error"`
	ErrorLog   *model.ErrorLog `json:"error_log"`
	UUID       string          `json:"uuid"`
	StatusCode int             `json:"status_code"`
}

type ProductRequest struct {
	NumItems int64  `json:"num_items" bson:"num_items"`
	Pages    int64  `json:"pages" bson:"pages"`
	Name     string `json:"name" bson:"name"`
	PriceGte int64  `json:"price_minimum" bson:"price_minimum"`
	PriceLte int64  `json:"price_maximum" bson:"price_maximum"`
}

type ProductUpdateRequest struct {
	Name  string `json:"name" bson:"name"`
	Price int64  `json:"price" bson:"price"`
	Stock int64  `json:"stock" bson:"stock"`
}

type Products struct {
	Products []*Product `json:"products"`
	Total    int64      `json:"total"`
}
