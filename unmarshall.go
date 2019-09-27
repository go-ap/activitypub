package activitypub

import (
	"github.com/buger/jsonparser"
	as "github.com/go-ap/activitystreams"
)

func JSONGetItemByType(typ as.ActivityVocabularyType) (as.Item, error) {
	obTyp := as.ActivityVocabularyTypes{as.ObjectType}
	if as.ObjectTypes.Contains(typ) || obTyp.Contains(typ) {
		return &Object{Parent: as.Object{Type: typ}}, nil
	}
	actTyp := as.ActivityVocabularyTypes{as.ActorType}
	if as.ActorTypes.Contains(typ) || actTyp.Contains(typ) {
		return &actor{Parent: Parent{Parent: as.Object{Type: typ}}}, nil
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
