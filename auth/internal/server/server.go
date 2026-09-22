package server

import (
	"auth/internal/model"
	"context"

	authpb "github.com/KTemka1234/go-micro/contracts/auth/go"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthService interface {
	Register(context.Context, model.Register) error
	Login(context.Context, model.Login) (model.TokenPair, error)
	Refresh(context.Context, string) (model.TokenPair, error)
	Validate(context.Context, string) (uint64, error)
	Logout(context.Context, string) error
}

type Server struct {
	authpb.UnimplementedAuthServer

	authService AuthService
	logger      *zerolog.Logger
}

func New(authService AuthService, logger *zerolog.Logger) *Server {
	return &Server{authService: authService, logger: logger}
}

func (s *Server) Register(ctx context.Context, req *authpb.RegisterRequest) (*emptypb.Empty, error) {
	r := model.PbRegisterToRegisterModel(req)
	if err := s.authService.Register(ctx, r); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	pair, err := s.authService.Login(ctx, model.PbLoginToLoginModel(req))
	if err != nil {
		return nil, err
	}
	return &authpb.LoginResponse{TokenPair: &authpb.TokenPair{AccessToken: pair.AccessToken, RefreshToken: pair.RefreshToken}}, nil
}

func (s *Server) Refresh(ctx context.Context, req *authpb.RefreshRequest) (*authpb.RefreshResponse, error) {
	pair, err := s.authService.Refresh(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}
	return &authpb.RefreshResponse{TokenPair: &authpb.TokenPair{AccessToken: pair.AccessToken, RefreshToken: pair.RefreshToken}}, nil
}

func (s *Server) Validate(ctx context.Context, req *authpb.ValidateRequest) (*authpb.ValidateResponse, error) {
	uid, err := s.authService.Validate(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}
	return &authpb.ValidateResponse{UserId: uid}, nil
}

func (s *Server) Logout(ctx context.Context, req *authpb.LogoutRequest) (*emptypb.Empty, error) {
	if err := s.authService.Logout(ctx, req.RefreshToken); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
