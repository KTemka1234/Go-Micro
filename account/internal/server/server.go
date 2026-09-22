package server

import (
	"account/internal/model"
	"context"

	accountpb "github.com/KTemka1234/go-micro/contracts/account/go"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AccountService interface {
	CreateUser(context.Context, model.CreateUser) (model.User, error)
	GetUser(ctx context.Context, userId uint64) (model.User, error)
	GetUsers(ctx context.Context, limit int, offset int) ([]model.User, error)
	DeleteUser(ctx context.Context, userId uint64) error
	UpdateUser(ctx context.Context, userId uint64, user model.UpdateUser) error
}

type Server struct {
	accountpb.UnimplementedAccountServer

	accountService AccountService
	logger         *zerolog.Logger
}

func New(accountService AccountService, logger *zerolog.Logger) *Server {
	return &Server{accountService: accountService, logger: logger}
}

func (h *Server) CreateUser(ctx context.Context, req *accountpb.CreateUserRequest) (*accountpb.CreateUserResponse, error) {
	newUser := model.PbToUserCreate(req.GetUser())
	user, err := h.accountService.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}
	return &accountpb.CreateUserResponse{User: model.UserToPb(user)}, nil
}

func (h *Server) GetUser(ctx context.Context, req *accountpb.GetUserRequest) (*accountpb.GetUserResponse, error) {
	res, err := h.accountService.GetUser(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &accountpb.GetUserResponse{
		User: model.UserToPb(res),
	}, nil
}

func (h *Server) GetUsers(ctx context.Context, req *accountpb.GetUsersRequest) (*accountpb.GetUsersResponse, error) {
	res, err := h.accountService.GetUsers(ctx, int(req.GetPagination().GetLimit()), int(req.GetPagination().GetOffset()))
	if err != nil {
		return nil, err
	}
	return &accountpb.GetUsersResponse{
		Users: model.UsersToPbs(res),
	}, nil
}

func (h *Server) UpdateUser(ctx context.Context, req *accountpb.UpdateUserRequest) (*emptypb.Empty, error) {
	user := model.PbToUserUpdate(req.User)
	if err := h.accountService.UpdateUser(ctx, uint64(req.GetUserId()), user); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h *Server) DeleteUser(ctx context.Context, req *accountpb.DeleteUserRequest) (*emptypb.Empty, error) {
	if err := h.accountService.DeleteUser(ctx, uint64(req.GetUserId())); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}