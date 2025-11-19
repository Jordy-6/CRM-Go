package storage

import (
	"fmt"

	"gorm.io/gorm"
)

type GormStore struct {
	db *gorm.DB
}

func NewGormStore(db *gorm.DB) *GormStore {
	return &GormStore{db: db}
}

func (g *GormStore) Add(contact *Contact) error {
	return g.db.Create(contact).Error
}

func (g *GormStore) Update(id int, contact *Contact) error {
	return g.db.Model(&Contact{}).Where("id = ?", id).Updates(contact).Error
}

func (g *GormStore) Delete(id int) error {
	return g.db.Delete(&Contact{}, id).Error
}

func (g *GormStore) GetAll() error {
	var contacts []Contact
	result := g.db.Find(&contacts)
	if result.Error != nil {
		return result.Error
	}

	if len(contacts) == 0 {
		return fmt.Errorf("no contacts found")
	}

	fmt.Println("📋 Contacts:")
	for _, c := range contacts {
		fmt.Printf("[%d] %s - %s\n", c.ID, c.Name, c.Email)
	}

	return nil
}
