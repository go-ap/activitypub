package activitypub

import "time"

// CreateActivity is the type for a create activity message
type CreateActivity struct {
	Activity  *Create
	Published time.Time
	To        ObjectsArr
	CC        ObjectsArr
}

// CreateActivityNew initializes a new CreateActivity message
func CreateActivityNew(id ObjectID, a ObjectOrLink, o ObjectOrLink) CreateActivity {
	act := CreateNew(id, o)

	if a != nil {
		typ := a.GetType()
		switch typ {
		case ApplicationType:
			var app Application
			app, _ = a.(Application)
			if app.Inbox == nil {
				app.Inbox = InboxNew()
			}
			app.Inbox.Append(o)
			act.Actor = app
		case GroupType:
			var grp Group
			grp, _ = a.(Group)
			if grp.Inbox == nil {
				grp.Inbox = InboxNew()
			}
			grp.Inbox.Append(o)
			act.Actor = grp
		case OrganizationType:
			var org Organization
			org, _ = a.(Organization)
			if org.Inbox == nil {
				org.Inbox = InboxNew()
			}
			org.Inbox.Append(o)
			act.Actor = org
		case PersonType:
			var pers Person
			pers, _ = a.(Person)
			if pers.Inbox == nil {
				pers.Inbox = InboxNew()
			}
			pers.Inbox.Append(o)
			act.Actor = pers
		case ServiceType:
			var serv Service
			serv, _ = a.(Service)
			serv.Inbox.Append(o)
			act.Actor = serv
		default:
			actor, _ := a.(Actor)
			if actor.Inbox == nil {
				actor.Inbox = InboxNew()
			}
			actor.Inbox.Append(o)
			act.Actor = actor
		}
	}

	c := CreateActivity{
		Activity:  act,
		Published: time.Now(),
	}

	return c
}
