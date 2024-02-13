package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `gorm:"type:string;size:20;not null,unique"`
}

type User struct {
	gorm.Model
	FirstName string `gorm:"type:string;size:20;not null"`
	LastName  string `gorm:"type:string;size:30;null"`
	Username  string `gorm:"type:string;size:20;not null;unique"`
	Email     string `gorm:"type:string;size:64;not null;unique;"`
	Password  string `gorm:"type:string;size:64;not null" json:"-"`
	RoleID    int
	Role      Role
}
