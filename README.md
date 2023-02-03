# Activity Pub for Go

[![MIT Licensed](https://img.shields.io/github/license/go-ap/activitypub.svg)](https://raw.githubusercontent.com/go-ap/activitypub/master/LICENSE)
[![Build Status](https://builds.sr.ht/~mariusor/activitypub.svg)](https://builds.sr.ht/~mariusor/activitypub)
[![Test Coverage](https://img.shields.io/codecov/c/github/go-ap/activitypub.svg)](https://codecov.io/gh/go-ap/activitypub)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-ap/activitypub)](https://goreportcard.com/report/github.com/go-ap/activitypub)

Basic package for using [ActivityPub](https://www.w3.org/TR/activitypub/#Overview) API in Go.

It contains types for most of the ActivityStreams vocabulary and the ActivityPub extension.
They are documented accordingly with annotations from the specification.

## Usage

```go
import "github.com/go-ap/activitypub"
```

## Note about generics

The helper functions exposed by the package come in two flavours: 
explicit `OnXXX` and `ToXXX` functions corresponding to each type and,
a generic pair of functions `On[T]` and `To[T]`.

Before using them you should consider that the latter comes with a pretty heavy performance penalty:

```
goos: linux
goarch: amd64
pkg: github.com/go-ap/activitypub
cpu: Intel(R) Core(TM) i7-6700K CPU @ 4.00GHz
Benchmark_OnT_vs_On_T/OnObject-8                    752387791       1.633 ns/op
Benchmark_OnT_vs_On_T/On_T_Object-8                   4656264     261.8   ns/op
Benchmark_OnT_vs_On_T/OnActor-8                     739833261       1.596 ns/op
Benchmark_OnT_vs_On_T/On_T_Actor-8                    4035148     301.9   ns/op
Benchmark_OnT_vs_On_T/OnActivity-8                  751173854       1.604 ns/op
Benchmark_OnT_vs_On_T/On_T_Activity-8                 4062598     285.9   ns/op
Benchmark_OnT_vs_On_T/OnIntransitiveActivity-8      675824500       1.640 ns/op
Benchmark_OnT_vs_On_T/On_T_IntransitiveActivity-8     4372798     274.1   ns/op
PASS
ok  	github.com/go-ap/activitypub	11.350s
```
