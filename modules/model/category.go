package model

import (
	"time"
)

// Định nghĩa struct model Category
type Category struct {
	Categoryid     int            `json:"categoryid"  gorm:"column:categoryid;primaryKey"`
	Categoryname   string         `json:"categoryname" gorm:"column:categoryname;" binding:"required,alpha"`
	Created_at     *time.Time     `json:"created_at" gorm:"created_at;"`
	Update_at      *time.Time     `json:"update_at" gorm:"update_at;"`
	Categorystatus *ProductStatus `json:"categorystatus" gorm:"categorystatus;"`
}

// Cho GORM biết mapping struct này với table categories trong database.
func (Category) CategoryTableName() string {
	return "categories"
}
