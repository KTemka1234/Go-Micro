package model

import "account/internal/model"

func UserToRepoUser(user model.User) User {
	return User{
		ID:         user.ID,
		Login:      user.Login,
		Email:      user.Email,
		Phone:      user.Phone,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		Age:        user.Age,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}

func RepoUserToUser(user User) model.User {
	res := model.User{
		ID:         user.ID,
		Login:      user.Login,
		Email:      user.Email,
		Phone:      user.Phone,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		Age:        user.Age,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
	return res
}

func RepoUsersToUsers(users []User) []model.User {
	res := make([]model.User, len(users))
	for i, user := range users {
		res[i] = RepoUserToUser(user)
	}
	return res
}

func UpdateUserToRepoUser(user model.UpdateUser) User {
	return User{
		Email:      user.Email,
		Phone:      user.Phone,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		Age:        user.Age,
	}
}

func CreateUserToRepoUser(user model.CreateUser) User {
	return User{
		Login:      user.Login,
		Email:      user.Email,
		Phone:      user.Phone,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		Age:        user.Age,
	}
}
