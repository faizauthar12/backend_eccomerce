package models

import (
	"github.com/faizauthar12/backend_eccomerce/global-utils/model"
)

type Product struct {
	UUID      string `json:"uuid,omitempty" bson:"uuid,omitempty"`
	Name      string `json:"name,omitempty" bson:"name,omitempty"  binding:"required"`
	Price     int64  `json:"price,omitempty" bson:"price,omitempty"`
	Stock     int64  `json:"stock" bson:"stock,omitempty" binding:"required"`
	CreatedAt int64  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type ProductChan struct {
	Product    *Product        `json:"product,omitempty"`
	Total      int64           `json:"total,omitempty"`
	Error      error           `json:"error,omitempty"`
	ErrorLog   *model.ErrorLog `json:"error_log,omitempty"`
	UUID       string          `json:"uuid,omitempty"`
	StatusCode int             `json:"status_code,omitempty"`
}

type ProductsChan struct {
	Products   []*Product      `json:"products,omitempty"`
	Total      int64           `json:"total,omitempty"`
	Error      error           `json:"error,omitempty"`
	ErrorLog   *model.ErrorLog `json:"error_log,omitempty"`
	UUID       string          `json:"uuid,omitempty"`
	StatusCode int             `json:"status_code,omitempty"`
}

type ProductRequest struct {
	NumItems int64  `json:"num_items,omitempty" bson:"num_items,omitempty"`
	Pages    int64  `json:"pages,omitempty" bson:"pages,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	PriceGte int64  `json:"price_minimum,omitempty" bson:"price_minimum,omitempty"`
	PriceLte int64  `json:"price_maximum,omitempty" bson:"price_maximum,omitempty"`
}

type ProductUpdateRequest struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Price int64  `json:"price,omitempty" bson:"price,omitempty"`
	Stock int64  `json:"stock,omitempty" bson:"stock,omitempty"`
}

type Products struct {
	Products []*Product `json:"products,omitempty"`
	Total    int64      `json:"total,omitempty"`
}
