package activitypub

import (
	"time"

	as "github.com/mariusor/activitypub.go/activitystreams"
)

// CreateActivity is the type for a create activity message
type CreateActivity struct {
	Activity  *as.Create        `jsonld:"activity"`
	Published time.Time         `jsonld:"published"`
	To        as.ItemCollection `jsonld:"to,omitempty,collapsible"`
	CC        as.ItemCollection `jsonld:"cc,omitempty,collapsible"`
}

func loadActorWithInboxObject(a as.Item, o as.Item) as.Item {
	typ := a.GetType()
	switch typ {
	case as.ApplicationType:
		var app as.Application
		app, _ = a.(as.Application)
		var inbox *as.OrderedCollection
		if app.Inbox == nil {
			inbox = InboxNew()
		} else {
			inbox = app.Inbox.(*as.OrderedCollection)
		}
		inbox.Append(o)
		app.Inbox = inbox
		return app
	case as.GroupType:
		var grp as.Group
		grp, _ = a.(as.Group)
		var inbox *as.OrderedCollection
		if grp.Inbox == nil {
			inbox = InboxNew()
		} else {
			inbox = grp.Inbox.(*as.OrderedCollection)
		}
		inbox.Append(o)
		grp.Inbox = inbox
		return grp
	case as.OrganizationType:
		var org as.Organization
		org, _ = a.(as.Organization)
		var inbox *as.OrderedCollection
		if org.Inbox == nil {
			inbox = InboxNew()
		} else {
			inbox = org.Inbox.(*as.OrderedCollection)
		}
		inbox.Append(o)
		org.Inbox = inbox
		return org
	case as.PersonType:
		var pers as.Person
		pers, _ = a.(as.Person)
		var inbox *as.OrderedCollection
		if pers.Inbox == nil {
			inbox = InboxNew()
		} else {
			inbox = pers.Inbox.(*as.OrderedCollection)
		}
		inbox.Append(o)
		pers.Inbox = inbox
		return pers
	case as.ServiceType:
		var serv as.Service
		serv, _ = a.(as.Service)
		var inbox *as.OrderedCollection
		if serv.Inbox == nil {
			inbox = InboxNew()
		} else {
			inbox = serv.Inbox.(*as.OrderedCollection)
		}
		inbox.Append(o)
		serv.Inbox = inbox
		return serv
	default:
		actor, _ := a.(as.Actor)
		var inbox *as.OrderedCollection
		if actor.Inbox == nil {
			inbox = InboxNew()
		} else {
			inbox = actor.Inbox.(*as.OrderedCollection)
		}
		inbox.Append(o)
		actor.Inbox = inbox
		return actor
	}
}

// CreateActivityNew initializes a new CreateActivity message
func CreateActivityNew(id as.ObjectID, a as.Item, o as.Item) CreateActivity {
	act := as.CreateNew(id, o)

	if a != nil {
		if a.IsObject() {
			act.Actor = loadActorWithInboxObject(a, o)
		}
		if a.IsLink() {
			act.Actor = a
		}
	}

	act.RecipientsDeduplication()

	c := CreateActivity{
		Activity:  act,
		Published: time.Now(),
	}

	return c
}
