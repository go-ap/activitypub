package activitypub

import "time"

// CreateActivity is the type for a create activity message
type CreateActivity struct {
	Activity  *Create
	Published time.Time
	To        *Actor
	CC        *Actor
}

// CreateActivityNew initializes a new CreateActivity message
func CreateActivityNew(id ObjectID, a ObjectOrLink, o ObjectOrLink) *CreateActivity {
	act := CreateNew(id, o)

	if a != nil {
		act.Actor = Actor(a.(Person))
	}

	c := CreateActivity{
		Activity:  act,
		Published: time.Now(),
	}

	return &c
}
