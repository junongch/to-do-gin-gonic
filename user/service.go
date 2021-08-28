package user

import (
	"github.com/junongch/to-do-gin-gonic/config"
)

func GetAllUsersSer(users *[]User) (err error) {
	if err = config.DB.Find(users).Error; err != nil {
		return err
	}

	return nil
}

func GetUserSer(user *User, id string)(err error) {
	if err = config.DB.First(user, id).Error; err != nil {
		return err
	}

	return nil
}

func CreateUserSer(user *User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func UpdateUserSer(user *User) (err error) {
	if err = config.DB.Model(user).Updates(*user).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUserSer(id string) (err error) {
	if err = config.DB.Delete(&User{}, id).Error; err != nil {
		return err
	}

	return nil
}