// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: movie.proto

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

// Client API for MovieService service

type MovieService interface {
	Create(ctx context.Context, in *Movie, opts ...client.CallOption) (*DatabaseResponse, error)
	Update(ctx context.Context, in *Movie, opts ...client.CallOption) (*DatabaseResponse, error)
	SoftDelete(ctx context.Context, in *Movie, opts ...client.CallOption) (*DatabaseResponse, error)
	ForceDelete(ctx context.Context, in *Movie, opts ...client.CallOption) (*DatabaseResponse, error)
	FindAll(ctx context.Context, in *DatabaseEmpty, opts ...client.CallOption) (*DatabaseResponse, error)
	FindOnlyIDs(ctx context.Context, in *DatabaseEmpty, opts ...client.CallOption) (*DatabaseResponse, error)
	FindDownloaded(ctx context.Context, in *DatabaseEmpty, opts ...client.CallOption) (*DatabaseResponse, error)
	FindOneByID(ctx context.Context, in *DatabaseID, opts ...client.CallOption) (*DatabaseResponse, error)
	FindManyByID(ctx context.Context, in *DatabaseIDs, opts ...client.CallOption) (*DatabaseResponse, error)
	FindOneByTitle(ctx context.Context, in *DatabaseString, opts ...client.CallOption) (*DatabaseResponse, error)
	FindManyByTitle(ctx context.Context, in *DatabaseStrings, opts ...client.CallOption) (*DatabaseResponse, error)
	FindOneByOriginalTitle(ctx context.Context, in *DatabaseString, opts ...client.CallOption) (*DatabaseResponse, error)
	FindManyByOriginalTitle(ctx context.Context, in *DatabaseStrings, opts ...client.CallOption) (*DatabaseResponse, error)
	FindOneByLocalTitle(ctx context.Context, in *DatabaseString, opts ...client.CallOption) (*DatabaseResponse, error)
	FindManyByLocalTitle(ctx context.Context, in *DatabaseStrings, opts ...client.CallOption) (*DatabaseResponse, error)
}

type movieService struct {
	c    client.Client
	name string
}

func NewMovieService(name string, c client.Client) MovieService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "proto"
	}
	return &movieService{
		c:    c,
		name: name,
	}
}

func (c *movieService) Create(ctx context.Context, in *Movie, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "MovieService.Create", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) Update(ctx context.Context, in *Movie, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "MovieService.Update", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) SoftDelete(ctx context.Context, in *Movie, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "MovieService.SoftDelete", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) ForceDelete(ctx context.Context, in *Movie, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "MovieService.ForceDelete", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) FindAll(ctx context.Context, in *DatabaseEmpty, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "MovieService.FindAll", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) FindOnlyIDs(ctx context.Context, in *DatabaseEmpty, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "MovieService.FindOnlyIDs", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) FindDownloaded(ctx context.Context, in *DatabaseEmpty, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "MovieService.FindDownloaded", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) FindOneByID(ctx context.Context, in *DatabaseID, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "MovieService.FindOneByID", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) FindManyByID(ctx context.Context, in *DatabaseIDs, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "MovieService.FindManyByID", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) FindOneByTitle(ctx context.Context, in *DatabaseString, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "MovieService.FindOneByTitle", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) FindManyByTitle(ctx context.Context, in *DatabaseStrings, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "MovieService.FindManyByTitle", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) FindOneByOriginalTitle(ctx context.Context, in *DatabaseString, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "MovieService.FindOneByOriginalTitle", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) FindManyByOriginalTitle(ctx context.Context, in *DatabaseStrings, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "MovieService.FindManyByOriginalTitle", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) FindOneByLocalTitle(ctx context.Context, in *DatabaseString, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "MovieService.FindOneByLocalTitle", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) FindManyByLocalTitle(ctx context.Context, in *DatabaseStrings, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "MovieService.FindManyByLocalTitle", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MovieService service

type MovieServiceHandler interface {
	Create(context.Context, *Movie, *DatabaseResponse) error
	Update(context.Context, *Movie, *DatabaseResponse) error
	SoftDelete(context.Context, *Movie, *DatabaseResponse) error
	ForceDelete(context.Context, *Movie, *DatabaseResponse) error
	FindAll(context.Context, *DatabaseEmpty, *DatabaseResponse) error
	FindOnlyIDs(context.Context, *DatabaseEmpty, *DatabaseResponse) error
	FindDownloaded(context.Context, *DatabaseEmpty, *DatabaseResponse) error
	FindOneByID(context.Context, *DatabaseID, *DatabaseResponse) error
	FindManyByID(context.Context, *DatabaseIDs, *DatabaseResponse) error
	FindOneByTitle(context.Context, *DatabaseString, *DatabaseResponse) error
	FindManyByTitle(context.Context, *DatabaseStrings, *DatabaseResponse) error
	FindOneByOriginalTitle(context.Context, *DatabaseString, *DatabaseResponse) error
	FindManyByOriginalTitle(context.Context, *DatabaseStrings, *DatabaseResponse) error
	FindOneByLocalTitle(context.Context, *DatabaseString, *DatabaseResponse) error
	FindManyByLocalTitle(context.Context, *DatabaseStrings, *DatabaseResponse) error
}

func RegisterMovieServiceHandler(s server.Server, hdlr MovieServiceHandler, opts ...server.HandlerOption) error {
	type movieService interface {
		Create(ctx context.Context, in *Movie, out *DatabaseResponse) error
		Update(ctx context.Context, in *Movie, out *DatabaseResponse) error
		SoftDelete(ctx context.Context, in *Movie, out *DatabaseResponse) error
		ForceDelete(ctx context.Context, in *Movie, out *DatabaseResponse) error
		FindAll(ctx context.Context, in *DatabaseEmpty, out *DatabaseResponse) error
		FindOnlyIDs(ctx context.Context, in *DatabaseEmpty, out *DatabaseResponse) error
		FindDownloaded(ctx context.Context, in *DatabaseEmpty, out *DatabaseResponse) error
		FindOneByID(ctx context.Context, in *DatabaseID, out *DatabaseResponse) error
		FindManyByID(ctx context.Context, in *DatabaseIDs, out *DatabaseResponse) error
		FindOneByTitle(ctx context.Context, in *DatabaseString, out *DatabaseResponse) error
		FindManyByTitle(ctx context.Context, in *DatabaseStrings, out *DatabaseResponse) error
		FindOneByOriginalTitle(ctx context.Context, in *DatabaseString, out *DatabaseResponse) error
		FindManyByOriginalTitle(ctx context.Context, in *DatabaseStrings, out *DatabaseResponse) error
		FindOneByLocalTitle(ctx context.Context, in *DatabaseString, out *DatabaseResponse) error
		FindManyByLocalTitle(ctx context.Context, in *DatabaseStrings, out *DatabaseResponse) error
	}
	type MovieService struct {
		movieService
	}
	h := &movieServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MovieService{h}, opts...))
}

