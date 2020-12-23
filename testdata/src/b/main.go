package main

import (
	"context" // no rules

	"github.com/goccy/go-yaml" // no rules

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry" // valid alias

	"github.com/pkg/errors"           // no alias, invalid
	pkgerr "github.com/pkg/errors"    // valid alias
	pkgerrors "github.com/pkg/errors" // invalid alis

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"          // no aliases
	ddtracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer" // use alias, invalid
)

func main() {
	var (
		_ context.Context
		_ yaml.BytesMarshaler
		_ tracer.Span
		_ ddtracer.Span
		_ grpc_retry.BackoffFunc
	)
	_ = errors.New("error")
	_ = pkgerr.New("error")
	_ = pkgerrors.New("error")
}
