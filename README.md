# About GoActivityPub: Vocabulary

[![MIT Licensed](https://img.shields.io/github/license/go-ap/activitypub.svg)](https://raw.githubusercontent.com/go-ap/activitypub/master/LICENSE)
[![Build Status](https://builds.sr.ht/~mariusor/activitypub.svg)](https://builds.sr.ht/~mariusor/activitypub)
[![Test coverage](https://img.shields.io/coverallsCoverage/github/go-ap/activitypub)](https://coveralls.io/github/go-ap/activitypub)

This project is part of the [GoActivityPub](https://github.com/go-ap) library which helps with creating ActivityPub applications using the Go programming language.

It contains data types for most of the [Activity Vocabulary](https://www.w3.org/TR/activitystreams-vocabulary/) and the [ActivityPub](https://www.w3.org/TR/activitypub/) extension.
They are documented accordingly with annotations from these specifications.

<!-- Expanded [documentation](https://man.sr.ht/~mariusor/go-activitypub/lib/index.md) -->
For more details you can have a look at [the expanded documentation](https://go-activitypub.federated.id/lib).

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

There are more examples in the [official documentation](https://pkg.go.dev/github.com/go-ap/activitypub#pkg-examples).