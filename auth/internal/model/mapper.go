package model

import authpb "github.com/KTemka1234/go-micro/contracts/auth/go"

func PbRegisterToRegisterModel(user *authpb.RegisterRequest) Register {
	return Register{
		Login:    user.Login,
		Email:    user.Email,
		Password: user.Password,
	}
}

func PbLoginToLoginModel(user *authpb.LoginRequest) Login {
	return Login{
		LoginOrEmail: user.LoginOrEmail,
		Password:     user.Password,
	}
}
