// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/auth.proto

/*
Package org_dakstudios_srv_auth_auth is a generated protocol buffer package.

It is generated from these files:
	proto/auth/auth.proto

It has these top-level messages:
	Token
	AuthenticateRequest
	AuthenticateResponse
	AuthorizeRequest
	AuthorizeResponse
*/
package org_dakstudios_srv_auth_auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Token struct {
	// jwt token
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthenticateRequest struct {
	// user email
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	// user password
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *AuthenticateRequest) Reset()                    { *m = AuthenticateRequest{} }
func (m *AuthenticateRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthenticateRequest) ProtoMessage()               {}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AuthenticateRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthenticateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthenticateResponse struct {
	// jwt token
	Token *Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *AuthenticateResponse) Reset()                    { *m = AuthenticateResponse{} }
func (m *AuthenticateResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthenticateResponse) ProtoMessage()               {}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AuthenticateResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type AuthorizeRequest struct {
	// jwt token
	Token *Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	// permission
	Permission string `protobuf:"bytes,2,opt,name=permission" json:"permission,omitempty"`
}

func (m *AuthorizeRequest) Reset()                    { *m = AuthorizeRequest{} }
func (m *AuthorizeRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthorizeRequest) ProtoMessage()               {}
func (*AuthorizeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AuthorizeRequest) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *AuthorizeRequest) GetPermission() string {
	if m != nil {
		return m.Permission
	}
	return ""
}

type AuthorizeResponse struct {
	// true if user authorized for the action
	Authorized bool `protobuf:"varint,1,opt,name=authorized" json:"authorized,omitempty"`
}

func (m *AuthorizeResponse) Reset()                    { *m = AuthorizeResponse{} }
func (m *AuthorizeResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthorizeResponse) ProtoMessage()               {}
func (*AuthorizeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AuthorizeResponse) GetAuthorized() bool {
	if m != nil {
		return m.Authorized
	}
	return false
}

func init() {
	proto.RegisterType((*Token)(nil), "org.dakstudios.srv.auth.auth.Token")
	proto.RegisterType((*AuthenticateRequest)(nil), "org.dakstudios.srv.auth.auth.AuthenticateRequest")
	proto.RegisterType((*AuthenticateResponse)(nil), "org.dakstudios.srv.auth.auth.AuthenticateResponse")
	proto.RegisterType((*AuthorizeRequest)(nil), "org.dakstudios.srv.auth.auth.AuthorizeRequest")
	proto.RegisterType((*AuthorizeResponse)(nil), "org.dakstudios.srv.auth.auth.AuthorizeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthClient interface {
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...client.CallOption) (*AuthenticateResponse, error)
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*AuthorizeResponse, error)
}

type authClient struct {
	c           client.Client
	serviceName string
}

func NewAuthClient(serviceName string, c client.Client) AuthClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "org.dakstudios.srv.auth.auth"
	}
	return &authClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...client.CallOption) (*AuthenticateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Authenticate", in)
	out := new(AuthenticateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*AuthorizeResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Authorize", in)
	out := new(AuthorizeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	Authenticate(context.Context, *AuthenticateRequest, *AuthenticateResponse) error
	Authorize(context.Context, *AuthorizeRequest, *AuthorizeResponse) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Auth{hdlr}, opts...))
}

type Auth struct {
	AuthHandler
}

func (h *Auth) Authenticate(ctx context.Context, in *AuthenticateRequest, out *AuthenticateResponse) error {
	return h.AuthHandler.Authenticate(ctx, in, out)
}

func (h *Auth) Authorize(ctx context.Context, in *AuthorizeRequest, out *AuthorizeResponse) error {
	return h.AuthHandler.Authorize(ctx, in, out)
}

func init() { proto.RegisterFile("proto/auth/auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x5d, 0xb1, 0xd2, 0x8e, 0x1e, 0x74, 0xac, 0x50, 0x16, 0x2d, 0x12, 0x2f, 0x9e, 0x52,
	0xdc, 0x9e, 0x3c, 0x7a, 0xf2, 0xec, 0xe2, 0x0b, 0x44, 0x37, 0xd8, 0x50, 0x37, 0xb3, 0x66, 0xb2,
	0x16, 0x7c, 0x69, 0x5f, 0x41, 0x92, 0x76, 0x97, 0x08, 0x52, 0xff, 0x5c, 0xc2, 0x7e, 0x93, 0x99,
	0xef, 0xfb, 0xed, 0x10, 0x38, 0x6d, 0x1c, 0x79, 0x9a, 0xa9, 0xd6, 0x2f, 0xe2, 0x21, 0xa3, 0xc6,
	0x33, 0x72, 0xcf, 0xb2, 0x52, 0x4b, 0xf6, 0x6d, 0x65, 0x88, 0x25, 0xbb, 0x37, 0x19, 0xaf, 0xc3,
	0x21, 0xce, 0x61, 0xf0, 0x40, 0x4b, 0x6d, 0x71, 0x0c, 0x03, 0x1f, 0x3e, 0x26, 0xd9, 0x45, 0x76,
	0x35, 0x2a, 0xd7, 0x42, 0xdc, 0xc1, 0xc9, 0x6d, 0xeb, 0x17, 0xda, 0x7a, 0xf3, 0xa4, 0xbc, 0x2e,
	0xf5, 0x6b, 0xab, 0xd9, 0x87, 0x66, 0x5d, 0x2b, 0xf3, 0xd2, 0x35, 0x47, 0x81, 0x39, 0x0c, 0x1b,
	0xc5, 0xbc, 0x22, 0x57, 0x4d, 0x76, 0xe3, 0x45, 0xaf, 0xc5, 0x3d, 0x8c, 0xbf, 0x1a, 0x71, 0x43,
	0x96, 0x35, 0xde, 0xa4, 0xb1, 0x07, 0xc5, 0xa5, 0xdc, 0x46, 0x2b, 0x23, 0x6a, 0xc7, 0x56, 0xc3,
	0x51, 0xb0, 0x24, 0x67, 0xde, 0x7b, 0xb0, 0xff, 0xdb, 0xe1, 0x14, 0xa0, 0xd1, 0xae, 0x36, 0xcc,
	0x86, 0xec, 0x86, 0x3f, 0xa9, 0x88, 0x39, 0x1c, 0x27, 0x71, 0x1b, 0xfc, 0x29, 0x80, 0xea, 0x8a,
	0x55, 0x0c, 0x1d, 0x96, 0x49, 0xa5, 0xf8, 0xc8, 0x60, 0x2f, 0x4c, 0xe1, 0x0a, 0x0e, 0xd3, 0xff,
	0xc7, 0xeb, 0xed, 0x64, 0xdf, 0x2c, 0x3d, 0x2f, 0xfe, 0x32, 0xb2, 0xe6, 0x13, 0x3b, 0x68, 0x61,
	0xd4, 0x63, 0xa3, 0xfc, 0xd9, 0x22, 0x5d, 0x67, 0x3e, 0xfb, 0x75, 0x7f, 0x97, 0xf7, 0xb8, 0x1f,
	0x5f, 0xdd, 0xfc, 0x33, 0x00, 0x00, 0xff, 0xff, 0x23, 0x62, 0x0c, 0xab, 0x8e, 0x02, 0x00, 0x00,
}
