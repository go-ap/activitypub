package build

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	vocab "github.com/go-ap/activitypub"
)

func Actor(initFn ...InitFn) *vocab.Actor {
	act := new(vocab.Actor)
	for _, maybeFn := range initFn {
		switch fn := maybeFn.(type) {
		case vocab.WithObjectFn:
			_ = vocab.OnObject(act, fn)
		case vocab.WithActorFn:
			_ = vocab.OnActor(act, fn)
		}
	}
	return act
}

func Person(initFn ...InitFn) *vocab.Actor {
	initFn = append([]InitFn{WithType(vocab.PersonType)}, initFn...)
	return Actor(initFn...)
}

func Group(initFn ...InitFn) *vocab.Actor {
	initFn = append([]InitFn{WithType(vocab.GroupType)}, initFn...)
	return Actor(initFn...)
}

func Application(initFn ...InitFn) *vocab.Actor {
	initFn = append([]InitFn{WithType(vocab.ApplicationType)}, initFn...)
	return Actor(initFn...)
}

func Service(initFn ...InitFn) *vocab.Actor {
	initFn = append([]InitFn{WithType(vocab.ServiceType)}, initFn...)
	return Actor(initFn...)
}

func WithPreferredUsername[T ~string | vocab.NaturalLanguageValues](c T) vocab.WithActorFn {
	v := nlv(c)
	return func(act *vocab.Actor) error {
		act.PreferredUsername = v
		return nil
	}
}

func WithInbox(i ...vocab.Item) vocab.WithActorFn {
	return func(act *vocab.Actor) error {
		if i == nil && !vocab.EmptyID.Equal(act.ID) {
			i = vocab.ItemCollection{vocab.Inbox.IRI(act)}
		}
		act.Inbox = vocab.ItemCollection(i).Normalize()
		return nil
	}
}

func WithOutbox(i ...vocab.Item) vocab.WithActorFn {
	return func(act *vocab.Actor) error {
		if i == nil && !vocab.EmptyID.Equal(act.ID) {
			i = vocab.ItemCollection{vocab.Outbox.IRI(act)}
		}
		act.Outbox = vocab.ItemCollection(i).Normalize()
		return nil
	}
}

func WithLiked(l ...vocab.Item) vocab.WithActorFn {
	return func(act *vocab.Actor) error {
		if l == nil && !vocab.EmptyID.Equal(act.ID) {
			l = vocab.ItemCollection{vocab.Liked.IRI(act)}
		}
		act.Liked = vocab.ItemCollection(l).Normalize()
		return nil
	}
}

func WithFollowing(f ...vocab.Item) vocab.WithActorFn {
	return func(act *vocab.Actor) error {
		if f == nil && !vocab.EmptyID.Equal(act.ID) {
			f = vocab.ItemCollection{vocab.Following.IRI(act)}
		}
		act.Following = vocab.ItemCollection(f).Normalize()
		return nil
	}
}

func WithFollowers(f ...vocab.Item) vocab.WithActorFn {
	return func(act *vocab.Actor) error {
		if f == nil && !vocab.EmptyID.Equal(act.ID) {
			f = vocab.ItemCollection{vocab.Followers.IRI(act)}
		}
		act.Followers = vocab.ItemCollection(f).Normalize()
		return nil
	}
}

func WithAuthEndpoint(i vocab.IRI) vocab.WithActorFn {
	return func(act *vocab.Actor) error {
		if act.Endpoints == nil {
			act.Endpoints = new(vocab.Endpoints)
		}
		act.Endpoints.OauthAuthorizationEndpoint = i
		return nil
	}
}

func WithTokenEndpoint(i vocab.IRI) vocab.WithActorFn {
	return func(act *vocab.Actor) error {
		if act.Endpoints == nil {
			act.Endpoints = new(vocab.Endpoints)
		}
		act.Endpoints.OauthTokenEndpoint = i
		return nil
	}
}

func WithProxyURL(i vocab.IRI) vocab.WithActorFn {
	return func(act *vocab.Actor) error {
		if act.Endpoints == nil {
			act.Endpoints = new(vocab.Endpoints)
		}
		act.Endpoints.ProxyURL = i
		return nil
	}
}

func WithSharedInbox(i vocab.IRI) vocab.WithActorFn {
	return func(act *vocab.Actor) error {
		if act.Endpoints == nil {
			act.Endpoints = new(vocab.Endpoints)
		}
		act.Endpoints.SharedInbox = i
		return nil
	}
}

func WithStreams(i ...vocab.Item) vocab.WithActorFn {
	s := vocab.ItemCollection(i)
	return func(act *vocab.Actor) error {
		return act.Streams.Append(s...)
	}
}

func WithPublicKey(k crypto.PublicKey) vocab.WithActorFn {
	pubEnc, _ := x509.MarshalPKIXPublicKey(k)
	pubEncoded := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pubEnc})
	return func(act *vocab.Actor) error {
		act.PublicKey = vocab.PublicKey{
			ID:           vocab.IRI(fmt.Sprintf("%s#main", act.ID)),
			Owner:        act.ID,
			PublicKeyPem: string(pubEncoded),
		}
		return nil
	}
}
