package activitypub

import "time"

// LikeActivity is the type for a create activity message
type LikeActivity struct {
	Activity  *Like          `jsonld:"activity"`
	Published time.Time      `jsonld:"published"`
	To        ItemCollection `jsonld:"to,omitempty,collapsible"`
	CC        ItemCollection `jsonld:"cc,omitempty,collapsible"`
}

// DislikeActivity is the type for a create activity message
type DislikeActivity struct {
	Activity  *Dislike       `jsonld:"activity"`
	Published time.Time      `jsonld:"published"`
	To        ItemCollection `jsonld:"to,omitempty,collapsible"`
	CC        ItemCollection `jsonld:"cc,omitempty,collapsible"`
}

func loadActorWithLikedObject(a Item, o Item) Item {
	typ := a.GetType()
	switch typ {
	case ApplicationType:
		var app Application
		app, _ = a.(Application)
		var liked *Liked
		if app.Liked == nil {
			liked = LikedNew()
		} else {
			liked = app.Liked.(*Liked)
		}
		liked.Append(o)
		app.Liked = liked
		return app
	case GroupType:
		var grp Group
		grp, _ = a.(Group)
		var liked *Liked
		if grp.Liked == nil {
			liked = LikedNew()
		} else {
			liked = grp.Liked.(*Liked)
		}
		liked.Append(o)
		grp.Liked = liked
		return grp
	case OrganizationType:
		var org Organization
		org, _ = a.(Organization)
		var liked *Liked
		if org.Liked == nil {
			liked = LikedNew()
		} else {
			liked = org.Liked.(*Liked)
		}
		liked.Append(o)
		org.Liked = liked
		return org
	case PersonType:
		var pers Person
		pers, _ = a.(Person)
		var liked *Liked
		if pers.Liked == nil {
			liked = LikedNew()
		} else {
			liked = pers.Liked.(*Liked)
		}
		liked.Append(o)
		pers.Liked = liked
		return pers
	case ServiceType:
		var serv Service
		serv, _ = a.(Service)
		var liked *Liked
		if serv.Liked == nil {
			liked = LikedNew()
		} else {
			liked = serv.Liked.(*Liked)
		}
		liked.Append(o)
		serv.Liked = liked
		return serv
	default:
		actor, _ := a.(Actor)
		var liked *Liked
		if actor.Liked == nil {
			liked = LikedNew()
		} else {
			liked = actor.Liked.(*Liked)
		}
		liked.Append(o)
		actor.Liked = liked
		return actor
	}
}

// LikeActivityNew initializes a new LikeActivity message
func LikeActivityNew(id ObjectID, a Item, o Item) LikeActivity {
	act := LikeNew(id, o)

	if a != nil {
		if a.IsObject() {
			act.Actor = loadActorWithLikedObject(a, o)
		}
		if a.IsLink() {
			act.Actor = a
		}
	}

	act.RecipientsDeduplication()

	c := LikeActivity{
		Activity:  act,
		Published: time.Now(),
	}

	return c
}

// DislikeActivityNew initializes a new LikeActivity message
func DislikeActivityNew(id ObjectID, a Item, o Item) DislikeActivity {
	act := DislikeNew(id, o)

	if a != nil {
		if a.IsObject() {
			act.Actor = loadActorWithLikedObject(a, o)
		}
		if a.IsLink() {
			act.Actor = a
		}
	}

	act.RecipientsDeduplication()

	d := DislikeActivity{
		Activity:  act,
		Published: time.Now(),
	}

	return d
}

// UnmarshalJSON
func (l *Like) UnmarshalJSON(data []byte) error {
	a := Activity(*l)
	err := a.UnmarshalJSON(data)

	*l = Like(a)

	return err
}

// UnmarshalJSON
func (d *Dislike) UnmarshalJSON(data []byte) error {
	a := Activity(*d)
	err := a.UnmarshalJSON(data)

	*d = Dislike(a)

	return err
}
