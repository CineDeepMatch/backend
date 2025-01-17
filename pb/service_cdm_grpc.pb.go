// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: service_cdm.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CineDeepMatchClient is the client API for CineDeepMatch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CineDeepMatchClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	CreateActivity(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*CreateActivityResponse, error)
	GetActivities(ctx context.Context, in *GetActivitiesRequest, opts ...grpc.CallOption) (*GetActivitiesResponse, error)
	GetMovie(ctx context.Context, in *GetMovieRequest, opts ...grpc.CallOption) (*GetMovieResponse, error)
	GetManyMovies(ctx context.Context, in *GetManyMoviesRequest, opts ...grpc.CallOption) (*GetManyMoviesResponse, error)
	CreateFavMovies(ctx context.Context, in *CreateFavMoviesRequest, opts ...grpc.CallOption) (*CreateFavMoviesResponse, error)
	GetFavMovies(ctx context.Context, in *GetFavMoviesRequest, opts ...grpc.CallOption) (*GetFavMoviesResponse, error)
	UpdateFavMovies(ctx context.Context, in *UpdateFavMoviesRequest, opts ...grpc.CallOption) (*UpdateFavMoviesResponse, error)
	VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error)
	RenewAccessToken(ctx context.Context, in *RenewAccessTokenRequest, opts ...grpc.CallOption) (*RenewAccessTokenResponse, error)
}

type cineDeepMatchClient struct {
	cc grpc.ClientConnInterface
}

func NewCineDeepMatchClient(cc grpc.ClientConnInterface) CineDeepMatchClient {
	return &cineDeepMatchClient{cc}
}

