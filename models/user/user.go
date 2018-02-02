package user

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Storage metodos del crud
type Storage interface {
	Create(user *User) error
	Update(user *User) error
	Delete(user *User) error
	GetAll() (Users, error)
	GetByID(user *User) error
}

var storage Storage

// SetStorage establece el storage
func SetStorage(s Storage) {
	storage = s
}

// User structura del modelo
type User struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	UserName     string
	Email        string
	Password     string
	HashPassword []byte
	CreateAt     time.Time
	UpdateAt     time.Time
	DeleteAt     time.Time
}

// Create crea un usuario
func (u *User) Create() error {
	return storage.Create(u)
}

// Update actualiza un registro de usuario
func (u *User) Update() error {
	return storage.Update(u)
}

// Delete elimina un usuario
func (u *User) Delete() error {
	return storage.Delete(u)
}

// GetAll obtiene todos los usuarios
func (u *User) GetAll() (Users, error) {
	return storage.GetAll()
}

// GetByID obtiene un usuario segun el Id
func (u *User) GetByID() error {
	return storage.GetByID(u)
}

// Users slice de usuarios
type Users []User

// New crea una nueva instancia de usuario
func New() *User {
	return &User{}
}
