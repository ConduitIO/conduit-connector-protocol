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

type grpcDestinationPlugin struct {
	plugin.NetRPCUnsupportedPlugin
}

var _ plugin.Plugin = (*grpcDestinationPlugin)(nil)

func (p *grpcDestinationPlugin) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, cc *grpc.ClientConn) (interface{}, error) {
	return &destinationPlugin{client: connectorv1.NewDestinationPluginClient(cc)}, nil
}

// GRPCServer always returns an error; we're only implementing the client half
// of the interface.
func (p *grpcDestinationPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	return errors.New("this package only implements gRPC clients")
}

func NewDestinationPlugin(client connectorv1.DestinationPluginClient) cpluginv1.DestinationPlugin {
	return &destinationPlugin{client: client}
}

type destinationPlugin struct {
	client connectorv1.DestinationPluginClient
}

func (s *destinationPlugin) Configure(ctx context.Context, goReq cpluginv1.DestinationConfigureRequest) (cpluginv1.DestinationConfigureResponse, error) {
	protoReq, err := toproto.DestinationConfigureRequest(goReq)
	if err != nil {
		return cpluginv1.DestinationConfigureResponse{}, err
	}
	protoResp, err := s.client.Configure(ctx, protoReq)
	if err != nil {
		return cpluginv1.DestinationConfigureResponse{}, err
	}
	goResp, err := fromproto.DestinationConfigureResponse(protoResp)
	if err != nil {
		return cpluginv1.DestinationConfigureResponse{}, err
	}
	return goResp, nil
}

func (s *destinationPlugin) Start(ctx context.Context, goReq cpluginv1.DestinationStartRequest) (cpluginv1.DestinationStartResponse, error) {
	protoReq, err := toproto.DestinationStartRequest(goReq)
	if err != nil {
		return cpluginv1.DestinationStartResponse{}, err
	}
	protoResp, err := s.client.Start(ctx, protoReq)
	if err != nil {
		return cpluginv1.DestinationStartResponse{}, err
	}
	goResp, err := fromproto.DestinationStartResponse(protoResp)
	if err != nil {
		return cpluginv1.DestinationStartResponse{}, err
	}
	return goResp, nil
}

func (s *destinationPlugin) Stop(ctx context.Context, goReq cpluginv1.DestinationStopRequest) (cpluginv1.DestinationStopResponse, error) {
	protoReq, err := toproto.DestinationStopRequest(goReq)
	if err != nil {
		return cpluginv1.DestinationStopResponse{}, err
	}
	protoResp, err := s.client.Stop(ctx, protoReq)
	if err != nil {
		return cpluginv1.DestinationStopResponse{}, err
	}
	goResp, err := fromproto.DestinationStopResponse(protoResp)
	if err != nil {
		return cpluginv1.DestinationStopResponse{}, err
	}
	return goResp, nil
}

func (s *destinationPlugin) Run(ctx context.Context, stream cpluginv1.DestinationRunStream) error {
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
		goResp, err := fromproto.DestinationRunResponse(protoResp)
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
			protoReq, err := toproto.DestinationRunRequest(goReq)
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
