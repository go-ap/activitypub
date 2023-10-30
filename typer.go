package activitypub

import (
	"path/filepath"
	"strings"

	"github.com/go-ap/errors"
)

// CollectionPath
type CollectionPath string

// CollectionPaths
type CollectionPaths []CollectionPath

const (
	Unknown = CollectionPath("")
	// Outbox
	//
	// https://www.w3.org/TR/activitypub/#outbox
	//
	// The outbox is discovered through the outbox property of an actor's profile.
	// The outbox MUST be an OrderedCollection.
	//
	// The outbox stream contains activities the user has published, subject to the ability of the requestor
	// to retrieve the activity (that is, the contents of the outbox are filtered by the permissions
	// of the person reading it). If a user submits a request without Authorization the server should respond
	// with all of the Public posts. This could potentially be all relevant objects published by the user,
	// though the number of available items is left to the discretion of those implementing and deploying the server.
	//
	// The outbox accepts HTTP POST requests, with behaviour described in Client to Server Interactions.
	Outbox = CollectionPath("outbox")
	// Inbox
	//
	// https://www.w3.org/TR/activitypub/#inbox
	//
	// The inbox is discovered through the inbox property of an actor's profile. The inbox MUST be an OrderedCollection.
	//
	// The inbox stream contains all activities received by the actor. The server SHOULD filter content according
	// to the requester's permission. In general, the owner of an inbox is likely to be able to access
	// all of their inbox contents. Depending on access control, some other content may be public,
	// whereas other content may require authentication for non-owner users, if they can access the inbox at all.
	//
	// The server MUST perform de-duplication of activities returned by the inbox. Duplication can occur
	// if an activity is addressed both to an actor's followers, and a specific actor who also follows
	// the recipient actor, and the server has failed to de-duplicate the recipients list.
	// Such deduplication MUST be performed by comparing the id of the activities and dropping any activities already seen.
	//
	// The inboxes of actors on federated servers accepts HTTP POST requests, with behaviour described in Delivery.
	// Non-federated servers SHOULD return a 405 Method Not Allowed upon receipt of a POST request.
	Inbox = CollectionPath("inbox")
	// Followers
	//
	// https://www.w3.org/TR/activitypub/#followers
	//
	// Every actor SHOULD have a followers collection. This is a list of everyone who has sent a Follow activity
	// for the actor, added as a side effect. This is where one would find a list of all the actors that are following
	// the actor. The followers collection MUST be either an OrderedCollection or a Collection and MAY be filtered
	// on privileges of an authenticated user or as appropriate when no authentication is given.
	//
	// NOTE: Default for notification targeting
	// The follow activity generally is a request to see the objects an actor creates.
	// This makes the Followers collection an appropriate default target for delivery of notifications.
	Followers = CollectionPath("followers")
	// Following
	//
	// https://www.w3.org/TR/activitypub/#following
	//
	// Every actor SHOULD have a following collection. This is a list of everybody that the actor has followed,
	// added as a side effect. The following collection MUST be either an OrderedCollection or a Collection
	// and MAY be filtered on privileges of an authenticated user or as appropriate when no authentication is given.
	Following = CollectionPath("following")
	// Liked
	//
	// https://www.w3.org/TR/activitypub/#liked
	//
	// Every actor MAY have a liked collection. This is a list of every object from all of the actor's Like activities,
	// added as a side effect. The liked collection MUST be either an OrderedCollection or a Collection and
	// MAY be filtered on privileges of an authenticated user or as appropriate when no authentication is given.
	Liked = CollectionPath("liked")
	// Likes
	//
	// https://www.w3.org/TR/activitypub/#likes
	//
	// Every object MAY have a likes collection. This is a list of all Like activities with this object as the object
	// property, added as a side effect. The likes collection MUST be either an OrderedCollection or a Collection
	// and MAY be filtered on privileges of an authenticated user or as appropriate when no authentication is given.
	//
	// NOTE
	// Care should be taken to not confuse the the likes collection with the similarly named but different liked
	// collection. In sum:
	//
	// * liked: Specifically a property of actors. This is a collection of Like activities performed by the actor,
	// added to the collection as a side effect of delivery to the outbox.
	// * likes: May be a property of any object. This is a collection of Like activities referencing this object,
	// added to the collection as a side effect of delivery to the inbox.
	Likes = CollectionPath("likes")
	// Shares
	//
	// https://www.w3.org/TR/activitypub/#shares
	//
	// Every object MAY have a shares collection. This is a list of all Announce activities with this object
	// as the object property, added as a side effect. The shares collection MUST be either an OrderedCollection
	// or a Collection and MAY be filtered on privileges of an authenticated user or as appropriate when
	// no authentication is given.
	Shares  = CollectionPath("shares")
	Replies = CollectionPath("replies") // activitystreams
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
		if strings.EqualFold(string(typ), string(tt)) {
			return true
		}
	}
	return false
}

// Split splits the IRI in an actor IRI and its CollectionPath
// if the CollectionPath is found in the elements in the t CollectionPaths slice
func (t CollectionPaths) Split(i IRI) (IRI, CollectionPath) {
	if u, err := i.URL(); err == nil {
		maybeActor, maybeCol := filepath.Split(u.Path)
		if len(maybeActor) == 0 {
			return i, Unknown
		}
		tt := CollectionPath(maybeCol)
		if !t.Contains(tt) {
			tt = ""
		}
		u.Path = strings.TrimRight(maybeActor, "/")
		iri := IRI(u.String())
		return iri, tt
	}
	maybeActor, maybeCol := filepath.Split(i.String())
	if len(maybeActor) == 0 {
		return i, Unknown
	}
	tt := CollectionPath(maybeCol)
	if !t.Contains(tt) {
		return i, Unknown
	}
	maybeActor = strings.TrimRight(maybeActor, "/")
	return IRI(maybeActor), tt
}

// IRIf formats an IRI from an existing IRI and the CollectionPath type
func IRIf(i IRI, t CollectionPath) IRI {
	si := i.String()
	s := strings.Builder{}
	_, _ = s.WriteString(si)
	if l := len(si); l == 0 || si[l-1] != '/' {
		_, _ = s.WriteRune('/')
	}
	_, _ = s.WriteString(string(t))
	return IRI(s.String())
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
	if strings.EqualFold(maybeCol, string(t)) {
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
		if strings.EqualFold(string(typ), string(t)) {
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