func (c *cineDeepMatchClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/pb.CineDeepMatch/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cineDeepMatchClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/pb.CineDeepMatch/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cineDeepMatchClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, "/pb.CineDeepMatch/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cineDeepMatchClient) CreateActivity(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*CreateActivityResponse, error) {
	out := new(CreateActivityResponse)
	err := c.cc.Invoke(ctx, "/pb.CineDeepMatch/CreateActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cineDeepMatchClient) GetActivities(ctx context.Context, in *GetActivitiesRequest, opts ...grpc.CallOption) (*GetActivitiesResponse, error) {
	out := new(GetActivitiesResponse)
	err := c.cc.Invoke(ctx, "/pb.CineDeepMatch/GetActivities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cineDeepMatchClient) GetMovie(ctx context.Context, in *GetMovieRequest, opts ...grpc.CallOption) (*GetMovieResponse, error) {
	out := new(GetMovieResponse)
	err := c.cc.Invoke(ctx, "/pb.CineDeepMatch/GetMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cineDeepMatchClient) GetManyMovies(ctx context.Context, in *GetManyMoviesRequest, opts ...grpc.CallOption) (*GetManyMoviesResponse, error) {
	out := new(GetManyMoviesResponse)
	err := c.cc.Invoke(ctx, "/pb.CineDeepMatch/GetManyMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cineDeepMatchClient) CreateFavMovies(ctx context.Context, in *CreateFavMoviesRequest, opts ...grpc.CallOption) (*CreateFavMoviesResponse, error) {
	out := new(CreateFavMoviesResponse)
	err := c.cc.Invoke(ctx, "/pb.CineDeepMatch/CreateFavMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cineDeepMatchClient) GetFavMovies(ctx context.Context, in *GetFavMoviesRequest, opts ...grpc.CallOption) (*GetFavMoviesResponse, error) {
	out := new(GetFavMoviesResponse)
	err := c.cc.Invoke(ctx, "/pb.CineDeepMatch/GetFavMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cineDeepMatchClient) UpdateFavMovies(ctx context.Context, in *UpdateFavMoviesRequest, opts ...grpc.CallOption) (*UpdateFavMoviesResponse, error) {
	out := new(UpdateFavMoviesResponse)
	err := c.cc.Invoke(ctx, "/pb.CineDeepMatch/UpdateFavMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cineDeepMatchClient) VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error) {
	out := new(VerifyEmailResponse)
	err := c.cc.Invoke(ctx, "/pb.CineDeepMatch/VerifyEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cineDeepMatchClient) RenewAccessToken(ctx context.Context, in *RenewAccessTokenRequest, opts ...grpc.CallOption) (*RenewAccessTokenResponse, error) {
	out := new(RenewAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/pb.CineDeepMatch/RenewAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CineDeepMatchServer is the server API for CineDeepMatch service.
// All implementations must embed UnimplementedCineDeepMatchServer
// for forward compatibility
type CineDeepMatchServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	CreateActivity(context.Context, *CreateActivityRequest) (*CreateActivityResponse, error)
	GetActivities(context.Context, *GetActivitiesRequest) (*GetActivitiesResponse, error)
	GetMovie(context.Context, *GetMovieRequest) (*GetMovieResponse, error)
	GetManyMovies(context.Context, *GetManyMoviesRequest) (*GetManyMoviesResponse, error)
	CreateFavMovies(context.Context, *CreateFavMoviesRequest) (*CreateFavMoviesResponse, error)
	GetFavMovies(context.Context, *GetFavMoviesRequest) (*GetFavMoviesResponse, error)
	UpdateFavMovies(context.Context, *UpdateFavMoviesRequest) (*UpdateFavMoviesResponse, error)
	VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error)
	RenewAccessToken(context.Context, *RenewAccessTokenRequest) (*RenewAccessTokenResponse, error)
	mustEmbedUnimplementedCineDeepMatchServer()
}

// UnimplementedCineDeepMatchServer must be embedded to have forward compatible implementations.
type UnimplementedCineDeepMatchServer struct {
}

func (UnimplementedCineDeepMatchServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedCineDeepMatchServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedCineDeepMatchServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedCineDeepMatchServer) CreateActivity(context.Context, *CreateActivityRequest) (*CreateActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateActivity not implemented")
}
func (UnimplementedCineDeepMatchServer) GetActivities(context.Context, *GetActivitiesRequest) (*GetActivitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivities not implemented")
}
func (UnimplementedCineDeepMatchServer) GetMovie(context.Context, *GetMovieRequest) (*GetMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovie not implemented")
}
func (UnimplementedCineDeepMatchServer) GetManyMovies(context.Context, *GetManyMoviesRequest) (*GetManyMoviesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManyMovies not implemented")
}
func (UnimplementedCineDeepMatchServer) CreateFavMovies(context.Context, *CreateFavMoviesRequest) (*CreateFavMoviesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFavMovies not implemented")
}
func (UnimplementedCineDeepMatchServer) GetFavMovies(context.Context, *GetFavMoviesRequest) (*GetFavMoviesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavMovies not implemented")
}
func (UnimplementedCineDeepMatchServer) UpdateFavMovies(context.Context, *UpdateFavMoviesRequest) (*UpdateFavMoviesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFavMovies not implemented")
}
func (UnimplementedCineDeepMatchServer) VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmail not implemented")
}
func (UnimplementedCineDeepMatchServer) RenewAccessToken(context.Context, *RenewAccessTokenRequest) (*RenewAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewAccessToken not implemented")
}
func (UnimplementedCineDeepMatchServer) mustEmbedUnimplementedCineDeepMatchServer() {}

// UnsafeCineDeepMatchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CineDeepMatchServer will
// result in compilation errors.
type UnsafeCineDeepMatchServer interface {
	mustEmbedUnimplementedCineDeepMatchServer()
}

func RegisterCineDeepMatchServer(s grpc.ServiceRegistrar, srv CineDeepMatchServer) {
	s.RegisterService(&CineDeepMatch_ServiceDesc, srv)
}

func _CineDeepMatch_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CineDeepMatchServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CineDeepMatch/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CineDeepMatchServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CineDeepMatch_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CineDeepMatchServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CineDeepMatch/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CineDeepMatchServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CineDeepMatch_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CineDeepMatchServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CineDeepMatch/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CineDeepMatchServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CineDeepMatch_CreateActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CineDeepMatchServer).CreateActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CineDeepMatch/CreateActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CineDeepMatchServer).CreateActivity(ctx, req.(*CreateActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CineDeepMatch_GetActivities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActivitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CineDeepMatchServer).GetActivities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CineDeepMatch/GetActivities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CineDeepMatchServer).GetActivities(ctx, req.(*GetActivitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CineDeepMatch_GetMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CineDeepMatchServer).GetMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CineDeepMatch/GetMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CineDeepMatchServer).GetMovie(ctx, req.(*GetMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CineDeepMatch_GetManyMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManyMoviesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CineDeepMatchServer).GetManyMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CineDeepMatch/GetManyMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CineDeepMatchServer).GetManyMovies(ctx, req.(*GetManyMoviesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CineDeepMatch_CreateFavMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFavMoviesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CineDeepMatchServer).CreateFavMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CineDeepMatch/CreateFavMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CineDeepMatchServer).CreateFavMovies(ctx, req.(*CreateFavMoviesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CineDeepMatch_GetFavMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFavMoviesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CineDeepMatchServer).GetFavMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CineDeepMatch/GetFavMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CineDeepMatchServer).GetFavMovies(ctx, req.(*GetFavMoviesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CineDeepMatch_UpdateFavMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFavMoviesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CineDeepMatchServer).UpdateFavMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CineDeepMatch/UpdateFavMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CineDeepMatchServer).UpdateFavMovies(ctx, req.(*UpdateFavMoviesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CineDeepMatch_VerifyEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CineDeepMatchServer).VerifyEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CineDeepMatch/VerifyEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CineDeepMatchServer).VerifyEmail(ctx, req.(*VerifyEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CineDeepMatch_RenewAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenewAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CineDeepMatchServer).RenewAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CineDeepMatch/RenewAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CineDeepMatchServer).RenewAccessToken(ctx, req.(*RenewAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CineDeepMatch_ServiceDesc is the grpc.ServiceDesc for CineDeepMatch service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CineDeepMatch_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CineDeepMatch",
	HandlerType: (*CineDeepMatchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _CineDeepMatch_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _CineDeepMatch_UpdateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _CineDeepMatch_LoginUser_Handler,
		},
		{
			MethodName: "CreateActivity",
			Handler:    _CineDeepMatch_CreateActivity_Handler,
		},
		{
			MethodName: "GetActivities",
			Handler:    _CineDeepMatch_GetActivities_Handler,
		},
		{
			MethodName: "GetMovie",
			Handler:    _CineDeepMatch_GetMovie_Handler,
		},
		{
			MethodName: "GetManyMovies",
			Handler:    _CineDeepMatch_GetManyMovies_Handler,
		},
		{
			MethodName: "CreateFavMovies",
			Handler:    _CineDeepMatch_CreateFavMovies_Handler,
		},
		{
			MethodName: "GetFavMovies",
			Handler:    _CineDeepMatch_GetFavMovies_Handler,
		},
		{
			MethodName: "UpdateFavMovies",
			Handler:    _CineDeepMatch_UpdateFavMovies_Handler,
		},
		{
			MethodName: "VerifyEmail",
			Handler:    _CineDeepMatch_VerifyEmail_Handler,
		},
		{
			MethodName: "RenewAccessToken",
			Handler:    _CineDeepMatch_RenewAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_cdm.proto",
}
