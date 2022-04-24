package models

type TaskCategory struct {
	ID          *[]byte `gorm:"Column:id" sql:"type:binary(16);not null"`
	CategoryKey *int    `gorm:"Column:category_key" sql:"type:tinyint(4);not null"`
	Name        *string `gorm:"Column:name" sql:"type:varchar(255);default:null"`
}
