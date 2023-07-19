# About GoActivityPub: Vocabulary

[![MIT Licensed](https://img.shields.io/github/license/go-ap/activitypub.svg)](https://raw.githubusercontent.com/go-ap/activitypub/master/LICENSE)
[![Build Status](https://builds.sr.ht/~mariusor/activitypub.svg)](https://builds.sr.ht/~mariusor/activitypub)
[![Test Coverage](https://img.shields.io/codecov/c/github/go-ap/activitypub.svg)](https://codecov.io/gh/go-ap/activitypub)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-ap/activitypub)](https://goreportcard.com/report/github.com/go-ap/activitypub)

This project is part of the [GoActivityPub](https://github.com/go-ap) library which helps with creating ActivityPub applications using the Go programming language.

It contains data types for most of the [Activity Vocabulary](https://www.w3.org/TR/activitystreams-vocabulary/) and the [ActivityPub](https://www.w3.org/TR/activitypub/) extension.
They are documented accordingly with annotations from these specifications.

You can find an expanded documentation about the whole library [on SourceHut](https://man.sr.ht/~mariusor/go-activitypub/go-ap/index.md).

For discussions about the projects you can write to the discussions mailing list: [~mariusor/go-activitypub-discuss@lists.sr.ht](mailto:~mariusor/go-activitypub-discuss@lists.sr.ht)

For patches and bug reports please use the dev mailing list: [~mariusor/go-activitypub-dev@lists.sr.ht](mailto:~mariusor/go-activitypub-dev@lists.sr.ht)

## Usage

```go
import vocab "github.com/go-ap/activitypub"

follow := vocab.Activity{
    Type: vocab.FollowType,
    Actor: vocab.IRI("https://example.com/alice"),
    Object: vocab.IRI("https://example.com/janedoe"),
}

```

## Note about generics

The module contains helper functions which make it simpler to deal with the `vocab.Item` 
interfaces and they come in two flavours: explicit `OnXXX` and `ToXXX` functions corresponding 
to each type and, a generic pair of functions `On[T]` and `To[T]`.

```go
import (
    "fmt"

    vocab "github.com/go-ap/activitypub"
)

var it vocab.Item = ... // an ActivityPub object unmarshaled from a request

err := vocab.OnActivity(it, func(act *vocab.Activity) error {
    if vocab.ContentManagementActivityTypes.Contains(act.Type) {
        fmt.Printf("This is a Content Management type activity: %q", act.Type)
    }
    return nil
})

err := vocab.On[vocab.Activity](it, func(act *vocab.Activity) error {
    if vocab.ReactionsActivityTypes.Contains(act.Type) {
        fmt.Printf("This is a Reaction type activity: %q", act.Type)
    }
    return nil
})

```

Before using the generic versions you should consider that they come with a pretty heavy performance penalty:

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
