package models

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	Comment string `gorm:"type:string;size:512;not null" json:"comment"`
	UserID  int
	User    User
}
