package model

import (
	accountpb "github.com/KTemka1234/go-micro/contracts/account/go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PbToUser(userpb *accountpb.User) User {
	return User{
		ID: userpb.Id,
		Login: userpb.Login,
		Email: userpb.Email,
		Phone: userpb.Phone,
		FirstName: userpb.FirstName,
		LastName: userpb.LastName,
		MiddleName: userpb.MiddleName,
		Age: userpb.Age,
		CreatedAt: userpb.CreatedAt.AsTime(),
		UpdatedAt: userpb.UpdatedAt.AsTime(),
	}
}

func UserToPb(user User) *accountpb.User {
	return &accountpb.User{
		Id: user.ID,
		Login: user.Login,
		Email: user.Email,
		Phone: user.Phone,
		FirstName: user.FirstName,
		LastName: user.LastName,
		MiddleName: user.MiddleName,
		Age: user.Age,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func UsersToPbs(users []User) []*accountpb.User {
	pbs := make([]*accountpb.User, len(users))
	for i, user := range users {
		pbs[i] = UserToPb(user)
	}
	return pbs
}

func PbsToUsers(pbs []*accountpb.User) []User {
	users := make([]User, len(pbs))
	for i, pb := range pbs {
		users[i] = PbToUser(pb)
	}
	return users
}

func PbToUserCreate(accountpbUser *accountpb.CreateUser) CreateUser {
	return CreateUser{
		Login:      accountpbUser.Login,
		Email:      accountpbUser.Email,
		Phone:      accountpbUser.Phone,
		FirstName:  accountpbUser.FirstName,
		LastName:   accountpbUser.LastName,
		MiddleName: accountpbUser.MiddleName,
		Age:        accountpbUser.Age,
	}
}

func PbToUserUpdate(userpb *accountpb.User) UpdateUser {
	return UpdateUser{
		Email:      userpb.Email,
		Phone:      userpb.Phone,
		FirstName:  userpb.FirstName,
		LastName:   userpb.LastName,
		MiddleName: userpb.MiddleName,
		Age:        userpb.Age,
	}
}
