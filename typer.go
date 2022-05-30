package activitypub

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/go-ap/errors"
)

// collectionPath
type collectionPath string

// CollectionPaths
type CollectionPaths []collectionPath

const (
	Unknown   = collectionPath("")
	Outbox    = collectionPath("outbox")
	Inbox     = collectionPath("inbox")
	Shares    = collectionPath("shares")
	Replies   = collectionPath("replies") // activitystreams
	Following = collectionPath("following")
	Followers = collectionPath("followers")
	Liked     = collectionPath("liked")
	Likes     = collectionPath("likes")
)

func CollectionPath(s string) collectionPath {
	return collectionPath(s)
}

// Typer is the static package variable that determines a collectionPath type for a particular request
// It can be overloaded from outside packages.
// @TODO(marius): This should be moved as a property on an instantiable package object, instead of keeping it here
var Typer CollectionTyper = pathTyper{}

// CollectionTyper allows external packages to tell us which collectionPath the current HTTP request addresses
type CollectionTyper interface {
	Type(r *http.Request) collectionPath
}

type pathTyper struct{}

func (d pathTyper) Type(r *http.Request) collectionPath {
	if r.URL == nil || len(r.URL.Path) == 0 {
		return Unknown
	}
	col := Unknown
	pathElements := strings.Split(r.URL.Path[1:], "/") // Skip first /
	for i := len(pathElements) - 1; i >= 0; i-- {
		col = collectionPath(pathElements[i])
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
	validActivityCollection = CollectionPaths{
		Outbox,
		Inbox,
		Likes,
		Shares,
		Replies, // activitystreams
	}
	OfObject = CollectionPaths{
		Likes,
		Shares,
		Replies,
	}
	OfActor = CollectionPaths{
		Outbox,
		Inbox,
		Liked,
		Following,
		Followers,
	}

	ActivityPubCollections = CollectionPaths{
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

func (t CollectionPaths) Contains(typ collectionPath) bool {
	for _, tt := range t {
		if strings.ToLower(string(typ)) == string(tt) {
			return true
		}
	}
	return false
}

// Split splits the IRI in an actor IRI and its collectionPath
// if the collectionPath is found in the elements in the t CollectionPaths slice
func (t CollectionPaths) Split(i IRI) (IRI, collectionPath) {
	maybeActor, maybeCol := path.Split(i.String())
	tt := collectionPath(maybeCol)
	if !t.Contains(tt) {
		tt = ""
		maybeActor = i.String()
	}
	iri := IRI(strings.TrimRight(maybeActor, "/"))
	return iri, tt
}

// IRIf formats an IRI from an existing IRI and the collectionPath type
func IRIf(i IRI, t collectionPath) IRI {
	onePastLast := len(i)
	if onePastLast > 1 && i[onePastLast-1] == '/' {
		i = i[:onePastLast-1]
	}
	return IRI(fmt.Sprintf("%s/%s", i, t))
}

// IRI gives us the IRI of the t collectionPath type corresponding to the i Item,
// or generates a new one if not found.
func (t collectionPath) IRI(i Item) IRI {
	if IsNil(i) {
		return IRIf("", t)
	}
	if IsObject(i) {
		if it := t.Of(i); !IsNil(it) {
			return it.GetLink()
		}
	}
	return IRIf(i.GetLink(), t)
}

// Of gives us the property of the i Item that corresponds to the t collectionPath type.
func (t collectionPath) Of(i Item) Item {
	if IsNil(i) || !i.IsObject() {
		return nil
	}
	var it Item
	if OfActor.Contains(t) && ActorTypes.Contains(i.GetType()) {
		OnActor(i, func(a *Actor) error {
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
	OnObject(i, func(o *Object) error {
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

// OfActor returns the base IRI of received i, if i represents an IRI matching collectionPath type t
func (t collectionPath) OfActor(i IRI) (IRI, error) {
	maybeActor, maybeCol := path.Split(i.String())
	if strings.ToLower(maybeCol) == strings.ToLower(string(t)) {
		maybeActor = strings.TrimRight(maybeActor, "/")
		return IRI(maybeActor), nil
	}
	return EmptyIRI, errors.Newf("IRI does not represent a valid %s collectionPath", t)
}

// Split returns the base IRI of received i, if i represents an IRI matching collectionPath type t
func Split(i IRI) (IRI, collectionPath) {
	return ActivityPubCollections.Split(i)
}

func getValidActivityCollection(t collectionPath) collectionPath {
	if validActivityCollection.Contains(t) {
		return t
	}
	return Unknown
}

// ValidActivityCollection shows if the current ActivityPub end-point type is a valid one for handling Activities
func ValidActivityCollection(typ collectionPath) bool {
	return getValidActivityCollection(typ) != Unknown
}

var validObjectCollection = []collectionPath{
	Following,
	Followers,
	Liked,
}

func getValidObjectCollection(typ collectionPath) collectionPath {
	for _, t := range validObjectCollection {
		if strings.ToLower(string(typ)) == string(t) {
			return t
		}
	}
	return Unknown
}

// ValidActivityCollection shows if the current ActivityPub end-point type is a valid one for handling Objects
func ValidObjectCollection(typ collectionPath) bool {
	return getValidObjectCollection(typ) != Unknown
}

func getValidCollection(typ collectionPath) collectionPath {
	if typ := getValidActivityCollection(typ); typ != Unknown {
		return typ
	}
	if typ := getValidObjectCollection(typ); typ != Unknown {
		return typ
	}
	return Unknown
}

func ValidCollection(typ collectionPath) bool {
	return getValidCollection(typ) != Unknown
}

func ValidCollectionIRI(i IRI) bool {
	_, t := Split(i)
	return getValidCollection(t) != Unknown
}

// AddTo adds collectionPath type IRI on the corresponding property of the i Item
func (t collectionPath) AddTo(i Item) (IRI, bool) {
	if IsNil(i) || !i.IsObject() {
		return NilIRI, false
	}
	status := false
	var iri IRI
	if OfActor.Contains(t) {
		OnActor(i, func(a *Actor) error {
			if status = t == Inbox && IsNil(a.Inbox); status {
				a.Inbox = IRIf(a.GetLink(), t)
				iri = a.Inbox.GetLink()
			} else if status = t == Outbox && IsNil(a.Outbox); status {
				a.Outbox = IRIf(a.GetLink(), t)
				iri = a.Outbox.GetLink()
			} else if status = t == Liked && IsNil(a.Liked); status {
				a.Liked = IRIf(a.GetLink(), t)
				iri = a.Liked.GetLink()
			} else if status = t == Following && IsNil(a.Following); status {
				a.Following = IRIf(a.GetLink(), t)
				iri = a.Following.GetLink()
			} else if status = t == Followers && IsNil(a.Followers); status {
				a.Followers = IRIf(a.GetLink(), t)
				iri = a.Followers.GetLink()
			}
			return nil
		})
	} else if OfObject.Contains(t) {
		OnObject(i, func(o *Object) error {
			if status = t == Likes && IsNil(o.Likes); status {
				o.Likes = IRIf(o.GetLink(), t)
				iri = o.Likes.GetLink()
			} else if status = t == Shares && IsNil(o.Shares); status {
				o.Shares = IRIf(o.GetLink(), t)
				iri = o.Shares.GetLink()
			} else if status = t == Replies && IsNil(o.Replies); status {
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
