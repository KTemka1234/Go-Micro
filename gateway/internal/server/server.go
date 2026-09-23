package server

import (
	"context"

	"gateway/internal/model"
	"gateway/internal/server/interceptor"

	gatewaypb "github.com/KTemka1234/go-micro/contracts/gateway/go"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/rs/zerolog"
)

type GatewayService interface {
	Register(context.Context, model.User, string) (model.User, model.TokenPair, error)
	Login(context.Context, string, string) (model.User, model.TokenPair, error)
	Logout(context.Context, string) error
	Refresh(context.Context, string) (model.TokenPair, error)
	ValidateToken(context.Context, string) (uint64, bool, error)
	CreateUser(context.Context, model.CreateUser) error
	GetUser(ctx context.Context, userID uint64) (model.User, error)
	GetUsers(ctx context.Context, limit int, offset int) ([]model.User, error)
	DeleteUser(context.Context, uint64) error
	UpdateUser(context.Context, uint64, model.UpdateUser) error
}

type Server struct {
	gatewaypb.UnimplementedGatewayServer

	gatewayService GatewayService
	logger         *zerolog.Logger
}

func New(gatewayService GatewayService, logger *zerolog.Logger) *Server {
	return &Server{gatewayService: gatewayService, logger: logger}
}

func (s *Server) Register(ctx context.Context, req *gatewaypb.RegisterRequest) (*gatewaypb.RegisterResponse, error) {
	user := model.PbToUser(req.GetUser())

	createdUser, tokens, err := s.gatewayService.Register(ctx, user, req.Password)
	if err != nil {
		return nil, err
	}

	return &gatewaypb.RegisterResponse{
		User:   model.UserToPb(createdUser),
		Tokens: model.TokenPairToPb(tokens),
	}, nil
}

func (s *Server) Login(ctx context.Context, req *gatewaypb.LoginRequest) (*gatewaypb.LoginResponse, error) {
	user, tokens, err := s.gatewayService.Login(ctx, req.GetLoginOrEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	return &gatewaypb.LoginResponse{
		User:   model.UserToPb(user),
		Tokens: model.TokenPairToPb(tokens),
	}, nil
}

func (s *Server) Refresh(ctx context.Context, req *gatewaypb.RefreshRequest) (*gatewaypb.RefreshResponse, error) {
	tokens, err := s.gatewayService.Refresh(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &gatewaypb.RefreshResponse{
		TokenPair: model.TokenPairToPb(tokens),
	}, nil
}

func (s *Server) Logout(ctx context.Context, req *gatewaypb.LogoutRequest) (*emptypb.Empty, error) {
	err := s.gatewayService.Logout(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) ValidateToken(ctx context.Context, req *gatewaypb.ValidateTokenRequest) (*gatewaypb.ValidateTokenResponse, error) {
	userID, isValid, err := s.gatewayService.ValidateToken(ctx, req.GetAccessToken())
	if err != nil {
		return nil, err
	}

	return &gatewaypb.ValidateTokenResponse{
		UserId:  userID,
		IsValid: isValid,
	}, nil
}

func (s *Server) CreateUser(ctx context.Context, req *gatewaypb.CreateUserRequest) (*emptypb.Empty, error) {
	user := model.PbToUserCreate(req.GetUser())
	if err := s.gatewayService.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) GetUser(ctx context.Context, req *gatewaypb.GetUserRequest) (*gatewaypb.GetUserResponse, error) {
	res, err := s.gatewayService.GetUser(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &gatewaypb.GetUserResponse{
		User: model.UserToPb(res),
	}, nil
}

func (s *Server) GetCurrentUser(ctx context.Context, req *emptypb.Empty) (*gatewaypb.GetCurrentUserResponse, error) {
	userID, err := interceptor.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.gatewayService.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &gatewaypb.GetCurrentUserResponse{
		User: model.UserToPb(res),
	}, nil
}

func (s *Server) GetUsers(ctx context.Context, req *gatewaypb.GetUsersRequest) (*gatewaypb.GetUsersResponse, error) {
	pagination := req.GetPagination()
	limit := int(pagination.GetLimit())
	offset := int(pagination.GetOffset())

	res, err := s.gatewayService.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return &gatewaypb.GetUsersResponse{
		Users:      model.UsersToPbs(res),
		Pagination: pagination,
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *gatewaypb.UpdateUserRequest) (*emptypb.Empty, error) {
	user := model.PbToUserUpdate(req.GetUser())
	if err := s.gatewayService.UpdateUser(ctx, req.GetUserId(), user); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) UpdateCurrentUser(ctx context.Context, req *gatewaypb.UpdateCurrentUserRequest) (*emptypb.Empty, error) {
	userID, err := interceptor.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user := model.PbToUserUpdate(req.GetUser())
	if err := s.gatewayService.UpdateUser(ctx, userID, user); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteUser(ctx context.Context, req *gatewaypb.DeleteUserRequest) (*emptypb.Empty, error) {
	if err := s.gatewayService.DeleteUser(ctx, req.GetUserId()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteCurrentUser(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	userID, err := interceptor.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if err := s.gatewayService.DeleteUser(ctx, userID); err != nil {
		return nil, err
	}
	
	return &emptypb.Empty{}, nil
}
