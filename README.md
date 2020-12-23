# aliasimport

[![pkg.go.dev][pkg.go.dev-badge]][pkg.go.dev]
[![Test][test-badge]][test]
[![reviewdog][reviewdog-badge]][reviewdog]
[![Releases][release-badge]][release]
[![codecov][codecov-badge]][codecov]

`aliasimport` defines some rules about import statement alias.

## Define rules via YAML

You can define two type rules as `aliases` and `noaliases`, looks like the following code.

```yaml:rules.yml
---
aliases:
  grpc_retry: github.com/grpc-ecosystem/go-grpc-middleware/retry
  pkgerr: "github.com/pkg/errors"  # quoted
noaliases:
  - gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer
```

- `aliases`
  - The list of definitions about alias naming
  - Only one alias can be defined for each package.
- `noaliases`
  - The list of package names that don't need any aliases
  - These packages should be imported nakedly.


```go
package main

import (
	"context" // no rules

	"github.com/goccy/go-yaml" // no rules

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry" // valid alias

	"github.com/pkg/errors"           // no alias, invalid
	pkgerr "github.com/pkg/errors"    // valid alias
	pkgerrors "github.com/pkg/errors" // invalid alias

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
```

```sh
go vet -vettool=`which aliasimport` -aliasimport.rule=$(pwd)/rules.yml main.go
./main.go:10:2: the package "github.com/pkg/errors" should be imported with the alias name pkgerr
./main.go:12:2: the alias name of "github.com/pkg/errors" should be pkgerr, not pkgerrors
./main.go:15:2: the package "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer" shouldn't be imported with any aliases, but with ddtracer
```

## False Positives

Analyzers ignore nodes which are annotated by [staticcheck's style comments](https://staticcheck.io/docs/#ignoring-problems) as belows.
A ignore comment includes analyzer names and reason of ignoring checking.
If you specify `aliasimport` as analyzer name, all analyzers ignore corresponding code.

```go
import (
	//lint:ignore aliasimport reason
	"github.com/pkg/errors"

	pkgerrors "github.com/pkg/errors" //lint:ignore aliasimport reason
)
```

## Commiters

- Motonori IWATA([@iwata](https://github.com/iwata))

## License

Licensed under the MIT License.

<!-- badge links -->
[test]: https://github.com/iwata/aliasimport/actions?query=workflow%3A%22test+and+coverage%22
[reviewdog]: https://github.com/iwata/aliasimport/actions?query=workflow%3Areviewdog
[codecov]: https://codecov.io/gh/iwata/aliasimport
[pkg.go.dev]: https://pkg.go.dev/github.com/iwata/aliasimport
[release]: https://github.com/iwata/aliasimport/releases/latest

[test-badge]: https://github.com/iwata/aliasimport/workflows/test%20and%20coverage/badge.svg
[reviewdog-badge]: https://github.com/iwata/aliasimport/workflows/reviewdog/badge.svg
[codecov-badge]: https://codecov.io/gh/iwata/aliasimport/branch/main/graph/badge.svg?token=yOEEZZqDse
[pkg.go.dev-badge]: https://pkg.go.dev/badge/github.com/iwata/aliasimport.svg
[release-badge]: https://img.shields.io/github/release/iwata/aliasimport.svg?style=flat&logo=github
