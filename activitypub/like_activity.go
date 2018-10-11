package activitypub

import (
	as "github.com/mariusor/activitypub.go/activitystreams"
	"time"
)

// LikeActivity is the type for a create activity message
type LikeActivity struct {
	Activity  *as.Like          `jsonld:"activity"`
	Published time.Time         `jsonld:"published"`
	To        as.ItemCollection `jsonld:"to,omitempty,collapsible"`
	CC        as.ItemCollection `jsonld:"cc,omitempty,collapsible"`
}

// DislikeActivity is the type for a create activity message
type DislikeActivity struct {
	Activity  *as.Dislike       `jsonld:"activity"`
	Published time.Time         `jsonld:"published"`
	To        as.ItemCollection `jsonld:"to,omitempty,collapsible"`
	CC        as.ItemCollection `jsonld:"cc,omitempty,collapsible"`
}

func loadActorWithLikedObject(a as.Item, o as.Item) as.Item {
	typ := a.GetType()
	switch typ {
	case as.ApplicationType:
		var app as.Application
		app, _ = a.(as.Application)
		var liked *as.OrderedCollection
		if app.Liked == nil {
			liked = LikedNew()
		} else {
			liked = app.Liked.(*as.OrderedCollection)
		}
		liked.Append(o)
		app.Liked = liked
		return app
	case as.GroupType:
		var grp as.Group
		grp, _ = a.(as.Group)
		var liked *as.OrderedCollection
		if grp.Liked == nil {
			liked = LikedNew()
		} else {
			liked = grp.Liked.(*as.OrderedCollection)
		}
		liked.Append(o)
		grp.Liked = liked
		return grp
	case as.OrganizationType:
		var org as.Organization
		org, _ = a.(as.Organization)
		var liked *as.OrderedCollection
		if org.Liked == nil {
			liked = LikedNew()
		} else {
			liked = org.Liked.(*as.OrderedCollection)
		}
		liked.Append(o)
		org.Liked = liked
		return org
	case as.PersonType:
		var pers as.Person
		pers, _ = a.(as.Person)
		var liked *as.OrderedCollection
		if pers.Liked == nil {
			liked = LikedNew()
		} else {
			liked = pers.Liked.(*as.OrderedCollection)
		}
		liked.Append(o)
		pers.Liked = liked
		return pers
	case as.ServiceType:
		var serv as.Service
		serv, _ = a.(as.Service)
		var liked *as.OrderedCollection
		if serv.Liked == nil {
			liked = LikedNew()
		} else {
			liked = serv.Liked.(*as.OrderedCollection)
		}
		liked.Append(o)
		serv.Liked = liked
		return serv
	default:
		actor, _ := a.(as.Actor)
		var liked *as.OrderedCollection
		if actor.Liked == nil {
			liked = LikedNew()
		} else {
			liked = actor.Liked.(*as.OrderedCollection)
		}
		liked.Append(o)
		actor.Liked = liked
		return actor
	}
}

// LikeActivityNew initializes a new LikeActivity message
func LikeActivityNew(id as.ObjectID, a as.Item, o as.Item) LikeActivity {
	act := as.LikeNew(id, o)

	if a != nil {
		if a.IsObject() {
			act.Actor = loadActorWithLikedObject(a, o)
		}
		if a.IsLink() {
			act.Actor = a
		}
	}

	act.RecipientsDeduplication()

	c := LikeActivity{
		Activity:  act,
		Published: time.Now(),
	}

	return c
}

// DislikeActivityNew initializes a new LikeActivity message
func DislikeActivityNew(id as.ObjectID, a as.Item, o as.Item) DislikeActivity {
	act := as.DislikeNew(id, o)

	if a != nil {
		if a.IsObject() {
			act.Actor = loadActorWithLikedObject(a, o)
		}
		if a.IsLink() {
			act.Actor = a
		}
	}

	act.RecipientsDeduplication()

	d := DislikeActivity{
		Activity:  act,
		Published: time.Now(),
	}

	return d
}
