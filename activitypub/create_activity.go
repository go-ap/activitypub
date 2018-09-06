package activitypub

import "time"

// CreateActivity is the type for a create activity message
type CreateActivity struct {
	Activity  *Create
	Published time.Time
	To        ObjectsArr
	CC        ObjectsArr
}

func loadActorWithInboxObject(a ObjectOrLink, o ObjectOrLink) ObjectOrLink {
	typ := a.GetType()
	switch typ {
	case ApplicationType:
		var app Application
		app, _ = a.(Application)
		if app.Inbox == nil {
			app.Inbox = InboxNew()
		}
		app.Inbox.Append(o)
		return app
	case GroupType:
		var grp Group
		grp, _ = a.(Group)
		if grp.Inbox == nil {
			grp.Inbox = InboxNew()
		}
		grp.Inbox.Append(o)
		return grp
	case OrganizationType:
		var org Organization
		org, _ = a.(Organization)
		if org.Inbox == nil {
			org.Inbox = InboxNew()
		}
		org.Inbox.Append(o)
		return org
	case PersonType:
		var pers Person
		pers, _ = a.(Person)
		if pers.Inbox == nil {
			pers.Inbox = InboxNew()
		}
		pers.Inbox.Append(o)
		return pers
	case ServiceType:
		var serv Service
		serv, _ = a.(Service)
		serv.Inbox.Append(o)
		return serv
	default:
		actor, _ := a.(Actor)
		if actor.Inbox == nil {
			actor.Inbox = InboxNew()
		}
		actor.Inbox.Append(o)
		return actor
	}
}

// CreateActivityNew initializes a new CreateActivity message
func CreateActivityNew(id ObjectID, a ObjectOrLink, o ObjectOrLink) CreateActivity {
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
