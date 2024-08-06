package entity

import (

	"gorm.io/gorm"

)

type Price struct {
	gorm.Model

	Name 	string	`json:"name"`

	Price 	uint	`json:"price"`
}