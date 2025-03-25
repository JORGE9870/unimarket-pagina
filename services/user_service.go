package services

import (
	"errors"

	"github.com/beego/beego/v2/client/orm"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	orm orm.Ormer
}

func NewUserService() *UserService {
	return &UserService{
		orm: orm.NewOrm(),
	}
}

func (s *UserService) Create(user *User) error {
	// Encriptar contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Validar email único
	exist := s.orm.QueryTable("usuarios").Filter("email", user.Email).Exist()
	if exist {
		return errors.New("email ya registrado")
	}

	// Crear usuario
	_, err = s.orm.Insert(user)
	return err
}

func (s *UserService) Authenticate(email, password string) (*User, error) {
	var user User
	err := s.orm.QueryTable("usuarios").Filter("email", email).One(&user)
	if err != nil {
		return nil, errors.New("credenciales inválidas")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("credenciales inválidas")
	}

	user.Password = "" // No devolver la contraseña
	return &user, nil
}
