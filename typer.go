package handlers

import (
	"fmt"
	pub "github.com/go-ap/activitypub"
	"github.com/go-ap/errors"
	"net/http"
	"path"
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
// @TODO(marius): This should be moved as a property on an instantiable package object, instead of keeping it here
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
	col := Unknown
	pathElements := strings.Split(r.URL.Path[1:], "/") // Skip first /
	for i := len(pathElements) - 1; i >= 0; i-- {
		col = CollectionType(pathElements[i])
		if typ := getValidActivityCollection(col); typ != Unknown {
			return typ
		}
		if typ := getValidObjectCollection(col); typ != Unknown {
			return typ
		}
	}

	return col
}

var (
	validActivityCollection = CollectionTypes{
		Outbox,
		Inbox,
		Likes,
		Shares,
		Replies, // activitystreams
	}
	OnObject = CollectionTypes{
		Likes,
		Shares,
		Replies,
	}
	OnActor = CollectionTypes{
		Outbox,
		Inbox,
		Liked,
		Following,
		Followers,
	}

	ActivityPubCollections = CollectionTypes{
		Outbox,
		Inbox,
		Liked,
		Following,
		Followers,
		Likes,
		Shares,
		Replies,
	}
)

func (t CollectionTypes) Contains(typ CollectionType) bool {
	for _, tt := range t {
		if strings.ToLower(string(typ)) == string(tt) {
			return true
		}
	}
	return false
}

// IRIf formats an IRI from an existing IRI and the collection type
func IRIf(i pub.IRI, t CollectionType) pub.IRI {
	return pub.IRI(fmt.Sprintf("%s/%s", i, t))
}

// IRI gives us the property of the i Item that corresponds to the t collection type
// or generates a new one if not found.
func (t CollectionType) IRI(i pub.Item) pub.IRI {
	var iri pub.IRI
	if i == nil {
		return pub.EmptyIRI
	}
	if i.IsObject() {
		if OnActor.Contains(t) {
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
		if OnObject.Contains(t) {
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

	if len(iri) > 0 {
		return iri
	}
	return IRIf(i.GetLink(), t)
}

// OfActor returns the base IRI of received i, if i represents an IRI matching collection type t
func (t CollectionType) OfActor(i pub.IRI) (pub.IRI, error) {
	maybeActor, maybeCol := path.Split(i.String())
	if strings.ToLower(maybeCol) == strings.ToLower(string(t)) {
		maybeActor = strings.TrimRight(maybeActor, "/")
		return pub.IRI(maybeActor), nil
	}
	return pub.EmptyIRI, errors.Newf("IRI does not represent a valid %s collection", t)
}

// OfActor returns the base IRI of received i, if i represents an IRI matching collection type t
func Split(i pub.IRI) (pub.IRI, CollectionType) {
	maybeActor, maybeCol := path.Split(i.String())
	t := CollectionType(maybeCol)
	iri := pub.IRI(strings.TrimRight(maybeActor, "/"))
	return iri, t
}

func getValidActivityCollection(t CollectionType) CollectionType {
	if validActivityCollection.Contains(t) {
		return t
	}
	return Unknown
}

// ValidActivityCollection shows if the current ActivityPub end-point type is a valid one for handling Activities
func ValidActivityCollection(typ CollectionType) bool {
	return getValidActivityCollection(typ) != Unknown
}

var validObjectCollection = []CollectionType{
	Following,
	Followers,
	Liked,
}

func getValidObjectCollection(typ CollectionType) CollectionType {
	for _, t := range validObjectCollection {
		if strings.ToLower(string(typ)) == string(t) {
			return t
		}
	}
	return Unknown
}

// ValidActivityCollection shows if the current ActivityPub end-point type is a valid one for handling Objects
func ValidObjectCollection(typ CollectionType) bool {
	return getValidObjectCollection(typ) != Unknown
}

func getValidCollection(typ CollectionType) CollectionType {
	if typ := getValidActivityCollection(typ); typ != Unknown {
		return typ
	}
	if typ := getValidObjectCollection(typ); typ != Unknown {
		return typ
	}
	return Unknown
}

func ValidCollection(typ CollectionType) bool {
	return getValidCollection(typ) != Unknown
}

// AddTo adds collection type IRI on the corresponding property of the i Item
func (t CollectionType) AddTo(i pub.Item) (pub.IRI, bool) {
	if i == nil || !i.IsObject() {
		return pub.NilIRI, false
	}
	status := false
	var iri pub.IRI
	if OnActor.Contains(t) {
		pub.OnActor(i, func(a *pub.Actor) error {
			if status = t == Inbox && a.Inbox == nil; status {
				a.Inbox = IRIf(a.GetLink(), t)
				iri = a.Inbox.GetLink()
			} else if status = t == Outbox && a.Outbox == nil; status {
				a.Outbox = IRIf(a.GetLink(), t)
				iri = a.Outbox.GetLink()
			} else if status = t == Liked && a.Liked == nil; status {
				a.Liked = IRIf(a.GetLink(), t)
				iri = a.Liked.GetLink()
			} else if status = t == Following && a.Following == nil; status {
				a.Following = IRIf(a.GetLink(), t)
				iri = a.Following.GetLink()
			} else if status = t == Followers && a.Followers == nil; status {
				a.Followers = IRIf(a.GetLink(), t)
				iri = a.Followers.GetLink()
			}
			return nil
		})
	} else if OnObject.Contains(t) {
		pub.OnObject(i, func(o *pub.Object) error {
			if status = t == Likes && o.Likes == nil; status {
				o.Likes = IRIf(o.GetLink(), t)
				iri = o.Likes.GetLink()
			} else if status = t == Shares && o.Shares == nil; status {
				o.Shares = IRIf(o.GetLink(), t)
				iri = o.Shares.GetLink()
			} else if status = t == Replies && o.Replies == nil; status {
				o.Replies = IRIf(o.GetLink(), t)
				iri = o.Replies.GetLink()
			}
			return nil
		})
	} else {
		iri = IRIf(i.GetLink(), t)
	}
	return iri, status
}
