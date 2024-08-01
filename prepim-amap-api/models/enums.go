package models

type Status struct {
	Code string `sql:"type:ENUM('ACTIVE', 'INACTIVE', 'CANCELLED')" gorm:"primaryKey:column:code"`
	Name string `gorm:"column:name"`
}
