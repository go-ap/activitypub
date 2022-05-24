package handlers

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	pub "github.com/go-ap/activitypub"
	"github.com/go-ap/errors"
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

// Split splits the IRI in an actor IRI and its collection
// if the collection is found in the elements in the t CollectionTypes slice
func (t CollectionTypes) Split(i pub.IRI) (pub.IRI, CollectionType) {
	maybeActor, maybeCol := path.Split(i.String())
	tt := CollectionType(maybeCol)
	if !t.Contains(tt) {
		tt = ""
		maybeActor = i.String()
	}
	iri := pub.IRI(strings.TrimRight(maybeActor, "/"))
	return iri, tt
}

// IRIf formats an IRI from an existing IRI and the collection type
func IRIf(i pub.IRI, t CollectionType) pub.IRI {
	onePastLast := len(i)
	if onePastLast > 1 && i[onePastLast-1] == '/' {
		i = i[:onePastLast-1]
	}
	return pub.IRI(fmt.Sprintf("%s/%s", i, t))
}

// IRI gives us the IRI of the t collection type corresponding to the i Item,
// or generates a new one if not found.
func (t CollectionType) IRI(i pub.Item) pub.IRI {
	if pub.IsNil(i) {
		return IRIf("", t)
	}
	if pub.IsObject(i) {
		if it := t.Of(i); !pub.IsNil(it) {
			return it.GetLink()
		}
	}
	return IRIf(i.GetLink(), t)
}

// Of gives us the property of the i Item that corresponds to the t collection type.
func (t CollectionType) Of(i pub.Item) pub.Item {
	if pub.IsNil(i) || !i.IsObject() {
		return nil
	}
	var it pub.Item
	if OnActor.Contains(t) && pub.ActorTypes.Contains(i.GetType()) {
		pub.OnActor(i, func(a *pub.Actor) error {
			switch t {
			case Inbox:
				it = a.Inbox
			case Outbox:
				it = a.Outbox
			case Liked:
				it = a.Liked
			case Following:
				it = a.Following
			case Followers:
				it = a.Followers
			}
			return nil
		})
	}
	pub.OnObject(i, func(o *pub.Object) error {
		switch t {
		case Likes:
			it = o.Likes
		case Shares:
			it = o.Shares
		case Replies:
			it = o.Replies
		}
		return nil
	})

	return it
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

// Split returns the base IRI of received i, if i represents an IRI matching collection type t
func Split(i pub.IRI) (pub.IRI, CollectionType) {
	return ActivityPubCollections.Split(i)
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

func ValidCollectionIRI(i pub.IRI) bool {
	_, t := Split(i)
	return getValidCollection(t) != Unknown
}

// AddTo adds collection type IRI on the corresponding property of the i Item
func (t CollectionType) AddTo(i pub.Item) (pub.IRI, bool) {
	if pub.IsNil(i) || !i.IsObject() {
		return pub.NilIRI, false
	}
	status := false
	var iri pub.IRI
	if OnActor.Contains(t) {
		pub.OnActor(i, func(a *pub.Actor) error {
			if status = t == Inbox && pub.IsNil(a.Inbox); status {
				a.Inbox = IRIf(a.GetLink(), t)
				iri = a.Inbox.GetLink()
			} else if status = t == Outbox && pub.IsNil(a.Outbox); status {
				a.Outbox = IRIf(a.GetLink(), t)
				iri = a.Outbox.GetLink()
			} else if status = t == Liked && pub.IsNil(a.Liked); status {
				a.Liked = IRIf(a.GetLink(), t)
				iri = a.Liked.GetLink()
			} else if status = t == Following && pub.IsNil(a.Following); status {
				a.Following = IRIf(a.GetLink(), t)
				iri = a.Following.GetLink()
			} else if status = t == Followers && pub.IsNil(a.Followers); status {
				a.Followers = IRIf(a.GetLink(), t)
				iri = a.Followers.GetLink()
			}
			return nil
		})
	} else if OnObject.Contains(t) {
		pub.OnObject(i, func(o *pub.Object) error {
			if status = t == Likes && pub.IsNil(o.Likes); status {
				o.Likes = IRIf(o.GetLink(), t)
				iri = o.Likes.GetLink()
			} else if status = t == Shares && pub.IsNil(o.Shares); status {
				o.Shares = IRIf(o.GetLink(), t)
				iri = o.Shares.GetLink()
			} else if status = t == Replies && pub.IsNil(o.Replies); status {
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
