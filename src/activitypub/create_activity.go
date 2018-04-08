package activitypub

import "time"

// CreateActivity is the type for a create activity message
type CreateActivity struct {
	Activity  *Create
	Published time.Time
	To        *ObjectOrLink
	CC        *ObjectOrLink
}

// CreateActivityNew initializes a new CreateActivity message
func CreateActivityNew(id ObjectID, a ObjectOrLink, o ObjectOrLink) CreateActivity {
	act := CreateNew(id, o)

	ok := false
	if a != nil {
		typ := a.Object().Type
		if typ == ApplicationType {
			act.Actor, ok = a.(Application)
		}
		if typ == GroupType {
			act.Actor, ok = a.(Group)
		}
		if typ == OrganizationType {
			act.Actor, ok = a.(Organization)
		}
		if typ == PersonType {
			act.Actor, ok = a.(Person)
		}
		if typ == ServiceType {
			act.Actor, ok = a.(Service)
		}
		if !ok {
			act.Actor, ok = a.(Actor)
		}
	}

	c := CreateActivity{
		Activity:  act,
		Published: time.Now(),
	}

	return c
}
