package activitypub

import (
	"github.com/buger/jsonparser"
	as "github.com/go-ap/activitystreams"
)

func JSONGetItemByType(typ as.ActivityVocabularyType) (as.Item, error) {
	var ret as.Item
	var err error

	switch typ {
	case as.ObjectType:
		o := &Object{}
		o.Type = typ
		ret = o
	case as.ActorType:
		o := &actor{}
		o.Type = typ
		ret = o
	case as.ApplicationType:
		a := &Application{}
		a.Type = typ
		ret = a
	case as.GroupType:
		g := &Group{}
		g.Type = typ
		ret = g
	case as.OrganizationType:
		o := &Organization{}
		o.Type = typ
		ret = o
	case as.PersonType:
		p := &Person{}
		p.Type = typ
		ret = p
	case as.ServiceType:
		s := &Service{}
		s.Type = typ
		ret = s
	case "":
		// when no type is available use a plain Object
		ret = &Object{}
	default:
		return as.JSONGetItemByType(typ)
	}
	return ret, err
}

func JSONGetActorEndpoints(data []byte, prop string) *Endpoints {
	str, _ := jsonparser.GetUnsafeString(data, prop)

	var e *Endpoints
	if len(str) > 0 {
		e = &Endpoints{}
		e.UnmarshalJSON([]byte(str))
	}

	return e
}
