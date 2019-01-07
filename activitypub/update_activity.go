package activitypub

import (
	as "github.com/go-ap/activitypub.go/activitystreams"
	"time"
)

// UpdateActivity is the type for a Update activity message
type UpdateActivity struct {
	Activity  *as.Update        `jsonld:"activity"`
	Published time.Time         `jsonld:"published"`
	To        as.ItemCollection `jsonld:"to,omitempty,collapsible"`
	CC        as.ItemCollection `jsonld:"cc,omitempty,collapsible"`
}

// UpdateActivityNew initializes a new UpdateActivity message
func UpdateActivityNew(id as.ObjectID, a as.Item, o as.Item) UpdateActivity {
	act := as.UpdateNew(id, o)

	if a != nil {
		if a.IsObject() {
			act.Actor = loadActorWithInboxObject(a, o)
		}
		if a.IsLink() {
			act.Actor = a
		}
	}

	act.RecipientsDeduplication()

	c := UpdateActivity{
		Activity:  act,
		Published: time.Now(),
	}

	return c
}
