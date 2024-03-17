package models

import "gorm.io/gorm"

type InvalidToken struct {
	gorm.Model
	Username string
	Token    string
}
