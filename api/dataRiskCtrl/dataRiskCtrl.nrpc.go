// This code was autogenerated from dataRiskCtrl/dataRiskCtrl.proto, do not edit.
package dataRiskCtrl

import (
	"context"
	"log"
	"time"

	"google.golang.org/protobuf/proto"
	"github.com/nats-io/nats.go"
	github_com_golang_protobuf_ptypes_empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/nats-rpc/nrpc"
)

// DataRiskCtrlServer is the interface that providers of the service
// DataRiskCtrl should implement.
type DataRiskCtrlServer interface {
	RpcAlive(ctx context.Context, req *github_com_golang_protobuf_ptypes_empty.Empty) (*github_com_golang_protobuf_ptypes_empty.Empty, error)
	QueryCnt(ctx context.Context, req *QueryReq) (*QueryRes, error)
}

// DataRiskCtrlHandler provides a NATS subscription handler that can serve a
// subscription using a given DataRiskCtrlServer implementation.
type DataRiskCtrlHandler struct {
	ctx     context.Context
	workers *nrpc.WorkerPool
	nc      nrpc.NatsConn
	server  DataRiskCtrlServer

	encodings []string
}

func NewDataRiskCtrlHandler(ctx context.Context, nc nrpc.NatsConn, s DataRiskCtrlServer) *DataRiskCtrlHandler {
	return &DataRiskCtrlHandler{
		ctx:    ctx,
		nc:     nc,
		server: s,

		encodings: []string{"protobuf"},
	}
}

func NewDataRiskCtrlConcurrentHandler(workers *nrpc.WorkerPool, nc nrpc.NatsConn, s DataRiskCtrlServer) *DataRiskCtrlHandler {
	return &DataRiskCtrlHandler{
		workers: workers,
		nc:      nc,
		server:  s,
	}
}

// SetEncodings sets the output encodings when using a '*Publish' function
func (h *DataRiskCtrlHandler) SetEncodings(encodings []string) {
	h.encodings = encodings
}

func (h *DataRiskCtrlHandler) Subject() string {
	return "DataRiskCtrl.>"
}

func (h *DataRiskCtrlHandler) Handler(msg *nats.Msg) {
	var ctx context.Context
	if h.workers != nil {
		ctx = h.workers.Context
	} else {
		ctx = h.ctx
	}
	request := nrpc.NewRequest(ctx, h.nc, msg.Subject, msg.Reply)
	// extract method name & encoding from subject
	_, _, name, tail, err := nrpc.ParseSubject(
		"", 0, "DataRiskCtrl", 0, msg.Subject)
	if err != nil {
		log.Printf("DataRiskCtrlHanlder: DataRiskCtrl subject parsing failed: %v", err)
		return
	}

	request.MethodName = name
	request.SubjectTail = tail

	// call handler and form response
	var immediateError *nrpc.Error
	switch name {
	case "RpcAlive":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("RpcAliveHanlder: RpcAlive subject parsing failed: %v", err)
			break
		}
		var req github_com_golang_protobuf_ptypes_empty.Empty
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("RpcAliveHandler: RpcAlive request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.RpcAlive(ctx, &req)
				if err != nil {
					return nil, err
				}
				return innerResp, err
			}
		}
	case "QueryCnt":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("QueryCntHanlder: QueryCnt subject parsing failed: %v", err)
			break
		}
		var req QueryReq
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("QueryCntHandler: QueryCnt request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.QueryCnt(ctx, &req)
				if err != nil {
					return nil, err
				}
				return innerResp, err
			}
		}
	default:
		log.Printf("DataRiskCtrlHandler: unknown name %q", name)
		immediateError = &nrpc.Error{
			Type: nrpc.Error_CLIENT,
			Message: "unknown name: " + name,
		}
	}
	if immediateError == nil {
		if h.workers != nil {
			// Try queuing the request
			if err := h.workers.QueueRequest(request); err != nil {
				log.Printf("nrpc: Error queuing the request: %s", err)
			}
		} else {
			// Run the handler synchronously
			request.RunAndReply()
		}
	}

	if immediateError != nil {
		if err := request.SendReply(nil, immediateError); err != nil {
			log.Printf("DataRiskCtrlHandler: DataRiskCtrl handler failed to publish the response: %s", err)
		}
	} else {
	}
}

type DataRiskCtrlClient struct {
	nc      nrpc.NatsConn
	Subject string
	Encoding string
	Timeout time.Duration
}

func NewDataRiskCtrlClient(nc nrpc.NatsConn) *DataRiskCtrlClient {
	return &DataRiskCtrlClient{
		nc:      nc,
		Subject: "DataRiskCtrl",
		Encoding: "protobuf",
		Timeout: 5 * time.Second,
	}
}

func (c *DataRiskCtrlClient) RpcAlive(req *github_com_golang_protobuf_ptypes_empty.Empty) (*github_com_golang_protobuf_ptypes_empty.Empty, error) {

	subject := c.Subject + "." + "RpcAlive"

	// call
	var resp = github_com_golang_protobuf_ptypes_empty.Empty{}
	if err := nrpc.Call(req, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *DataRiskCtrlClient) QueryCnt(req *QueryReq) (*QueryRes, error) {

	subject := c.Subject + "." + "QueryCnt"

	// call
	var resp = QueryRes{}
	if err := nrpc.Call(req, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
		return nil, err
	}

	return &resp, nil
}

type Client struct {
	nc      nrpc.NatsConn
	defaultEncoding string
	defaultTimeout time.Duration
	DataRiskCtrl *DataRiskCtrlClient
}

func NewClient(nc nrpc.NatsConn) *Client {
	c := Client{
		nc: nc,
		defaultEncoding: "protobuf",
		defaultTimeout: 5*time.Second,
	}
	c.DataRiskCtrl = NewDataRiskCtrlClient(nc)
	return &c
}

func (c *Client) SetEncoding(encoding string) {
	c.defaultEncoding = encoding
	if c.DataRiskCtrl != nil {
		c.DataRiskCtrl.Encoding = encoding
	}
}

func (c *Client) SetTimeout(t time.Duration) {
	c.defaultTimeout = t
	if c.DataRiskCtrl != nil {
		c.DataRiskCtrl.Timeout = t
	}
}