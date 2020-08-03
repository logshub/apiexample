package main

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// User represents user entity
type User struct {
	UserID    string    `sql:"type:uuid;primary_key" json:"user_id"`
	Active    bool      `gorm:"active;not null" json:"active"`
	Name      string    `gorm:"name;not null" json:"name"`
	CreatedAt time.Time `gorm:"not null" json:"-"`
	UpdatedAt time.Time `gorm:"not null" json:"-"`
}

// isValid returns bool whether the user object is valid. ID is checked only if not empty
func (t *User) isValid() bool {
	length := len(t.Name)
	if length < 2 || length > 50 {
		return false
	}
	if t.UserID != "" && !t.isValidID() {
		return false
	}

	return true
}

// isValidID returns bool whether ID is valid (it must be valid UUID)
func (t *User) isValidID() bool {
	if _, err := uuid.Parse(t.UserID); err != nil {
		return false
	}

	return true
}

// add adds a new entry to the database
func (t *User) add(db *gorm.DB) error {
	if dbc := db.Create(t); dbc.Error != nil {
		return dbc.Error
	}

	return nil
}

// load loads the entry from the database
func (t *User) load(db *gorm.DB) error {
	if dbc := db.Where("user_id = ?", t.UserID).First(t); dbc.Error != nil {
		return dbc.Error
	}

	return nil
}

// delete deletes the entry from the database
func (t *User) delete(db *gorm.DB) error {
	if !t.isValidID() {
		return errors.New("No user_id provided")
	}
	if dbc := db.Where("user_id = ?", t.UserID).Delete(t); dbc.Error != nil {
		return dbc.Error
	}

	return nil
}

// BeforeCreate gorm callback
func (t *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("user_id", uuid.New().String())
	return nil
}

// TableName gorm callback
func (User) TableName() string {
	return "users"
}

// findAllUsers returns the slice of User objects
func findAllUsers(db *gorm.DB, filter usersFilter) []User {
	users := []User{}
	if filter.active == -1 {
		db.Find(&users)
	} else if filter.active == 0 {
		db.Where("active = ?", false).Find(&users)
	} else {
		db.Where("active = ?", true).Find(&users)
	}

	return users
}

// usersFilter used for filtering of users
type usersFilter struct {
	active int8
}

// SetActiveFilter sets the active filtering param
func (t *usersFilter) SetActiveFilter(activeParam string) error {
	if activeParam == "" {
		return nil
	} else if activeParam == "1" {
		t.active = 1
		return nil
	} else if activeParam == "0" {
		t.active = 0
		return nil
	}

	return errors.New("filter not allowed")
}

// newUsersFilter returns the new instance of filtering object
func newUsersFilter() usersFilter {
	return usersFilter{
		active: -1,
	}
}
