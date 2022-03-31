// Code generated by protoc-gen-go-micro. DO NOT EDIT.
// versions:
// - protoc-gen-go-micro 23815987fd5c65de1306066d8380981c5ec77123
// - protoc              v3.19.4
// source: agent/proto/bot.proto

package proto

import (
	api "c-z.dev/go-micro/api"
	client "c-z.dev/go-micro/client"
	server "c-z.dev/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// NewCommandEndpoints API Endpoints for Command service
func NewCommandEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// CommandService is the client API for Command service.
type CommandService interface {
	Help(ctx context.Context, in *HelpRequest, opts ...client.CallOption) (*HelpResponse, error)
	Exec(ctx context.Context, in *ExecRequest, opts ...client.CallOption) (*ExecResponse, error)
}

type commandService struct {
	c    client.Client
	name string
}

func NewCommandService(name string, c client.Client) CommandService {
	return &commandService{
		c:    c,
		name: name,
	}
}

func (c *commandService) Help(ctx context.Context, in *HelpRequest, opts ...client.CallOption) (*HelpResponse, error) {
	req := c.c.NewRequest(c.name, "Command.Help", in)
	out := new(HelpResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandService) Exec(ctx context.Context, in *ExecRequest, opts ...client.CallOption) (*ExecResponse, error) {
	req := c.c.NewRequest(c.name, "Command.Exec", in)
	out := new(ExecResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandHandler is the server API for Command service.
type CommandHandler interface {
	Help(context.Context, *HelpRequest, *HelpResponse) error
	Exec(context.Context, *ExecRequest, *ExecResponse) error
}

func RegisterCommandHandler(s server.Server, hdlr CommandHandler, opts ...server.HandlerOption) error {
	type command interface {
		Help(ctx context.Context, in *HelpRequest, out *HelpResponse) error
		Exec(ctx context.Context, in *ExecRequest, out *ExecResponse) error
	}
	type Command struct {
		command
	}
	h := &commandHandler{hdlr}
	return s.Handle(s.NewHandler(&Command{h}, opts...))
}

type commandHandler struct {
	CommandHandler
}

func (h *commandHandler) Help(ctx context.Context, in *HelpRequest, out *HelpResponse) error {
	return h.CommandHandler.Help(ctx, in, out)
}

func (h *commandHandler) Exec(ctx context.Context, in *ExecRequest, out *ExecResponse) error {
	return h.CommandHandler.Exec(ctx, in, out)
}
