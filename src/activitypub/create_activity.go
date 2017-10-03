package activitypub

import "time"

type CreateActivity struct {
	Activity  *Create
	Published time.Time
	To        *Actor
	CC        *Actor
}

func CreateActivityNew(id ObjectId, o *ObjectOrLink) *CreateActivity {
	c := CreateActivity{
		Activity:  CreateNew(id, o),
		Published: time.Now(),
	}

	return &c
}
