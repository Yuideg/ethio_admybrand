package model

import (
	"time"
)

type Role struct {
	Name      string    `json:"name,omitempty" gorm:"primaryKey=true;not null;unique"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"default:current_timestamp"`
}

func (p *Role) TableName() string {
	return "roles"
}
