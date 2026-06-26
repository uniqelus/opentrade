package userapi

import (
	"context"
	"errors"

	idp "github.com/uniqelus/opentrade/api/gen/go/identity_provider/v1"
	sdkgrpc "github.com/uniqelus/opentrade/sdk/go/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	userdmn "github.com/uniqelus/opentrade/identity-provider/internal/domains/user"
)

type UserService interface {
	userdmn.Creator
}

type api struct {
	idp.UnimplementedUserServiceServer
	userService UserService
}

func NewServiceRegistration(userService UserService) sdkgrpc.ServiceRegistration {
	return func(sr grpc.ServiceRegistrar) {
		idp.RegisterUserServiceServer(sr, &api{userService: userService})
	}
}

func (a *api) GetUser(context.Context, *idp.GetUserRequest) (*idp.User, error) {
	return nil, status.Error(codes.Unimplemented, "method GetUser not implemented")
}

func (a *api) ListUsers(context.Context, *idp.ListUsersRequest) (*idp.ListUsersResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method ListUsers not implemented")
}

func (a *api) DeleteUser(context.Context, *idp.DeleteUserRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "method DeleteUser not implemented")
}

func (a *api) CreateUser(ctx context.Context, req *idp.CreateUserRequest) (*idp.User, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := a.userService.CreateUser(ctx,
		userdmn.WithCreateUserFirstName(req.GetUser().GetFirstName()),
		userdmn.WithCreateUserFirstName(req.GetUser().GetFirstName()),
		userdmn.WithCreateUserLastName(req.GetUser().GetLastName()),
	)
	if err != nil {
		return nil, convertErrToProto(err)
	}

	return convertUserToProto(user), nil
}

func (a *api) UpdateUser(context.Context, *idp.UpdateUserRequest) (*idp.User, error) {
	return nil, status.Error(codes.Unimplemented, "method UpdateUser not implemented")
}

func convertErrToProto(err error) error {
	if errors.Is(err, userdmn.ErrUserNotFound) {
		return status.Error(codes.NotFound, err.Error())
	}

	return status.Error(codes.Internal, "internal error")
}

func convertUserToProto(user *userdmn.User) *idp.User {
	return &idp.User{
		Name:       user.Name.String(),
		CreateTime: timestamppb.New(user.CreateTime),
		UpdateTime: timestamppb.New(user.UpdateTime),
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		State:      idp.UserState(user.State),
	}
}
