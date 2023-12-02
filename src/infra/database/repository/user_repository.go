package repository

import (
	"github.com/gcors88/read_files_go/src/infra/database"
	"github.com/gcors88/read_files_go/src/infra/database/entities"
)

func CreateUser(u entities.User) {
	database.GetConnection().Create(u)
}

func FindUserByName(Name string) entities.User {
	var user entities.User
	if err := database.GetConnection().Where("name = ?", Name).First(&user).Error; err != nil {
		return user
	}

	return user
}

func UpdateUser(user entities.User) {
	database.GetConnection().Where("name = ?", user.Name).Save(&user)
}
