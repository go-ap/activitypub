package activitypub

import "time"

// CreateActivity is the type for a create activity message
type CreateActivity struct {
	Activity  *Create        `jsonld:"activity"`
	Published time.Time      `jsonld:"published"`
	To        ItemCollection `jsonld:"to,omitempty,collapsible"`
	CC        ItemCollection `jsonld:"cc,omitempty,collapsible"`
}

func loadActorWithInboxObject(a Item, o Item) Item {
	typ := a.GetType()
	switch typ {
	case ApplicationType:
		var app Application
		app, _ = a.(Application)
		var inbox *Inbox
		if app.Inbox == nil {
			inbox = InboxNew()
		} else {
			inbox = app.Inbox.(*Inbox)
		}
		inbox.Append(o)
		app.Inbox = inbox
		return app
	case GroupType:
		var grp Group
		grp, _ = a.(Group)
		var inbox *Inbox
		if grp.Inbox == nil {
			inbox = InboxNew()
		} else {
			inbox = grp.Inbox.(*Inbox)
		}
		inbox.Append(o)
		grp.Inbox = inbox
		return grp
	case OrganizationType:
		var org Organization
		org, _ = a.(Organization)
		var inbox *Inbox
		if org.Inbox == nil {
			inbox = InboxNew()
		} else {
			inbox = org.Inbox.(*Inbox)
		}
		inbox.Append(o)
		org.Inbox = inbox
		return org
	case PersonType:
		var pers Person
		pers, _ = a.(Person)
		var inbox *Inbox
		if pers.Inbox == nil {
			inbox = InboxNew()
		} else {
			inbox = pers.Inbox.(*Inbox)
		}
		inbox.Append(o)
		pers.Inbox = inbox
		return pers
	case ServiceType:
		var serv Service
		serv, _ = a.(Service)
		var inbox *Inbox
		if serv.Inbox == nil {
			inbox = InboxNew()
		} else {
			inbox = serv.Inbox.(*Inbox)
		}
		inbox.Append(o)
		serv.Inbox = inbox
		return serv
	default:
		actor, _ := a.(Actor)
		var inbox *Inbox
		if actor.Inbox == nil {
			inbox = InboxNew()
		} else {
			inbox = actor.Inbox.(*Inbox)
		}
		inbox.Append(o)
		actor.Inbox = inbox
		return actor
	}
}

// CreateActivityNew initializes a new CreateActivity message
func CreateActivityNew(id ObjectID, a Item, o Item) CreateActivity {
	act := CreateNew(id, o)

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

// UnmarshalJSON
func (c *Create) UnmarshalJSON(data []byte) error {
	a := Activity(*c)
	err := a.UnmarshalJSON(data)

	*c = Create(a)

	return err
}
