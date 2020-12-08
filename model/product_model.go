package model

import (
	"database/sql"

	"github.com/tjgurwara99/golang_database_utility/entity"
)

// ProductModel Database model for Product entity
type ProductModel struct {
	DB *sql.DB
}

// CreateProduct Creates a new Product record in the database
func (productModel *ProductModel) CreateProduct(newProduct *entity.Product) {

}
