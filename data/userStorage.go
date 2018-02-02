package data

import (
	"time"

	"github.com/TheoRev/GolangMongoDB/models/user"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserStorage struct {
	c       *mgo.Collection
	context *Context
}

func (us *UserStorage) setContext() {
	us.context = NewContext()
	us.c = us.context.DBCollection(CName)
}

// Create crea un usuario
func (us UserStorage) Create(u *user.User) error {
	u.ID = bson.NewObjectId()
	hpass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		u = nil
		return err
	}
	u.HashPassword = hpass
	u.Password = ""
	u.CreateAt = time.Now()

	us.setContext()
	defer us.context.Close()
	err = us.c.Insert(u)
	u.HashPassword = []byte{}
	return err
}

// Update actualiza un registro de usuario
func (us UserStorage) Update(u *user.User) error {
	us.setContext()
	u.UpdateAt = time.Now()

	err := us.c.Update(
		bson.M{"_id": u.ID},
		bson.M{"$set": bson.M{
			"username": u.UserName,
			"email":    u.Email,
			"updateat": time.Now(),
		}},
	)
	defer us.context.Close()
	return err
}

// Delete elimina un usuario
func (us UserStorage) Delete(u *user.User) error {
	us.setContext()
	err := us.c.RemoveId(u.ID)
	defer us.context.Close()
	return err
}

// GetAll obtiene todos los usuarios
func (us UserStorage) GetAll() (user.Users, error) {
	var users user.Users
	us.setContext()
	// Acendente
	iter := us.c.Find(nil).Sort("username").Iter()
	// Descendenete
	// iter := us.c.Find(nil).Sort("-username").Iter()
	u := user.New()
	for iter.Next(u) {
		u.HashPassword = []byte{}
		users = append(users, *u)
	}
	defer us.context.Close()
	return users, nil
}

// GetByID obtiene un usuario segun el Id
func (us UserStorage) GetByID(u *user.User) error {
	us.setContext()
	err := us.c.FindId(u.ID).One(u)
	u.HashPassword = []byte{}
	defer us.context.Close()
	return err
}
