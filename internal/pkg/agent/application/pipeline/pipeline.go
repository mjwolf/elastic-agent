// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package pipeline

import (
	"github.com/elastic/elastic-agent/internal/pkg/agent/application/info"
	"github.com/elastic/elastic-agent/internal/pkg/agent/configrequest"
	"github.com/elastic/elastic-agent/internal/pkg/agent/program"
	"github.com/elastic/elastic-agent/internal/pkg/agent/storage/store"
	"github.com/elastic/elastic-agent/internal/pkg/agent/transpiler"
	"github.com/elastic/elastic-agent/internal/pkg/config"
	"github.com/elastic/elastic-agent/internal/pkg/core/logger"
	"github.com/elastic/elastic-agent/internal/pkg/fleetapi"
	"github.com/elastic/elastic-agent/internal/pkg/sorted"
)

// ConfigHandler is capable of handling configrequest.
type ConfigHandler interface {
	HandleConfig(configrequest.Request) error
	Close() error
	Shutdown()
}

// DefaultRK default routing keys until we implement the routing key / config matrix.
var DefaultRK = "default"

// RoutingKey is used for routing as pipeline id.
type RoutingKey = string

// Router is an interface routing programs to the corresponding stream.
type Router interface {
	Routes() *sorted.Set
	Route(id string, grpProg map[RoutingKey][]program.Program) error
	Shutdown()
}

// StreamFunc creates a stream out of routing key.
type StreamFunc func(*logger.Logger, RoutingKey) (Stream, error)

// Stream is capable of executing configrequest change.
type Stream interface {
	Execute(configrequest.Request) error
	Close() error
	Shutdown()
}

// EmitterFunc emits configuration for processing.
type EmitterFunc func(*config.Config) error

// DecoratorFunc is a func for decorating a retrieved configuration before processing.
type DecoratorFunc = func(*info.AgentInfo, string, *transpiler.AST, []program.Program) ([]program.Program, error)

// FilterFunc is a func for filtering a retrieved configuration before processing.
type FilterFunc = func(*logger.Logger, *transpiler.AST) error

// ConfigModifiers is a collections of filters and decorators applied while processing configuration.
type ConfigModifiers struct {
	Filters    []FilterFunc
	Decorators []DecoratorFunc
}

// Dispatcher processes actions coming from fleet api.
type Dispatcher interface {
	Dispatch(acker store.FleetAcker, actions ...fleetapi.Action) error
}
