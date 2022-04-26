package model

import "gorm.io/gorm"

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{db}
}

func (m *UserModel) CreateUser(firstName, lastName string) (user User, err error) {
	user = User{
		FirstName: firstName,
		LastName:  lastName,
	}
	err = m.db.Create(&user).Error
	return
}

func (m *UserModel) FindUserById(id int) (user User, err error) {
	err = m.db.Find(&user, id).Error
	return
}

func (m *UserModel) DeleteUser(id int) (err error) {
	err = m.db.Delete(&User{}, id).Error
	return
}

func (m *UserModel) UpdateUser(id int, firstName, lastName string) (user User, err error) {
	user = User{
		FirstName: firstName,
		LastName:  lastName,
	}

	err = m.db.Model(&user).Where("id = ?", id).Updates(user).First(&user).Error
	return
}
