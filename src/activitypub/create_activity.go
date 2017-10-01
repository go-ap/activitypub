package activitypub

import "time"

type CreateActivity struct {
	Activity  Create
	Published time.Time
	To        Actor
	CC        Actor
}
