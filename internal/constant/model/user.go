package model

import (
	uuid "github.com/satori/go.uuid"
	"time"

	// "mime/multipart"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	"gorm.io/gorm"
)

//CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
type User struct {
	ID          uuid.UUID `json:"id,omitempty" gorm:"primaryKey;type:uuid;default:uuid_generate_v4();"`
	Name        string    `json:"name,omitempty" validate:"required"`
	UserName    string    `json:"username,omitempty" gorm:"unique;not null" validate:"required"`
	Password    string    `json:"password,omitempty" gorm:"not null" validate:"required,passwordcheck"`
	Address     string    `json:"address,omitempty" `
	Description string    `json:"description,omitempty" gorm:"default:pending"  validate:"required"`
	RoleName    string    `json:"role_name,omitempty"`
	Dob         time.Time `json:"dob,omitempty" gorm:"default:current_timestamp"  validate:"required"`
	Role        *Role     `json:"role,omitempty" gorm:"foreignKey:RoleName;references:Name" `
	CreatedAt   time.Time `json:"created_at,omitempty" gorm:"default:current_timestamp"`
}
type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (u *User) Sanitize() {
	u.Password = ""
}

type LoginCredential struct {
	UserName string `json:"username,omitempty"  validate:"required"`
	Password string `json:"password,omitempty"  validate:"required,passwordcheck"`
}
type Utils struct {
	Conn        *gorm.DB
	GoValidator *validator.Validate
	Translator  ut.Translator
	Timeout     time.Duration
}
type DeleteData struct {
	Result string `json:"result,omitempty"`
}

func (p *User) TableName() string {
	return "users"
}
