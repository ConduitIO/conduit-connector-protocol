// Copyright © 2021 Meroxa Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"context"
	"errors"
	"io"
	"sync"

	"github.com/conduitio/conduit-plugin/cpluginv1"
	"github.com/conduitio/conduit-plugin/cpluginv1/internal/fromproto"
	"github.com/conduitio/conduit-plugin/cpluginv1/internal/toproto"
	connectorv1 "github.com/conduitio/conduit-plugin/internal/connector/v1"
	"github.com/hashicorp/go-plugin"
	"go.uber.org/multierr"
	"google.golang.org/grpc"
)

type grpcSourcePlugin struct {
	plugin.NetRPCUnsupportedPlugin
}

var _ plugin.Plugin = (*grpcSourcePlugin)(nil)

func (p *grpcSourcePlugin) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, cc *grpc.ClientConn) (interface{}, error) {
	return &sourcePlugin{client: connectorv1.NewSourcePluginClient(cc)}, nil
}

// GRPCServer always returns an error; we're only implementing the client half
// of the interface.
func (p *grpcSourcePlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	return errors.New("this package only implements gRPC clients")
}

func NewSourcePlugin(client connectorv1.SourcePluginClient) cpluginv1.SourcePlugin {
	return &sourcePlugin{client: client}
}

type sourcePlugin struct {
	client connectorv1.SourcePluginClient
}

func (s *sourcePlugin) Configure(ctx context.Context, goReq cpluginv1.SourceConfigureRequest) (cpluginv1.SourceConfigureResponse, error) {
	protoReq, err := toproto.SourceConfigureRequest(goReq)
	if err != nil {
		return cpluginv1.SourceConfigureResponse{}, err
	}
	protoResp, err := s.client.Configure(ctx, protoReq)
	if err != nil {
		return cpluginv1.SourceConfigureResponse{}, err
	}
	goResp, err := fromproto.SourceConfigureResponse(protoResp)
	if err != nil {
		return cpluginv1.SourceConfigureResponse{}, err
	}
	return goResp, nil
}

func (s *sourcePlugin) Start(ctx context.Context, goReq cpluginv1.SourceStartRequest) (cpluginv1.SourceStartResponse, error) {
	protoReq, err := toproto.SourceStartRequest(goReq)
	if err != nil {
		return cpluginv1.SourceStartResponse{}, err
	}
	protoResp, err := s.client.Start(ctx, protoReq)
	if err != nil {
		return cpluginv1.SourceStartResponse{}, err
	}
	goResp, err := fromproto.SourceStartResponse(protoResp)
	if err != nil {
		return cpluginv1.SourceStartResponse{}, err
	}
	return goResp, nil
}

func (s *sourcePlugin) Stop(ctx context.Context, goReq cpluginv1.SourceStopRequest) (cpluginv1.SourceStopResponse, error) {
	protoReq, err := toproto.SourceStopRequest(goReq)
	if err != nil {
		return cpluginv1.SourceStopResponse{}, err
	}
	protoResp, err := s.client.Stop(ctx, protoReq)
	if err != nil {
		return cpluginv1.SourceStopResponse{}, err
	}
	goResp, err := fromproto.SourceStopResponse(protoResp)
	if err != nil {
		return cpluginv1.SourceStopResponse{}, err
	}
	return goResp, nil
}

func (s *sourcePlugin) Run(ctx context.Context, stream cpluginv1.SourceRunStream) error {
	runClient, err := s.client.Run(ctx)
	if err != nil {
		return err
	}

	errOut := make(chan error)

	var wg sync.WaitGroup
	wg.Add(2)

	// goroutine for receiving messages from the plugin
	go func() {
		defer wg.Done()
		protoResp, err := runClient.Recv()
		if err != nil {
			// TODO if we receive an error and the other goroutine is still
			//  running we won't notice it and run forever. This could happen if
			//  the plugin is not written correctly and its side of the stream.
			if errors.Is(err, io.EOF) {
				return // no error
			}
			errOut <- err
			return
		}
		goResp, err := fromproto.SourceRunResponse(protoResp)
		if err != nil {
			errOut <- err
			return
		}
		err = stream.Send(goResp)
		if err != nil {
			errOut <- err
			return
		}
	}()

	// goroutine for sending messages to the plugin
	go func() {
		defer wg.Done()
		defer func() {
			err := runClient.CloseSend()
			if err != nil {
				errOut <- err
			}
		}()
		for {
			goReq, err := stream.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return // no error
				}
				errOut <- err
				return
			}
			protoReq, err := toproto.SourceRunRequest(goReq)
			if err != nil {
				errOut <- err
				return
			}
			err = runClient.Send(protoReq)
			if err != nil {
				errOut <- err
				return
			}
		}
	}()

	// this goroutine ensures errOut will be closed when both goroutines stop
	go func() {
		wg.Wait()
		close(errOut)
	}()

	// capture errors coming from goroutines
	for e := range errOut {
		err = multierr.Append(err, e)
	}

	return err
}
