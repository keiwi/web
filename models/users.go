package models

import (
	"github.com/jinzhu/gorm"

	// import mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	gorm.Model `json:"-"`
	Username   string `gorm:"not null;unique" json:"username"`
	Email      string `gorm:"not null;unique" json:"email"`
	Password   string `gorm:"not null" json:"-"`
	UUID       string `gorm:"not null;unique" json:"uuid"`
}

// UserManager struct
type UserManager struct {
	db *DB
}

// NewUserManager - Creates a new *UserManager that can be used for managing users.
func NewUserManager(db *DB) (*UserManager, error) {
	db.AutoMigrate(&User{})
	manager := UserManager{}
	manager.db = db
	return &manager, nil
}

// HasUser - Check if given username exists.
func (state *UserManager) HasUser(username string) bool {
	if err := state.db.Where("username=?", username).Find(&User{}).Error; err != nil {
		return false
	}
	return true
}

// HasUserByEmail - Check if given email exists.
func (state *UserManager) HasUserByEmail(email string) bool {
	if err := state.db.Where("email=?", email).Find(&User{}).Error; err != nil {
		return false
	}
	return true
}

// FindUser - Tries to find an existing user based on their username
func (state *UserManager) FindUser(username string) *User {
	user := User{}
	state.db.Where("username=?", username).Find(&user)
	return &user
}

// FindUserByEmail - Tries to find an existing user based on their email
func (state *UserManager) FindUserByEmail(email string) *User {
	user := User{}
	state.db.Where("email=?", email).Find(&user)
	return &user
}

// FindUserByUUID - Tries to find an existing user based on their UUID
func (state *UserManager) FindUserByUUID(uuid string) *User {
	user := User{}
	state.db.Where("uuid=?", uuid).Find(&user)
	return &user
}

// AddUser - Creates a user and hashes the password
func (state *UserManager) AddUser(username, email, password string) *User {
	passwordHash := state.HashPassword(username, password)
	user := &User{
		Username: username,
		Email:    email,
		Password: passwordHash,
		UUID:     uuid.NewV4().String(),
	}
	state.db.Create(&user)
	return user
}

// HashPassword - Hash the password (takes a username as well, it can be used for salting).
func (state *UserManager) HashPassword(username, password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Permissions: bcrypt password hashing unsuccessful")
	}
	return string(hash)
}

// CheckPassword - compare a hashed password with a possible plaintext equivalent
func (state *UserManager) CheckPassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
