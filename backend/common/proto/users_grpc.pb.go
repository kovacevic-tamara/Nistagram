// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UsersDTO, error)
	GetUserById(ctx context.Context, in *RequestIdUsers, opts ...grpc.CallOption) (*UsersDTO, error)
	GetUsernameById(ctx context.Context, in *RequestIdUsers, opts ...grpc.CallOption) (*UsersDTO, error)
	GetAllUsers(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*UsersResponse, error)
	UpdateUserProfile(ctx context.Context, in *CreateUserDTORequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	UpdateUserPassword(ctx context.Context, in *CreatePasswordRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	SearchUser(ctx context.Context, in *SearchUserDtoRequest, opts ...grpc.CallOption) (*UsersResponse, error)
	// Svi
	SendEmail(ctx context.Context, in *SendEmailDtoRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	// Svi
	GetUserByEmail(ctx context.Context, in *RequestEmailUser, opts ...grpc.CallOption) (*UsersDTO, error)
	GetUserByUsername(ctx context.Context, in *RequestUsernameUser, opts ...grpc.CallOption) (*UsersDTO, error)
	// Svi
	ValidateResetCode(ctx context.Context, in *RequestResetCode, opts ...grpc.CallOption) (*EmptyResponse, error)
	// Svi
	ChangeForgottenPass(ctx context.Context, in *CreatePasswordRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	CheckIsApproved(ctx context.Context, in *RequestIdUsers, opts ...grpc.CallOption) (*BooleanResponseUsers, error)
	// Samo ulogovani
	ApproveAccount(ctx context.Context, in *CreatePasswordRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	GoogleAuth(ctx context.Context, in *GoogleAuthRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	//* Verification requests *
	SubmitVerificationRequest(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetPendingVerificationRequests(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*VerificationRequestsArray, error)
	ChangeVerificationRequestStatus(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/proto.Users/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UsersDTO, error) {
	out := new(UsersDTO)
	err := c.cc.Invoke(ctx, "/proto.Users/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUserById(ctx context.Context, in *RequestIdUsers, opts ...grpc.CallOption) (*UsersDTO, error) {
	out := new(UsersDTO)
	err := c.cc.Invoke(ctx, "/proto.Users/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUsernameById(ctx context.Context, in *RequestIdUsers, opts ...grpc.CallOption) (*UsersDTO, error) {
	out := new(UsersDTO)
	err := c.cc.Invoke(ctx, "/proto.Users/GetUsernameById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetAllUsers(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, "/proto.Users/GetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdateUserProfile(ctx context.Context, in *CreateUserDTORequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.Users/UpdateUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdateUserPassword(ctx context.Context, in *CreatePasswordRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.Users/UpdateUserPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) SearchUser(ctx context.Context, in *SearchUserDtoRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, "/proto.Users/SearchUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) SendEmail(ctx context.Context, in *SendEmailDtoRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.Users/SendEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUserByEmail(ctx context.Context, in *RequestEmailUser, opts ...grpc.CallOption) (*UsersDTO, error) {
	out := new(UsersDTO)
	err := c.cc.Invoke(ctx, "/proto.Users/GetUserByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUserByUsername(ctx context.Context, in *RequestUsernameUser, opts ...grpc.CallOption) (*UsersDTO, error) {
	out := new(UsersDTO)
	err := c.cc.Invoke(ctx, "/proto.Users/GetUserByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ValidateResetCode(ctx context.Context, in *RequestResetCode, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.Users/ValidateResetCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ChangeForgottenPass(ctx context.Context, in *CreatePasswordRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.Users/ChangeForgottenPass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) CheckIsApproved(ctx context.Context, in *RequestIdUsers, opts ...grpc.CallOption) (*BooleanResponseUsers, error) {
	out := new(BooleanResponseUsers)
	err := c.cc.Invoke(ctx, "/proto.Users/CheckIsApproved", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ApproveAccount(ctx context.Context, in *CreatePasswordRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.Users/ApproveAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GoogleAuth(ctx context.Context, in *GoogleAuthRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/proto.Users/GoogleAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) SubmitVerificationRequest(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.Users/SubmitVerificationRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetPendingVerificationRequests(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*VerificationRequestsArray, error) {
	out := new(VerificationRequestsArray)
	err := c.cc.Invoke(ctx, "/proto.Users/GetPendingVerificationRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ChangeVerificationRequestStatus(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.Users/ChangeVerificationRequestStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
// All implementations must embed UnimplementedUsersServer
// for forward compatibility
type UsersServer interface {
	LoginUser(context.Context, *LoginRequest) (*LoginResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*UsersDTO, error)
	GetUserById(context.Context, *RequestIdUsers) (*UsersDTO, error)
	GetUsernameById(context.Context, *RequestIdUsers) (*UsersDTO, error)
	GetAllUsers(context.Context, *EmptyRequest) (*UsersResponse, error)
	UpdateUserProfile(context.Context, *CreateUserDTORequest) (*EmptyResponse, error)
	UpdateUserPassword(context.Context, *CreatePasswordRequest) (*EmptyResponse, error)
	SearchUser(context.Context, *SearchUserDtoRequest) (*UsersResponse, error)
	// Svi
	SendEmail(context.Context, *SendEmailDtoRequest) (*EmptyResponse, error)
	// Svi
	GetUserByEmail(context.Context, *RequestEmailUser) (*UsersDTO, error)
	GetUserByUsername(context.Context, *RequestUsernameUser) (*UsersDTO, error)
	// Svi
	ValidateResetCode(context.Context, *RequestResetCode) (*EmptyResponse, error)
	// Svi
	ChangeForgottenPass(context.Context, *CreatePasswordRequest) (*EmptyResponse, error)
	CheckIsApproved(context.Context, *RequestIdUsers) (*BooleanResponseUsers, error)
	// Samo ulogovani
	ApproveAccount(context.Context, *CreatePasswordRequest) (*EmptyResponse, error)
	GoogleAuth(context.Context, *GoogleAuthRequest) (*LoginResponse, error)
	//* Verification requests *
	SubmitVerificationRequest(context.Context, *VerificationRequest) (*EmptyResponse, error)
	GetPendingVerificationRequests(context.Context, *EmptyRequest) (*VerificationRequestsArray, error)
	ChangeVerificationRequestStatus(context.Context, *VerificationRequest) (*EmptyResponse, error)
	mustEmbedUnimplementedUsersServer()
}

// UnimplementedUsersServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (UnimplementedUsersServer) LoginUser(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedUsersServer) CreateUser(context.Context, *CreateUserRequest) (*UsersDTO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUsersServer) GetUserById(context.Context, *RequestIdUsers) (*UsersDTO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUsersServer) GetUsernameById(context.Context, *RequestIdUsers) (*UsersDTO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsernameById not implemented")
}
func (UnimplementedUsersServer) GetAllUsers(context.Context, *EmptyRequest) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedUsersServer) UpdateUserProfile(context.Context, *CreateUserDTORequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfile not implemented")
}
func (UnimplementedUsersServer) UpdateUserPassword(context.Context, *CreatePasswordRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserPassword not implemented")
}
func (UnimplementedUsersServer) SearchUser(context.Context, *SearchUserDtoRequest) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUser not implemented")
}
func (UnimplementedUsersServer) SendEmail(context.Context, *SendEmailDtoRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}
func (UnimplementedUsersServer) GetUserByEmail(context.Context, *RequestEmailUser) (*UsersDTO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByEmail not implemented")
}
func (UnimplementedUsersServer) GetUserByUsername(context.Context, *RequestUsernameUser) (*UsersDTO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByUsername not implemented")
}
func (UnimplementedUsersServer) ValidateResetCode(context.Context, *RequestResetCode) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateResetCode not implemented")
}
func (UnimplementedUsersServer) ChangeForgottenPass(context.Context, *CreatePasswordRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeForgottenPass not implemented")
}
func (UnimplementedUsersServer) CheckIsApproved(context.Context, *RequestIdUsers) (*BooleanResponseUsers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIsApproved not implemented")
}
func (UnimplementedUsersServer) ApproveAccount(context.Context, *CreatePasswordRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveAccount not implemented")
}
func (UnimplementedUsersServer) GoogleAuth(context.Context, *GoogleAuthRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoogleAuth not implemented")
}
func (UnimplementedUsersServer) SubmitVerificationRequest(context.Context, *VerificationRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitVerificationRequest not implemented")
}
func (UnimplementedUsersServer) GetPendingVerificationRequests(context.Context, *EmptyRequest) (*VerificationRequestsArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPendingVerificationRequests not implemented")
}
func (UnimplementedUsersServer) ChangeVerificationRequestStatus(context.Context, *VerificationRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeVerificationRequestStatus not implemented")
}
func (UnimplementedUsersServer) mustEmbedUnimplementedUsersServer() {}

// UnsafeUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServer will
// result in compilation errors.
type UnsafeUsersServer interface {
	mustEmbedUnimplementedUsersServer()
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Users/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).LoginUser(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Users/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestIdUsers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Users/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserById(ctx, req.(*RequestIdUsers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUsernameById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestIdUsers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUsernameById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Users/GetUsernameById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUsernameById(ctx, req.(*RequestIdUsers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Users/GetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetAllUsers(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdateUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserDTORequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdateUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Users/UpdateUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdateUserProfile(ctx, req.(*CreateUserDTORequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdateUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdateUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Users/UpdateUserPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdateUserPassword(ctx, req.(*CreatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_SearchUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserDtoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).SearchUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Users/SearchUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).SearchUser(ctx, req.(*SearchUserDtoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailDtoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Users/SendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).SendEmail(ctx, req.(*SendEmailDtoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEmailUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Users/GetUserByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserByEmail(ctx, req.(*RequestEmailUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUserByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUsernameUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Users/GetUserByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserByUsername(ctx, req.(*RequestUsernameUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ValidateResetCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestResetCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ValidateResetCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Users/ValidateResetCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ValidateResetCode(ctx, req.(*RequestResetCode))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ChangeForgottenPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ChangeForgottenPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Users/ChangeForgottenPass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ChangeForgottenPass(ctx, req.(*CreatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_CheckIsApproved_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestIdUsers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).CheckIsApproved(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Users/CheckIsApproved",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).CheckIsApproved(ctx, req.(*RequestIdUsers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ApproveAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ApproveAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Users/ApproveAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ApproveAccount(ctx, req.(*CreatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GoogleAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoogleAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GoogleAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Users/GoogleAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GoogleAuth(ctx, req.(*GoogleAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_SubmitVerificationRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).SubmitVerificationRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Users/SubmitVerificationRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).SubmitVerificationRequest(ctx, req.(*VerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetPendingVerificationRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetPendingVerificationRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Users/GetPendingVerificationRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetPendingVerificationRequests(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ChangeVerificationRequestStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ChangeVerificationRequestStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Users/ChangeVerificationRequestStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ChangeVerificationRequestStatus(ctx, req.(*VerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginUser",
			Handler:    _Users_LoginUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Users_CreateUser_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _Users_GetUserById_Handler,
		},
		{
			MethodName: "GetUsernameById",
			Handler:    _Users_GetUsernameById_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _Users_GetAllUsers_Handler,
		},
		{
			MethodName: "UpdateUserProfile",
			Handler:    _Users_UpdateUserProfile_Handler,
		},
		{
			MethodName: "UpdateUserPassword",
			Handler:    _Users_UpdateUserPassword_Handler,
		},
		{
			MethodName: "SearchUser",
			Handler:    _Users_SearchUser_Handler,
		},
		{
			MethodName: "SendEmail",
			Handler:    _Users_SendEmail_Handler,
		},
		{
			MethodName: "GetUserByEmail",
			Handler:    _Users_GetUserByEmail_Handler,
		},
		{
			MethodName: "GetUserByUsername",
			Handler:    _Users_GetUserByUsername_Handler,
		},
		{
			MethodName: "ValidateResetCode",
			Handler:    _Users_ValidateResetCode_Handler,
		},
		{
			MethodName: "ChangeForgottenPass",
			Handler:    _Users_ChangeForgottenPass_Handler,
		},
		{
			MethodName: "CheckIsApproved",
			Handler:    _Users_CheckIsApproved_Handler,
		},
		{
			MethodName: "ApproveAccount",
			Handler:    _Users_ApproveAccount_Handler,
		},
		{
			MethodName: "GoogleAuth",
			Handler:    _Users_GoogleAuth_Handler,
		},
		{
			MethodName: "SubmitVerificationRequest",
			Handler:    _Users_SubmitVerificationRequest_Handler,
		},
		{
			MethodName: "GetPendingVerificationRequests",
			Handler:    _Users_GetPendingVerificationRequests_Handler,
		},
		{
			MethodName: "ChangeVerificationRequestStatus",
			Handler:    _Users_ChangeVerificationRequestStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
