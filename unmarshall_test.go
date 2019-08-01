package activitypub

import (
	as "github.com/go-ap/activitystreams"
	"reflect"
	"testing"
)

type testPairs map[as.ActivityVocabularyType]reflect.Type

var objectPtrType = reflect.TypeOf(new(*Object)).Elem()
var actorPtrType = reflect.TypeOf(new(*actor)).Elem()
var applicationPtrType = reflect.TypeOf(new(*Application)).Elem()
var servicePtrType = reflect.TypeOf(new(*Service)).Elem()
var personPtrType = reflect.TypeOf(new(*Person)).Elem()
var groupPtrType = reflect.TypeOf(new(*Group)).Elem()
var organizationPtrType = reflect.TypeOf(new(*Organization)).Elem()

var tests = testPairs{
	as.ObjectType:       objectPtrType,
	as.ActorType:        actorPtrType,
	as.ApplicationType:  applicationPtrType,
	as.ServiceType:      servicePtrType,
	as.PersonType:       personPtrType,
	as.GroupType:        groupPtrType,
	as.OrganizationType: organizationPtrType,
}

func TestJSONGetItemByType(t *testing.T) {
	for typ, test := range tests {
		t.Run(string(typ), func(t *testing.T) {
			v, err := JSONGetItemByType(typ)
			if err != nil {
				t.Error(err)
			}
			if reflect.TypeOf(v) != test {
				t.Errorf("Invalid type returned %T, expected %s", v, test.String())
			}
		})
	}
}

func TestJSONGetActorEndpoints(t *testing.T) {
	t.Skipf("TODO")
}
