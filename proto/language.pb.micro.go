// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: language.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for LanguageService service

type LanguageService interface {
	Create(ctx context.Context, in *Language, opts ...client.CallOption) (*DatabaseResponse, error)
	Update(ctx context.Context, in *Language, opts ...client.CallOption) (*DatabaseResponse, error)
	SoftDelete(ctx context.Context, in *Language, opts ...client.CallOption) (*DatabaseResponse, error)
	ForceDelete(ctx context.Context, in *Language, opts ...client.CallOption) (*DatabaseResponse, error)
	FindAll(ctx context.Context, in *DatabaseEmpty, opts ...client.CallOption) (*DatabaseResponse, error)
	FindOneByISO(ctx context.Context, in *DatabaseString, opts ...client.CallOption) (*DatabaseResponse, error)
	FindManyByISO(ctx context.Context, in *DatabaseStrings, opts ...client.CallOption) (*DatabaseResponse, error)
}

type languageService struct {
	c    client.Client
	name string
}

func NewLanguageService(name string, c client.Client) LanguageService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "proto"
	}
	return &languageService{
		c:    c,
		name: name,
	}
}

func (c *languageService) Create(ctx context.Context, in *Language, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "LanguageService.Create", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageService) Update(ctx context.Context, in *Language, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "LanguageService.Update", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageService) SoftDelete(ctx context.Context, in *Language, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "LanguageService.SoftDelete", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageService) ForceDelete(ctx context.Context, in *Language, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "LanguageService.ForceDelete", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageService) FindAll(ctx context.Context, in *DatabaseEmpty, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "LanguageService.FindAll", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageService) FindOneByISO(ctx context.Context, in *DatabaseString, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "LanguageService.FindOneByISO", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageService) FindManyByISO(ctx context.Context, in *DatabaseStrings, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "LanguageService.FindManyByISO", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LanguageService service

type LanguageServiceHandler interface {
	Create(context.Context, *Language, *DatabaseResponse) error
	Update(context.Context, *Language, *DatabaseResponse) error
	SoftDelete(context.Context, *Language, *DatabaseResponse) error
	ForceDelete(context.Context, *Language, *DatabaseResponse) error
	FindAll(context.Context, *DatabaseEmpty, *DatabaseResponse) error
	FindOneByISO(context.Context, *DatabaseString, *DatabaseResponse) error
	FindManyByISO(context.Context, *DatabaseStrings, *DatabaseResponse) error
}

func RegisterLanguageServiceHandler(s server.Server, hdlr LanguageServiceHandler, opts ...server.HandlerOption) error {
	type languageService interface {
		Create(ctx context.Context, in *Language, out *DatabaseResponse) error
		Update(ctx context.Context, in *Language, out *DatabaseResponse) error
		SoftDelete(ctx context.Context, in *Language, out *DatabaseResponse) error
		ForceDelete(ctx context.Context, in *Language, out *DatabaseResponse) error
		FindAll(ctx context.Context, in *DatabaseEmpty, out *DatabaseResponse) error
		FindOneByISO(ctx context.Context, in *DatabaseString, out *DatabaseResponse) error
		FindManyByISO(ctx context.Context, in *DatabaseStrings, out *DatabaseResponse) error
	}
	type LanguageService struct {
		languageService
	}
	h := &languageServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&LanguageService{h}, opts...))
}

type languageServiceHandler struct {
	LanguageServiceHandler
}

func (h *languageServiceHandler) Create(ctx context.Context, in *Language, out *DatabaseResponse) error {
	return h.LanguageServiceHandler.Create(ctx, in, out)
}

func (h *languageServiceHandler) Update(ctx context.Context, in *Language, out *DatabaseResponse) error {
	return h.LanguageServiceHandler.Update(ctx, in, out)
}

func (h *languageServiceHandler) SoftDelete(ctx context.Context, in *Language, out *DatabaseResponse) error {
	return h.LanguageServiceHandler.SoftDelete(ctx, in, out)
}

func (h *languageServiceHandler) ForceDelete(ctx context.Context, in *Language, out *DatabaseResponse) error {
	return h.LanguageServiceHandler.ForceDelete(ctx, in, out)
}

func (h *languageServiceHandler) FindAll(ctx context.Context, in *DatabaseEmpty, out *DatabaseResponse) error {
	return h.LanguageServiceHandler.FindAll(ctx, in, out)
}

func (h *languageServiceHandler) FindOneByISO(ctx context.Context, in *DatabaseString, out *DatabaseResponse) error {
	return h.LanguageServiceHandler.FindOneByISO(ctx, in, out)
}

func (h *languageServiceHandler) FindManyByISO(ctx context.Context, in *DatabaseStrings, out *DatabaseResponse) error {
	return h.LanguageServiceHandler.FindManyByISO(ctx, in, out)
}
