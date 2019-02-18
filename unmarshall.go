package activitypub

import (
	as "github.com/go-ap/activitystreams"
)

func JSONGetItemByType(typ as.ActivityVocabularyType) (as.Item, error) {
	var ret as.Item
	var err error

	switch typ {
	case as.ObjectType:
		o := Object{}
		o.Type = typ
		ret = &o
	case as.ActorType:
		ret = &Object{}
		o := ret.(*Object)
		o.Type = typ
	case as.ApplicationType:
		ret = &Application{}
		o := ret.(*Application)
		o.Type = typ
	case as.GroupType:
		ret = &Group{}
		o := ret.(*Group)
		o.Type = typ
	case as.OrganizationType:
		ret = &Organization{}
		o := ret.(*Organization)
		o.Type = typ
	case as.PersonType:
		ret = &Person{}
		o := ret.(*Person)
		o.Type = typ
	case as.ServiceType:
		ret = &Service{}
		o := ret.(*Service)
		o.Type = typ
	case "":
		// when no type is available use a plain Object
		ret = &Object{}
	default:
		return as.JSONGetItemByType(typ)
	}
	return ret, err
}
