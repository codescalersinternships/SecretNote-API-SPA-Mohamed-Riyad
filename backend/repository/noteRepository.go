package repository

import (
	"github.com/codescalersinternships/SecretNote-API-SPA-Mohamed-Riyad/models"
	"gorm.io/gorm"
)

type NoteRepository interface {
	Create(note *models.Note) error
	GetByID(id uint) (*models.Note, error)
	GetAllByUserID(userID uint) ([]*models.Note, error)
	Update(note *models.Note) error
	Delete(id uint) error
}

type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &noteRepository{db}
}

func (r *noteRepository) Create(note *models.Note) error {
	return r.db.Create(note).Error
}

func (r *noteRepository) GetByID(id uint) (*models.Note, error) {
	var note models.Note
	err := r.db.First(&note, id).Error
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func (r *noteRepository) GetAllByUserID(userID uint) ([]*models.Note, error) {
	var notes []*models.Note
	err := r.db.Where("user_id = ?", userID).Find(&notes).Error
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (r *noteRepository) Update(note *models.Note) error {
	return r.db.Save(note).Error
}

func (r *noteRepository) Delete(id uint) error {
	return r.db.Delete(&models.Note{}, id).Error
}
