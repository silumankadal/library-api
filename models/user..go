package models

import "time"

type User struct {
	ID	uint	`gorm:"primaryKey" json:"id"`
	CreatedAt	*time.time	`json:"created_at,omitempty"`
	DpdatedAt	*time.time	`json:"updated_at,omitempty"`
}