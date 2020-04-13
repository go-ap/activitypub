package handlers

import (
	"fmt"
	pub "github.com/go-ap/activitypub"
	"net/http"
	"strings"
)

// CollectionType
type CollectionType string

// CollectionTypes
type CollectionTypes []CollectionType

const (
	Unknown   = CollectionType("")
	Outbox    = CollectionType("outbox")
	Inbox     = CollectionType("inbox")
	Shares    = CollectionType("shares")
	Replies   = CollectionType("replies") // activitystreams
	Following = CollectionType("following")
	Followers = CollectionType("followers")
	Liked     = CollectionType("liked")
	Likes     = CollectionType("likes")
)

// Typer is the static package variable that determines a collection type for a particular request
// It can be overloaded from outside packages.
// TODO(marius): This should be moved as a property on an instantiable package object, instead of keeping it here
var Typer CollectionTyper = pathTyper{}

// CollectionTyper allows external packages to tell us which collection the current HTTP request addresses
type CollectionTyper interface {
	Type(r *http.Request) CollectionType
}

type pathTyper struct{}

func (d pathTyper) Type(r *http.Request) CollectionType {
	if r.URL == nil || len(r.URL.Path) == 0 {
		return Unknown
	}
	var col string
	pathElements := strings.Split(r.URL.Path[1:], "/") // Skip first /
	for i := len(pathElements) - 1; i >= 0; i-- {
		col = pathElements[i]
		if typ := getValidActivityCollection(col); typ != Unknown {
			return typ
		}
		if typ := getValidObjectCollection(col); typ != Unknown {
			return typ
		}
	}

	return CollectionType(strings.ToLower(col))
}

var validActivityCollection = CollectionTypes{
	Outbox,
	Inbox,
	Likes,
	Shares,
	Replies, // activitystreams
}

var onObject = CollectionTypes{
	Likes,
	Shares,
	Replies,
}

var onActor = CollectionTypes{
	Outbox,
	Inbox,
	Liked,
	Following,
	Followers,
}

func (t CollectionTypes) Contains(typ CollectionType) bool {
	for _, tt := range t {
		if strings.ToLower(string(typ)) == string(tt) {
			return true
		}
	}
	return false
}

func (t CollectionType) IRI(i pub.Item) pub.IRI {
	var iri pub.IRI
	if i == nil {
		return pub.EmptyIRI
	}
	if i.IsObject() {
		if onActor.Contains(t) {
			pub.OnActor(i, func(a *pub.Actor) error {
				if t == Inbox && a.Inbox != nil {
					iri = a.Inbox.GetLink()
				}
				if t == Outbox && a.Outbox != nil {
					iri = a.Outbox.GetLink()
				}
				if t == Liked && a.Liked != nil {
					iri = a.Liked.GetLink()
				}
				if t == Following && a.Following != nil {
					iri = a.Following.GetLink()
				}
				if t == Followers && a.Followers != nil {
					iri = a.Followers.GetLink()
				}
				return nil
			})
		}
		if onObject.Contains(t) {
			pub.OnObject(i, func(o *pub.Object) error {
				if t == Likes && o.Likes != nil {
					iri = o.Likes.GetLink()
				}
				if t == Shares && o.Shares != nil {
					iri = o.Shares.GetLink()
				}
				if t == Replies && o.Replies != nil {
					iri = o.Replies.GetLink()
				}
				return nil
			})
		}
	}

	if len(iri) == 0 {
		iri = pub.IRI(fmt.Sprintf("%s/%s", i.GetLink(), t))
	}
	return iri
}

func getValidActivityCollection(typ string) CollectionType {
	t := CollectionType(typ)
	if validActivityCollection.Contains(t) {
		return t
	}
	return Unknown
}

// ValidActivityCollection shows if the current ActivityPub end-point type is a valid one for handling Activities
func ValidActivityCollection(typ string) bool {
	return getValidActivityCollection(typ) != Unknown
}

var validObjectCollection = []CollectionType{
	Following,
	Followers,
	Liked,
}

func getValidObjectCollection(typ string) CollectionType {
	for _, t := range validObjectCollection {
		if strings.ToLower(typ) == string(t) {
			return t
		}
	}
	return Unknown
}

// ValidActivityCollection shows if the current ActivityPub end-point type is a valid one for handling Objects
func ValidObjectCollection(typ string) bool {
	return getValidObjectCollection(typ) != Unknown
}

func getValidCollection(typ string) CollectionType {
	if typ := getValidActivityCollection(typ); typ != Unknown {
		return typ
	}
	if typ := getValidObjectCollection(typ); typ != Unknown {
		return typ
	}
	return Unknown
}

func ValidCollection(typ string) bool {
	return getValidCollection(typ) != Unknown
}
