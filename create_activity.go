package activitypub

import "time"

type CreateActivity struct {
	ActivityObject
	Published time.Time
	To Actor
	CC Actor
}