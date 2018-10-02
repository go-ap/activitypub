package activitypub

import "time"

// UpdateActivity is the type for a Update activity message
type UpdateActivity struct {
	Activity  *Update    `jsonld:"activity"`
	Published time.Time  `jsonld:"published"`
	To        ObjectsArr `jsonld:"to,omitempty,collapsible"`
	CC        ObjectsArr `jsonld:"cc,omitempty,collapsible"`
}

// UpdateActivityNew initializes a new UpdateActivity message
func UpdateActivityNew(id ObjectID, a ObjectOrLink, o ObjectOrLink) UpdateActivity {
	act := UpdateNew(id, o)

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

// UnmarshalJSON
func (u *Update) UnmarshalJSON(data []byte) error {
	a := Activity(*u)
	err := a.UnmarshalJSON(data)

	*u = Update(a)

	return err
}
