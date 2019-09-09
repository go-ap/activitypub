package activitypub

import (
	"github.com/buger/jsonparser"
	as "github.com/go-ap/activitystreams"
)

func JSONGetItemByType(typ as.ActivityVocabularyType) (as.Item, error) {
	if as.ObjectTypes.Contains(typ) {
		return &Object{Parent: Parent{Type: typ}}, nil
	}
	if as.ActorTypes.Contains(typ) {
		return &actor{Parent: Parent{Type: typ}}, nil
	}
	return as.JSONGetItemByType(typ)
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
