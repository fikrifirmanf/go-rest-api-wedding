package model

import (
	"time"

	util "github.com/fikrifirmanf/go-rest-api-wedding/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"type:varchar(255);unique;not null" json:"username"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	Role      string    `gorm:"type:varchar(255);not null" json:"role"`
	IsActive  bool      `gorm:"type:tinyint(1);not null;default:1" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"-" json:"-"`
}

func (entity *Users) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.Password = util.HashPassword(entity.Password)
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Users) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
