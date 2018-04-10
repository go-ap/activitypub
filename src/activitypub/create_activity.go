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

	ok := false
	if a != nil {
		typ := a.Object().Type
		if typ == ApplicationType {
			var app Application
			app, ok = a.(Application)
			act.Actor = Actor(app)
		}
		if typ == GroupType {
			var grp Group
			grp, ok = a.(Group)
			act.Actor = Actor(grp)
		}
		if typ == OrganizationType {
			var org Organization
			org, ok = a.(Organization)
			act.Actor = Actor(org)
		}
		if typ == PersonType {
			var pers Person
			pers, ok = a.(Person)
			act.Actor = Actor(pers)
		}
		if typ == ServiceType {
			var serv Service
			serv, ok = a.(Service)
			act.Actor = Actor(serv)
		}
		if !ok {
			act.Actor, ok = a.(Actor)
		}
	}

	act.Actor.Inbox.Append(o)

	c := CreateActivity{
		Activity:  act,
		Published: time.Now(),
	}

	return c
}