type movieServiceHandler struct {
	MovieServiceHandler
}

func (h *movieServiceHandler) Create(ctx context.Context, in *Movie, out *DatabaseResponse) error {
	return h.MovieServiceHandler.Create(ctx, in, out)
}

func (h *movieServiceHandler) Update(ctx context.Context, in *Movie, out *DatabaseResponse) error {
	return h.MovieServiceHandler.Update(ctx, in, out)
}

func (h *movieServiceHandler) SoftDelete(ctx context.Context, in *Movie, out *DatabaseResponse) error {
	return h.MovieServiceHandler.SoftDelete(ctx, in, out)
}

func (h *movieServiceHandler) ForceDelete(ctx context.Context, in *Movie, out *DatabaseResponse) error {
	return h.MovieServiceHandler.ForceDelete(ctx, in, out)
}

func (h *movieServiceHandler) FindAll(ctx context.Context, in *DatabaseEmpty, out *DatabaseResponse) error {
	return h.MovieServiceHandler.FindAll(ctx, in, out)
}

func (h *movieServiceHandler) FindOnlyIDs(ctx context.Context, in *DatabaseEmpty, out *DatabaseResponse) error {
	return h.MovieServiceHandler.FindOnlyIDs(ctx, in, out)
}

func (h *movieServiceHandler) FindDownloaded(ctx context.Context, in *DatabaseEmpty, out *DatabaseResponse) error {
	return h.MovieServiceHandler.FindDownloaded(ctx, in, out)
}

func (h *movieServiceHandler) FindOneByID(ctx context.Context, in *DatabaseID, out *DatabaseResponse) error {
	return h.MovieServiceHandler.FindOneByID(ctx, in, out)
}

func (h *movieServiceHandler) FindManyByID(ctx context.Context, in *DatabaseIDs, out *DatabaseResponse) error {
	return h.MovieServiceHandler.FindManyByID(ctx, in, out)
}

func (h *movieServiceHandler) FindOneByTitle(ctx context.Context, in *DatabaseString, out *DatabaseResponse) error {
	return h.MovieServiceHandler.FindOneByTitle(ctx, in, out)
}

func (h *movieServiceHandler) FindManyByTitle(ctx context.Context, in *DatabaseStrings, out *DatabaseResponse) error {
	return h.MovieServiceHandler.FindManyByTitle(ctx, in, out)
}

func (h *movieServiceHandler) FindOneByOriginalTitle(ctx context.Context, in *DatabaseString, out *DatabaseResponse) error {
	return h.MovieServiceHandler.FindOneByOriginalTitle(ctx, in, out)
}

func (h *movieServiceHandler) FindManyByOriginalTitle(ctx context.Context, in *DatabaseStrings, out *DatabaseResponse) error {
	return h.MovieServiceHandler.FindManyByOriginalTitle(ctx, in, out)
}

func (h *movieServiceHandler) FindOneByLocalTitle(ctx context.Context, in *DatabaseString, out *DatabaseResponse) error {
	return h.MovieServiceHandler.FindOneByLocalTitle(ctx, in, out)
}

func (h *movieServiceHandler) FindManyByLocalTitle(ctx context.Context, in *DatabaseStrings, out *DatabaseResponse) error {
	return h.MovieServiceHandler.FindManyByLocalTitle(ctx, in, out)
}
