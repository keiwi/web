package models

import (
	"errors"
	"time"

	"github.com/keiwi/utils"
	"github.com/keiwi/utils/models"
	"github.com/nats-io/go-nats"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"

	"golang.org/x/crypto/bcrypt"
)

// UserManager struct
type UserManager struct {
	conn *nats.Conn
}

// NewUserManager - Creates a new *UserManager that can be used for managing users.
func NewUserManager(conn *nats.Conn) (*UserManager, error) {
	manager := UserManager{}
	manager.conn = conn
	return &manager, nil
}

func (state *UserManager) Has(filter utils.Filter) bool {
	requestData := utils.HasOptions{
		Filter: filter,
	}

	data, err := bson.MarshalJSON(&requestData)
	if err != nil {
		return false
	}

	msg, err := state.conn.Request("users.has", data, time.Duration(viper.GetInt("nats_delay"))*time.Second)
	if err != nil {
		return false
	}

	var has bool
	err = bson.UnmarshalJSON(msg.Data, has)
	if err != nil {
		return false
	}

	return has
}

// HasUser - Check if given username exists.
func (state *UserManager) HasUser(username string) bool {
	return state.Has(utils.Filter{"username": username})
}

// HasUserByEmail - Check if given email exists.
func (state *UserManager) HasUserByEmail(email string) bool {
	return state.Has(utils.Filter{"email": email})
}

// FindUser - Tries to find an existing user based on their username
func (state *UserManager) FindUser(username string) (*models.User, error) {
	requestData := utils.FindOptions{
		Filter: utils.Filter{"username": username},
		Sort:   utils.Sort{"created_at"},
		Limit:  1,
	}
	data, err := bson.MarshalJSON(&requestData)
	if err != nil {
		return nil, err
	}

	msg, err := state.conn.Request("users.retrieve.find", data, time.Duration(viper.GetInt("nats_delay"))*time.Second)
	if err != nil {
		return nil, err
	}

	users := []models.User{}
	err = bson.UnmarshalJSON(msg.Data, &users)
	if err != nil {
		return nil, err
	}
	if len(users) <= 0 {
		return nil, errors.New("could not find any users")
	}
	return &users[0], nil
}

// FindUserByEmail - Tries to find an existing user based on their email
func (state *UserManager) FindUserByEmail(email string) (*models.User, error) {
	requestData := utils.FindOptions{
		Filter: utils.Filter{"email": email},
		Sort:   utils.Sort{"created_at"},
		Limit:  1,
	}
	data, err := bson.MarshalJSON(&requestData)
	if err != nil {
		return nil, err
	}

	msg, err := state.conn.Request("users.retrieve.find", data, time.Duration(viper.GetInt("nats_delay"))*time.Second)
	if err != nil {
		return nil, err
	}

	users := []models.User{}
	err = bson.UnmarshalJSON(msg.Data, &users)
	if err != nil {
		return nil, err
	}
	if len(users) <= 0 {
		return nil, errors.New("could not find any users")
	}
	return &users[0], nil
}

// AddUser - Creates a user and hashes the password
func (state *UserManager) AddUser(username, email, password string) (*models.User, error) {
	passwordHash := state.HashPassword(username, password)
	user := &models.User{
		Username: username,
		Email:    email,
		Password: passwordHash,
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	data, err := bson.MarshalJSON(&user)
	if err != nil {
		return nil, err
	}
	return user, state.conn.Publish("users.create.send", data)
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
