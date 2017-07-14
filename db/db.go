package db

import (
	"errors"

	account "github.com/dakstudios/auth-srv/proto/account"
)

type DB interface {
	Init() error
	Account
	Auth
}

type Account interface {
	Read(id string) (*account.User, error)
	Create(user *account.User) error
	Update(user *account.User) error
	Delete(id string) error
	Search(limit, offset int64) ([]*account.User, error)
}

type Auth interface {
	Authorize(id, permission string) (bool, error)
}

var (
	db DB

	ErrNotFound = errors.New("not found")
)

func Regiter(backend DB) {
	db = backend
}

func Init() error {
	return db.Init()
}

// account
func Read(id string) (*account.User, error) {
	return db.Read(id)
}

func Create(user *account.User) error {
	return db.Create(user)
}

func Update(user *account.User) error {
	return db.Update(user)
}

func Delete(id string) error {
	return db.Delete(id)
}

func Search(limit, offset int64) ([]*account.User, error) {
	return db.Search(limit, offset)
}

// auth
func Authorize(id, permission string) (bool, error) {
	return db.Authorize(id, permission)
}
