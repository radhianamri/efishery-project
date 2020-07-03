package model

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/radhianamri/efishery-project/auth-go/config"
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

// User struct represents users table
type UserRegister struct {
	Name  string `gorm:"column:name; type:varchar(50); not null" json:"name,omitempty"`
	Phone string `gorm:"column:phone; type:varchar(20); not null" json:"phone,omitempty"`
	Role  string `gorm:"column:role; type:varchar(10); not null" json:"role,omitempty"`
}

// UserLogin struct represents login requirements
type UserLogin struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// TableName method for User struct
func (User) TableName() string {
	return "efishery.users"
}

// ValidateUser validates user information
func ValidateUser(u *User) (err error) {
	//validates name field
	u.Name = strings.TrimSpace(u.Name)
	if u.Name == "" {
		return fmt.Errorf("Missing name field")
	}

	//validates phone field
	u.Phone = strings.TrimSpace(u.Phone)
	if u.Phone == "" {
		return fmt.Errorf("Missing phone field")
	}

	if u.Phone[0] != '+' && u.Phone[0] != '0' {
		return fmt.Errorf("Invalid phone: %s. Needs to begin with '+' or '0' character", u.Phone)
	}
	if u.Phone[0] == '0' {
		u.Phone = "+62" + u.Phone[1:]
	}
	u.Phone = config.SimplifyPhoneNumber(u.Phone)
	if !regexp.MustCompile(config.RegexPhoneNumber).MatchString(u.Phone) {
		return fmt.Errorf("Invalid phone format: %s", u.Phone)
	}

	//validates role field
	u.Role = strings.TrimSpace(u.Role)
	if u.Role == "" {
		return fmt.Errorf("Invalid role: %s", u.Role)
	}

	//creates random 4 character password
	b := make([]rune, 4)
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	u.Password = string(b)
	return nil
}

// CreateUser attempts to insert user data
func CreateUser(u *User) (err error) {
	if err := ValidateUser(u); err != nil {
		return err
	}
	if err = config.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

// CheckLoginUser attempts to check user login and return data if found
func CheckLoginUser(u *User) (err error) {
	if err = config.DB.Where("phone = ? AND password = ?", u.Phone, u.Password).First(u).Error; err != nil {
		return err
	}
	return nil
}
