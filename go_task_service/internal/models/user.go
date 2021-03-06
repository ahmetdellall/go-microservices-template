package models

import uuid "github.com/satori/go.uuid"

type User struct {
	ID            uuid.UUID `gorm:"Column:id;type:uuid" sql:"type:binary(16);not null"`
	FirstName     string    `gorm:"Column:first_name" sql:"type:varchar(255);not null"`
	LastName      string    `gorm:"Column:last_name" sql:"type:varchar(255);default:null"`
	DocumentNotes string    `gorm:"Column:document_notes" sql:"type:varchar(255);default:null"`
}
