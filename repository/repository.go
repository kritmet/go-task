package repository

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Repository common repository
type Repository struct{}

// NewRepository new repository
func NewRepository() Repository {
	return Repository{}
}

// Create create
func (r *Repository) Create(db *gorm.DB, i interface{}) error {
	return db.Omit(clause.Associations).Create(i).Error
}

// Update update
func (r *Repository) Update(db *gorm.DB, i interface{}) error {
	return db.Omit(clause.Associations).Save(i).Error
}

// UpdateFieldByID update field by id
func (r *Repository) UpdateFieldByID(db *gorm.DB, id uint, field string, value, i interface{}) error {
	return db.Model(i).Where("id = ?", id).Update(field, value).Error
}

// Delete delete
func (r *Repository) Delete(db *gorm.DB, i interface{}) error {
	return db.Delete(i).Error
}

// FindOneByField find one by field
func (r *Repository) FindOneByField(db *gorm.DB, field string, value, i interface{}) error {
	return db.Where(fmt.Sprintf("%s = ?", field), value).First(i).Error
}
