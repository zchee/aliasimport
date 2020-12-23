package a

import (
	"context" // no rules

	"github.com/goccy/go-yaml" // no rules

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry" // valid alias

	//lint:ignore aliasimport reason
	perr "github.com/pkg/errors"
	perrors "github.com/pkg/errors" //lint:ignore aliasimport reason

	"github.com/pkg/errors"           // want `the package "github.com/pkg/errors" should be imported with the alias name pkgerr`
	_ "github.com/pkg/errors"         // want `the alias name of "github.com/pkg/errors" should be pkgerr, not _`
	pkgerr "github.com/pkg/errors"    // valid alias
	pkgerrors "github.com/pkg/errors" // want `the alias name of "github.com/pkg/errors" should be pkgerr, not pkgerrors`

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"          // no aliases
	ddtracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer" // want `the package "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer" shouldn't be imported with any aliases, but with ddtracer`
)

func f() {
	var (
		_ context.Context
		_ yaml.BytesMarshaler
		_ tracer.Span
		_ ddtracer.Span
		_ grpc_retry.BackoffFunc
	)
	_ = errors.New("error")
	_ = perrors.New("error")
	_ = perr.New("error")
	_ = pkgerr.New("error")
	_ = pkgerrors.New("error")
}
