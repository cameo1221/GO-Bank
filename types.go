package main

import (
	"time"

	"github.com/google/uuid"
)

type Admin struct {
	Id         uuid.UUID  `json:"id,omitempty" db:"id"`
	Name       string     `json:"name,omitempty" db:"name"`
	Email      string     `json:"email,omitempty" db:"email"`
	Password   string     `json:"password,omitempty" db:"password"`
	CreatedAt  time.Time  `json:"createdAt,omitempty" db:"created_at"`
	ArchivedAt *time.Time `json:"archivedAt,omitempty" db:"archived_at"`
}

func NewAdmin(name, email, password string) *Admin {
	return &Admin{

		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}
}

type Admin_session struct {
	Id         uuid.UUID  `json:"id,omitempty" db:"id"`
	AdminId    uuid.UUID  `json:"Adminid,omitempty" db:"Admin_id"`
	CreatedAt  time.Time  `json:"createdAt,omitempty" db:"created_at"`
	ArchivedAt *time.Time `json:"archivedAt,omitempty" db:"archived_at"`
}
type Asset struct {
	Id         uuid.UUID  `json:"id,omitempty" db:"id"`
	Model      string     `json:"Model,omitempty" db:"Model"`
	Company    string     `json:"company,omitempty" db:"company"`
	CreatedAt  time.Time  `json:"createdAt,omitempty" db:"created_at"`
	ArchivedAt *time.Time `json:"archivedAt,omitempty" db:"archived_at"`
}
type Emplyoee struct {
	Id         uuid.UUID  `json:"id,omitempty" db:"id"`
	Name       string     `json:"name,omitempty" db:"name"`
	Email      string     `json:"email,omitempty" db:"email"`
	Role       string     `json:"role,omitempty" db:"role"`
	CreatedAt  time.Time  `json:"createdAt,omitempty" db:"created_at"`
	ArchivedAt *time.Time `json:"archivedAt,omitempty" db:"archived_at"`
}
type EmployeeAssetMapping struct {
	Id         uuid.UUID  `json:"id,omitempty" db:"id"`
	AssetId    string     `json:"assetid,omitempty" db:"assetid"`
	EmplyoeeID string     `json:"emplyoeeid,omitempty" db:"emplyoeeid"`
	CreatedAt  time.Time  `json:"createdAt,omitempty" db:"created_at"`
	ArchivedAt *time.Time `json:"archivedAt,omitempty" db:"archived_at"`
}
