package storage

import "gorm.io/gorm"

type Storer interface {
	Add(contact *Contact) error
	Update(id int, contact *Contact) error
	Delete(id int) error
	GetAll() error
}

type Contact struct {
	gorm.Model
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"type:varchar(100)"`
	Email string `json:"email" gorm:"type:varchar(100)"`
}
