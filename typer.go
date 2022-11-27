package activitypub

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/go-ap/errors"
)

// CollectionPath
type CollectionPath string

// CollectionPaths
type CollectionPaths []CollectionPath

const (
	Unknown   = CollectionPath("")
	Outbox    = CollectionPath("outbox")
	Inbox     = CollectionPath("inbox")
	Shares    = CollectionPath("shares")
	Replies   = CollectionPath("replies") // activitystreams
	Following = CollectionPath("following")
	Followers = CollectionPath("followers")
	Liked     = CollectionPath("liked")
	Likes     = CollectionPath("likes")
)

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

func (t CollectionPaths) Contains(typ CollectionPath) bool {
	for _, tt := range t {
		if strings.ToLower(string(typ)) == string(tt) {
			return true
		}
	}
	return false
}

// Split splits the IRI in an actor IRI and its CollectionPath
// if the CollectionPath is found in the elements in the t CollectionPaths slice
func (t CollectionPaths) Split(i IRI) (IRI, CollectionPath) {
	maybeActor, maybeCol := filepath.Split(i.String())
	tt := CollectionPath(maybeCol)
	if !t.Contains(tt) {
		tt = ""
		maybeActor = i.String()
	}
	iri := IRI(strings.TrimRight(maybeActor, "/"))
	return iri, tt
}

// IRIf formats an IRI from an existing IRI and the CollectionPath type
func IRIf(i IRI, t CollectionPath) IRI {
	onePastLast := len(i)
	if onePastLast > 1 && i[onePastLast-1] == '/' {
		i = i[:onePastLast-1]
	}
	return IRI(fmt.Sprintf("%s/%s", i, t))
}

// IRI gives us the IRI of the t CollectionPath type corresponding to the i Item,
// or generates a new one if not found.
func (t CollectionPath) IRI(i Item) IRI {
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

func (t CollectionPath) ofItemCollection(col ItemCollection) Item {
	iriCol := make(ItemCollection, len(col))
	for i, it := range col {
		iriCol[i] = t.Of(it)
	}
	return iriCol
}

func (t CollectionPath) ofObject(ob *Object) Item {
	var it Item
	switch t {
	case Likes:
		it = ob.Likes
	case Shares:
		it = ob.Shares
	case Replies:
		it = ob.Replies
	}
	if it == nil {
		it = t.ofIRI(ob.ID)
	}
	return it
}
func (t CollectionPath) ofActor(a *Actor) Item {
	var it Item
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
	if it == nil {
		it = t.ofIRI(a.ID)
	}
	return it
}

func (t CollectionPath) ofIRI(iri IRI) Item {
	if len(iri) == 0 {
		return nil
	}
	return iri.AddPath(string(t))
}

func (t CollectionPath) ofItem(i Item) Item {
	var it Item
	return it
}

// Of gives us the property of the i Item that corresponds to the t CollectionPath type.
func (t CollectionPath) Of(i Item) Item {
	if IsNil(i) {
		return nil
	}
	it := t.ofIRI(i.GetLink())
	if IsItemCollection(i) {
		OnItemCollection(i, func(col *ItemCollection) error {
			it = t.ofItemCollection(*col)
			return nil
		})
	}
	if OfActor.Contains(t) && ActorTypes.Contains(i.GetType()) {
		OnActor(i, func(a *Actor) error {
			it = t.ofActor(a)
			return nil
		})
	}
	OnObject(i, func(o *Object) error {
		it = t.ofObject(o)
		return nil
	})
	return it
}

// OfActor returns the base IRI of received i, if i represents an IRI matching CollectionPath type t
func (t CollectionPath) OfActor(i IRI) (IRI, error) {
	maybeActor, maybeCol := filepath.Split(i.String())
	if strings.ToLower(maybeCol) == strings.ToLower(string(t)) {
		maybeActor = strings.TrimRight(maybeActor, "/")
		return IRI(maybeActor), nil
	}
	return EmptyIRI, errors.Newf("IRI does not represent a valid %s CollectionPath", t)
}

// Split returns the base IRI of received i, if i represents an IRI matching CollectionPath type t
func Split(i IRI) (IRI, CollectionPath) {
	return ActivityPubCollections.Split(i)
}

func getValidActivityCollection(t CollectionPath) CollectionPath {
	if validActivityCollection.Contains(t) {
		return t
	}
	return Unknown
}

// ValidActivityCollection shows if the current ActivityPub end-point type is a valid one for handling Activities
func ValidActivityCollection(typ CollectionPath) bool {
	return getValidActivityCollection(typ) != Unknown
}

var validObjectCollection = []CollectionPath{
	Following,
	Followers,
	Liked,
}

func getValidObjectCollection(typ CollectionPath) CollectionPath {
	for _, t := range validObjectCollection {
		if strings.ToLower(string(typ)) == string(t) {
			return t
		}
	}
	return Unknown
}

// ValidActivityCollection shows if the current ActivityPub end-point type is a valid one for handling Objects
func ValidObjectCollection(typ CollectionPath) bool {
	return getValidObjectCollection(typ) != Unknown
}

func getValidCollection(typ CollectionPath) CollectionPath {
	if typ := getValidActivityCollection(typ); typ != Unknown {
		return typ
	}
	if typ := getValidObjectCollection(typ); typ != Unknown {
		return typ
	}
	return Unknown
}

func ValidCollection(typ CollectionPath) bool {
	return getValidCollection(typ) != Unknown
}

func ValidCollectionIRI(i IRI) bool {
	_, t := Split(i)
	return getValidCollection(t) != Unknown
}

// AddTo adds CollectionPath type IRI on the corresponding property of the i Item
func (t CollectionPath) AddTo(i Item) (IRI, bool) {
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
