/*
Copyright 2018 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sdk

import (
	"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1"
	client "github.com/micro/go-micro/client"
	context "golang.org/x/net/context"
)

type key int

const (
	CloudEventsKey           key    = 0
	CloudEventsVersion       string = "v1.0"
	ContextExtensionErrorKey string = "error"
)

const (
	MetadataKey   string = "type"
	MetadataValue string = "signal"
)

// SignalMetadata is a map of metadata that must be included
// for micro services to be discoverable as signals.
var SignalMetadata = map[string]string{MetadataKey: MetadataValue}

// Listener handles signal lifecycle
// Must be implemented by all signal service providers
// TODO: this should be the interface implemented by all the signal sources to be a signal provider
type Listener interface {
	Listen(*v1alpha1.Signal, <-chan struct{}) (<-chan *v1alpha1.Event, error)
}

// SignalClient is the interface for signal clients
type SignalClient interface {
	Ping(context.Context) error
	Listen(context.Context, *v1alpha1.Signal, ...client.CallOption) (SignalService_ListenService, error)
	Handshake(*v1alpha1.Signal, SignalService_ListenService) error
}

// SignalServer is the interface for signal servers
type SignalServer interface {
	SignalServiceHandler
	handshake(SignalService_ListenStream, <-chan struct{}) (<-chan *v1alpha1.Event, error)
}
