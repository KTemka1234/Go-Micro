package repository

import (
	"auth/internal/model"
	repomodel "auth/internal/repository/model"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	db     *gorm.DB
	logger *zerolog.Logger
}

func NewRepository(db *gorm.DB, logger *zerolog.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

func (r *Repository) CreateUser(ctx context.Context, user model.User) error {
	userRepo := repomodel.UserToRepoUser(user)
	res := r.db.WithContext(ctx).Create(&userRepo)
	if res.Error != nil {
		r.logger.Err(res.Error).Msg("failed to save user")
		return fmt.Errorf("failed to save user: %w", res.Error)
	}
	return nil
}

func (r *Repository) GetUserByID(ctx context.Context, userID uint64) (model.User, error) {
	var user repomodel.User
	res := r.db.WithContext(ctx).
		Model(&repomodel.User{}).
		Where("id = ?", userID).
		First(&user)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return model.User{}, fmt.Errorf("user not found")
	} else if res.Error != nil {
		r.logger.Err(res.Error).Msg("failed to get user id")
		return model.User{}, res.Error
	}
	return repomodel.RepoUserToUser(user), nil
}

func (r *Repository) GetUserByLoginOrEmail(ctx context.Context, loginOrEmail string) (model.User, error) {
	var user repomodel.User
	res := r.db.WithContext(ctx).
		Model(&repomodel.User{}).
		Where("login = ? OR email = ?", loginOrEmail, loginOrEmail).
		First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return model.User{}, fmt.Errorf("user not found")
	} else if res.Error != nil {
		r.logger.Err(res.Error).Msg("failed to get user by login/email")
		return model.User{}, res.Error
	}
	return repomodel.RepoUserToUser(user), nil
}

func (r *Repository) UpdateRefreshToken(ctx context.Context, token model.RefreshToken) error {
	m := repomodel.RefreshTokenToRepoRefresh(token)
	res := r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}}, // по какому уникальному индексу конфликт
			DoUpdates: clause.AssignmentColumns([]string{"token", "expires_at", "revoked_at"}),
		}).
		Create(&m)
	if res.Error != nil {
		r.logger.Err(res.Error).Msg("failed to save refresh token")
		return fmt.Errorf("failed to save refresh token: %w", res.Error)
	}
	return nil
}

func (r *Repository) GetRefreshToken(ctx context.Context, token string) (model.RefreshToken, error) {
	var refreshToken repomodel.RefreshToken
	res := r.db.WithContext(ctx).
		Model(&repomodel.RefreshToken{}).
		Where("token = ?", token).
		First(&refreshToken)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return model.RefreshToken{}, fmt.Errorf("refresh token not found")
	} else if res.Error != nil {
		r.logger.Err(res.Error).Msg("failed to get refresh token")
		return model.RefreshToken{}, res.Error
	}
	return repomodel.RepoRefreshTokenToRefresh(refreshToken), nil
}

func (r *Repository) RevokeRefreshToken(ctx context.Context, token string) error {
	now := time.Now()
	res := r.db.WithContext(ctx).
		Model(&repomodel.RefreshToken{}).
		Where("token = ?", token).
		Update("revoked_at", now)
	if res.Error != nil {
		r.logger.Err(res.Error).Msg("failed to revoke refresh token")
		return fmt.Errorf("failed to revoke refresh token: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("refresh token not found")
	}
	return nil
}