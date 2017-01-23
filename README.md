[![Build Status](https://travis-ci.org/appdataspec/sdk-golang.svg?branch=master)](https://travis-ci.org/appdataspec/sdk-golang)[![Coverage](https://codecov.io/gh/appdataspec/sdk-golang/branch/master/graph/badge.svg)](https://codecov.io/gh/appdataspec/sdk-golang)

Golang SDK for the [app data spec](https://github.com/appdataspec/spec)

> *Be advised: this project is currently at Major version zero. Per the
> semantic versioning spec: "Major version zero (0.y.z) is for initial
> development. Anything may change at any time. The public API should
> not be considered stable."*

# Packages

[/pkg](pkg) contains all packages distributed by this SDK.

# Usage

```go
package myDummyPackage

import (
"github.com/appdataspec/sdk-golang/pkg/path"
"fmt"
)

func main() {
// use path package for working w/ ADS (app data spec) paths
adsPath := path.New()

fmt.Printf("Global ADS path is: %v\n", adsPath.Global())
fmt.Printf("Per user ADS path is: %v\n", adsPath.PerUser())
}
```

# Releases

All releases will be
[tagged](https://github.com/opspec-io/sdk-golang/tags) and made
available on the
[releases](https://github.com/opspec-io/sdk-golang/releases) page with
links to docs.

# Versioning

This project adheres to the [Semantic Versioning](http://semver.org/)
specification

# Contributing

see [CONTRIBUTING.md](CONTRIBUTING.md)

# Changelog

see [CHANGELOG.md](CHANGELOG.md)
