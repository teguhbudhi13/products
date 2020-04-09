package model

import (
	"github.com/jinzhu/gorm"
	// import mysql driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

// product schema
type Product struct {
	ID    int    `gorm:"primary_key; auto_increment" sql:"id" json:"id"`
	Name  string `gorm:"not null" json:"name"`
	Email string `gorm:"not null" json:"email"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Product{})
	return db
}
