package main

import (
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"

	"github.com/ThanawitGZS/Sa-Receipt/entity"
)


func main() {
	db, err := gorm.Open(sqlite.Open("Receipt.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
  db.AutoMigrate(

	&entity.Coupon{},

	&entity.Employee{},

	&entity.Gender{},

	&entity.Member{},

	&entity.Position{},

	&entity.Price{},

	&entity.Rank{},

	&entity.Receipt{},
)
	
}