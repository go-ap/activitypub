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
func CreateActivityNew(id ObjectID, o ObjectOrLink) *CreateActivity {
	c := CreateActivity{
		Activity:  CreateNew(id, o),
		Published: time.Now(),
	}

	return &c
}
