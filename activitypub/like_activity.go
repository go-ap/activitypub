package activitypub

import "time"

// LikeActivity is the type for a create activity message
type LikeActivity struct {
	Activity  *Like      `jsonld:"activity"`
	Published time.Time  `jsonld:"published"`
	To        ObjectsArr `jsonld:"to,omitempty,collapsible"`
	CC        ObjectsArr `jsonld:"cc,omitempty,collapsible"`
}

func loadActorWithLikedObject(a ObjectOrLink, o ObjectOrLink) ObjectOrLink {
	typ := a.GetType()
	switch typ {
	case ApplicationType:
		var app Application
		app, _ = a.(Application)
		if app.Liked == nil {
			app.Liked = LikedNew()
		}
		app.Liked.Append(o)
		return app
	case GroupType:
		var grp Group
		grp, _ = a.(Group)
		if grp.Liked == nil {
			grp.Liked = LikedNew()
		}
		grp.Liked.Append(o)
		return grp
	case OrganizationType:
		var org Organization
		org, _ = a.(Organization)
		if org.Liked == nil {
			org.Liked = LikedNew()
		}
		org.Liked.Append(o)
		return org
	case PersonType:
		var pers Person
		pers, _ = a.(Person)
		if pers.Liked == nil {
			pers.Liked = LikedNew()
		}
		pers.Liked.Append(o)
		return pers
	case ServiceType:
		var serv Service
		serv, _ = a.(Service)
		serv.Liked.Append(o)
		return serv
	default:
		actor, _ := a.(Actor)
		if actor.Liked == nil {
			actor.Liked = LikedNew()
		}
		actor.Liked.Append(o)
		return actor
	}
}

// LikeActivityNew initializes a new LikeActivity message
func LikeActivityNew(id ObjectID, a ObjectOrLink, o ObjectOrLink) LikeActivity {
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

// UnmarshalJSON
func (l *Like) UnmarshalJSON(data []byte) error {
	a := Activity(*l)
	err := a.UnmarshalJSON(data)

	*l = Like(a)

	return err
}
