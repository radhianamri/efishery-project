package model

import (
	"time"
)

// User struct represents users table
type User struct {
	ID           uint64     `gorm:"column:id; type:int(11) unsigned auto_increment; primary_key; not null" json:"id,omitempty"`
	Name         string     `gorm:"column:name; type:varchar(50); not null" json:"name,omitempty"`
	Phone        string     `gorm:"column:phone; type:varchar(20); not null" json:"phone,omitempty"`
	Role         string     `gorm:"column:role; type:varchar(10); not null" json:"role,omitempty"`
	Password     string     `gorm:"column:password; type:varchar(4); not null" json:"password,omitempty"`
	CreatedAt    *time.Time `gorm:"column:created_at; type:datetime; not null; default:current_timestamp" json:"created_at,omitempty"`
	LastModified *time.Time `gorm:"column:last_modified; type:datetime; default:null on update current_timestamp" json:"last_modified,omitempty"`
}

// TableName method for User struct
func (User) TableName() string {
	return "efishery.users"
}

// CreateUser attempts to insert user data
func CreateUser(u *User) (err error) {

	return nil
}

// CheckLoginUser attempts to check user login and return data if found
func CheckLoginUser(phone string, password string) (err error) {

}
