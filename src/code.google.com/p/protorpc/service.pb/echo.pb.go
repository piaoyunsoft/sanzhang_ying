// Code generated by protoc-gen-go.
// source: echo.proto
// DO NOT EDIT!

package service

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

import "io"
import "log"
import "net"
import "net/rpc"
import "time"
import protorpc "code.google.com/p/protorpc"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type EchoRequest struct {
	Msg              *string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}

func (m *EchoRequest) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

type EchoResponse struct {
	Msg              *string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}

func (m *EchoResponse) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

func init() {
}

type EchoService interface {
	Echo(in *EchoRequest, out *EchoResponse) error
	EchoTwice(in *EchoRequest, out *EchoResponse) error
}

// AcceptEchoServiceClient accepts connections on the listener and serves requests
// for each incoming connection.  Accept blocks; the caller typically
// invokes it in a go statement.
func AcceptEchoServiceClient(lis net.Listener, x EchoService) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("EchoService", x); err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeCodec(protorpc.NewServerCodec(conn))
	}
}

// RegisterEchoService publish the given EchoService implementation on the server.
func RegisterEchoService(srv *rpc.Server, x EchoService) error {
	if err := srv.RegisterName("EchoService", x); err != nil {
		return err
	}
	return nil
}

// NewEchoServiceServer returns a new EchoService Server.
func NewEchoServiceServer(x EchoService) *rpc.Server {
	srv := rpc.NewServer()
	if err := srv.RegisterName("EchoService", x); err != nil {
		log.Fatal(err)
	}
	return srv
}

// ListenAndServeEchoService listen announces on the local network address laddr
// and serves the given EchoService implementation.
func ListenAndServeEchoService(network, addr string, x EchoService) error {
	lis, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	defer lis.Close()

	srv := rpc.NewServer()
	if err := srv.RegisterName("EchoService", x); err != nil {
		return err
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeCodec(protorpc.NewServerCodec(conn))
	}
}

type EchoServiceClient struct {
	*rpc.Client
}

// NewEchoServiceClient returns a EchoService rpc.Client and stub to handle
// requests to the set of EchoService at the other end of the connection.
func NewEchoServiceClient(conn io.ReadWriteCloser) (*EchoServiceClient, *rpc.Client) {
	c := rpc.NewClientWithCodec(protorpc.NewClientCodec(conn))
	return &EchoServiceClient{c}, c
}

func (c *EchoServiceClient) Echo(in *EchoRequest, out *EchoResponse) error {
	return c.Call("EchoService.Echo", in, out)
}
func (c *EchoServiceClient) EchoTwice(in *EchoRequest, out *EchoResponse) error {
	return c.Call("EchoService.EchoTwice", in, out)
}

// DialEchoService connects to an EchoService at the specified network address.
func DialEchoService(network, addr string) (*EchoServiceClient, *rpc.Client, error) {
	c, err := protorpc.Dial(network, addr)
	if err != nil {
		return nil, nil, err
	}
	return &EchoServiceClient{c}, c, nil
}

// DialEchoServiceTimeout connects to an EchoService at the specified network address.
func DialEchoServiceTimeout(network, addr string,
	timeout time.Duration) (*EchoServiceClient, *rpc.Client, error) {
	c, err := protorpc.DialTimeout(network, addr, timeout)
	if err != nil {
		return nil, nil, err
	}
	return &EchoServiceClient{c}, c, nil
}
