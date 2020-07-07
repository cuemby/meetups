package models

import "github.com/netsaj/episodio8-api-accounts/internal/utils"

type User struct {
	Base
	Name  string `gorm:"size:225;not null" json:"name"`
	Email string `gorm:"size:225;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
func (User) TableName() string {
	return "users"
}

func (u User) CheckPassword(password string) bool {
	return utils.CheckPasswordHash(password, u.Password)
}
func (u *User) SetPassword(password string) bool {
	var err error
	u.Password, err = utils.HashPassword(password)
	return err == nil
}
