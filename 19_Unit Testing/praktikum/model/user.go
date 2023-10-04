package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserInterface interface {
	CreateUser(data User) *User
	LoginUser(data User) *User
	GetAllUser() []User
	GetUserByID(id uint64) *User
	DeleteUserByID(id uint64) *User
	UpdateUserByID(data User) *User
}

type User struct {
	gorm.Model
	ID       uint64 `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	// Book     []Book `gorm:"foreignKey:Penulis;reference:ID"`
	// Blog     []Blog `gorm:"foreignKey:IDUser;reference:ID"`
}

type LoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserModel struct {
	db *gorm.DB
}

func (um *UserModel) Init(db *gorm.DB) {
	um.db = db
}

func NewUserModel(db *gorm.DB) UserInterface {
	return &UserModel{
		db: db,
	}
}

func (um *UserModel) CreateUser(data User) *User {
	if err := um.db.Create(&data).Error; err != nil {
		logrus.Error("Model : Register User failed")
	}

	return &data
}

func (um *UserModel) LoginUser(data User) *User {
	var dataa = User{}
	if err := um.db.Where("email = ? AND password = ?", data.Email, data.Password).First(&data).Error; err != nil {
		logrus.Error("Login failed", err.Error())
	}

	if data.Name == "" {
		logrus.Error("User not found")
	}

	return &dataa

}

func (um *UserModel) GetAllUser() []User {
	var listUser = []User{}

	if err := um.db.Find(&listUser).Error; err != nil {
		logrus.Error("Model : Failed Get All Users", err.Error())
		return nil
	}
	return listUser
}

func (um *UserModel) GetUserByID(id uint64) *User {
	var user = User{}
	user.ID = id
	if err := um.db.First(&user, id); err != nil {
		logrus.Error("model : failed get user by id", err.Error)
	}

	return &user
}

func (um *UserModel) DeleteUserByID(id uint64) *User {
	var user = User{}
	user.ID = id
	if err := um.db.Delete(&user).Error; err != nil {
		logrus.Error("Error deleting user")
	}

	return &user
}

func (um *UserModel) UpdateUserByID(data User) *User {
	// var user = User{}
	var query = um.db.Table("users").Where("id = ?", data.ID).Update("name", data.Name)
	if err := query.Error; err != nil {
		logrus.Error("Model : Update Error")
	}

	if dataCount := query.RowsAffected; dataCount < 1 {
		logrus.Error("Model : Update Error, no data affected")
	}

	return &data
}
