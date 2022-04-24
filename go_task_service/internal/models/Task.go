package models

import "time"

type Task struct {
	ID               *[]byte    `gorm:"Column:id" sql:"type:binary(16);not null"`
	Title            *string    `gorm:"Column:title" sql:"type:varchar(255);not null"`
	TaskKey          *string    `gorm:"Column:task_key" sql:"type:varchar(255);default:null"`
	Details          *string    `gorm:"Column:details" sql:"type:varchar(255);default:null"`
	ExpectedDateTime *time.Time `gorm:"Column:expected_date_time"sql:"type:timestamp;default:current_timestamp"`
	Status           *int       `gorm:"Column:status" sql:"type:tinyint(4);not null"`
	CreatorID        *[]byte    `gorm:"Column:creator_id" sql:"type:binary(16);not null"`
	CreatedDate      *time.Time `gorm:"Column:created_date" sql:"type:timestamp;default:current_timestamp"`
	UpdatedDate      *time.Time `gorm:"Column:created_date" sql:"type:timestamp;default:current_timestamp"`
	StartDateTime    *time.Time `gorm:"Column:created_date" sql:"type:timestamp;default:0000-00-00 00:00:00"`
	EndDateTime      *time.Time `gorm:"Column:created_date" sql:"type:timestamp;default:0000-00-00 00:00:00"`
	CategoryID       *[]byte    `gorm:"Column:category_id" sql:"type:binary(16);default:null"`
}
