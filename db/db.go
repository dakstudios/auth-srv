package db

import (
	"errors"

	account "github.com/dakstudios/auth-srv/proto/account"
)

type DB interface {
	Init() error
	Account
	Role
	Auth
}

type Account interface {
	FindUser(email string) (*account.User, error)
	ReadUser(id string) (*account.User, error)
	CreateUser(user *account.User) error
	UpdateUser(user *account.User) error
	DeleteUser(id string) error
	SearchUsers(limit, offset int64) ([]*account.User, error)
}

type Role interface {
	ReadRole(id string) (*account.Role, error)
	ReadAllRoles() ([]*account.Role, error)
	CreateRole(role *account.Role) error
	UpdateRole(role *account.Role) error
	DeleteRole(id string) error
}

type Auth interface {
	Authorize(id, permission string) (bool, error)
}

var (
	db DB

	ErrNotFound = errors.New("not found")
)

func Register(backend DB) {
	db = backend
}

func Init() error {
	return db.Init()
}

// account
func FindUser(email string) (*account.User, error) {
	return db.FindUser(email)
}

func ReadUser(id string) (*account.User, error) {
	return db.ReadUser(id)
}

func CreateUser(user *account.User) error {
	return db.CreateUser(user)
}

func UpdateUser(user *account.User) error {
	return db.UpdateUser(user)
}

func DeleteUser(id string) error {
	return db.DeleteUser(id)
}

func SearchUsers(limit, offset int64) ([]*account.User, error) {
	return db.SearchUsers(limit, offset)
}

// role
func ReadRole(id string) (*account.Role, error) {
	return db.ReadRole(id)
}

func ReadAllRoles() ([]*account.Role, error) {
	return db.ReadAllRoles()
}

func CreateRole(role *account.Role) error {
	return db.CreateRole(role)
}

func UpdateRole(role *account.Role) error {
	return db.UpdateRole(role)
}

func DeleteRole(id string) error {
	return db.DeleteRole(id)
}

// auth
func Authorize(id, permission string) (bool, error) {
	return db.Authorize(id, permission)
}
