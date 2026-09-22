package model

import "auth/internal/model"

func UserToRepoUser(user model.User) User {
	return User{
		ID:           user.ID,
		Login:        user.Login,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func RepoUserToUser(user User) model.User {
	return model.User{
		ID:           user.ID,
		Login:        user.Login,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func RefreshTokenToRepoRefresh(refreshToken model.RefreshToken) RefreshToken {
	return RefreshToken{
		ID:        refreshToken.ID,
		UserID:    refreshToken.UserID,
		Token:     refreshToken.Token,
		ExpiresAt: refreshToken.ExpiresAt,
		RevokedAt: refreshToken.RevokedAt,
		CreatedAt: refreshToken.CreatedAt,
	}
}

func RepoRefreshTokenToRefresh(refreshToken RefreshToken) model.RefreshToken {
	return model.RefreshToken{
		ID:        refreshToken.ID,
		UserID:    refreshToken.UserID,
		Token:     refreshToken.Token,
		ExpiresAt: refreshToken.ExpiresAt,
		RevokedAt: refreshToken.RevokedAt,
		CreatedAt: refreshToken.CreatedAt,
	}
}
