package entity

import (
        "gorm.io/gorm"

        "time"
)

type Gender struct {

	gorm.Model

	GenderName string `json:"gender_name"`

    CreatedAt time.Time

    UpdatedAt time.Time

    DeletedAt gorm.DeletedAt `gorm:"index"`

	// Gender 0..* to 1 Employee
	Employee []Employee `gorm:"foreignKey:GenderID"`
}