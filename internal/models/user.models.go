package models

import "time"

type User struct {
	ID         string    `db:"id" json:"id" valid:"-"`
	Email      string    `db:"email" json:"email" valid:"email"`
	Phone      string    `db:"phone" json:"phone" valid:"-"`
	Password   string    `db:"password" json:"password" valid:"stringlength(6|100)~password minimum 6 characters"`
	Is_deleted *bool     `db:"is_deleted" json:"is_deleted" valid:"-"`
	CreatedAt  time.Time `db:"created_at" json:"created_at" valid:"-"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at" valid:"-"`
}

type UserDetail struct {
	ID         string    `db:"id" json:"id" valid:"-"`
	Email     string    `db:"email" json:"email" valid:"email"`
	Phone     string    `db:"phone" json:"phone" valid:"_"`
}

type UserLogin struct {
	ID         string    `db:"id" json:"id" valid:"-"`
	Email     string    `db:"email" json:"email" valid:"email"`
	Password   string    `db:"password" json:"password" valid:"stringlength(6|100)~password minimum 6 characters"`
	Phone     string    `db:"phone" json:"phone" valid:"_"`
}

type Users []User
type UserDetails []UserDetail
