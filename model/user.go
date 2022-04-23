package model

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateUser(firstName, lastName string) (user User, err error) {
	user = User{
		FirstName: firstName,
		LastName:  lastName,
	}
	err = db.Create(&user).Error
	return
}

func FindUserById(id int) (user User, err error) {
	err = db.Find(&user, id).Error
	return
}

func DeleteUser(id int) (err error) {
	err = db.Delete(&User{}, id).Error
	return
}

func UpdateUser(id int, firstName, lastName string) (user User, err error) {
	user = User{
		FirstName: firstName,
		LastName:  lastName,
	}

	err = db.Model(&user).Where("id = ?", id).Updates(user).First(&user).Error
	return
}
