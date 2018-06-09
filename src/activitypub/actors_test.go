package activitypub

import (
	"reflect"
	"testing"
)

func TestActorNew(t *testing.T) {
	var testValue = ObjectID("test")
	var testType = ApplicationType

	o := ActorNew(testValue, testType)

	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != testType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, testType)
	}

	n := ActorNew(testValue, "")
	if n.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", n.ID, testValue)
	}
	if n.Type != ActorType {
		t.Errorf("APObject Type '%v' different than expected '%v'", n.Type, ActorType)
	}
}

func TestPersonNew(t *testing.T) {
	var testValue = ObjectID("test")

	o := PersonNew(testValue)
	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != PersonType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, PersonType)
	}
}

func TestApplicationNew(t *testing.T) {
	var testValue = ObjectID("test")

	o := ApplicationNew(testValue)
	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != ApplicationType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, ApplicationType)
	}
}

func TestGroupNew(t *testing.T) {
	var testValue = ObjectID("test")

	o := GroupNew(testValue)
	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != GroupType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, GroupType)
	}
}

func TestOrganizationNew(t *testing.T) {
	var testValue = ObjectID("test")

	o := OrganizationNew(testValue)
	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != OrganizationType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, OrganizationType)
	}
}

func TestServiceNew(t *testing.T) {
	var testValue = ObjectID("test")

	o := ServiceNew(testValue)
	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != ServiceType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, ServiceType)
	}
}

func TestValidActorType(t *testing.T) {
	var invalidType ActivityVocabularyType = "RandomType"

	if ValidActorType(invalidType) {
		t.Errorf("APObject Type '%v' should not be valid", invalidType)
	}
	for _, validType := range validActorTypes {
		if !ValidActorType(validType) {
			t.Errorf("APObject Type '%v' should be valid", validType)
		}
	}
}

func TestActor_IsLink(t *testing.T) {
	m := ActorNew("test", ActorType)
	if m.IsLink() {
		t.Errorf("%#v should not be a valid GetLink", m.Type)
	}
}

func TestActor_IsObject(t *testing.T) {
	m := ActorNew("test", ActorType)
	if !m.IsObject() {
		t.Errorf("%#v should be a valid object", m.Type)
	}
}

func TestActor_Object(t *testing.T) {
	m := ActorNew("test", ActorType)
	if reflect.DeepEqual(BaseObject{}, m.GetObject()) {
		t.Errorf("%#v should not be an empty activity pub object", m.GetObject())
	}
}

func TestActor_Link(t *testing.T) {
	m := ActorNew("test", ActorType)
	if !reflect.DeepEqual(Link{}, m.Link()) {
		t.Errorf("%#v should be an empty GetLink object", m.Link())
	}
}

func TestPerson_IsLink(t *testing.T) {
	m := PersonNew("test")
	if m.IsLink() {
		t.Errorf("%#v should not be a valid GetLink", m.Type)
	}
}

func TestPerson_IsObject(t *testing.T) {
	m := PersonNew("test")
	if !m.IsObject() {
		t.Errorf("%#v should be a valid object", m.Type)
	}
}

func TestPerson_Object(t *testing.T) {
	m := PersonNew("test")
	if reflect.DeepEqual(BaseObject{}, m.GetObject()) {
		t.Errorf("%#v should not be an empty activity pub object", m.GetObject())
	}
}

func TestPerson_Link(t *testing.T) {
	m := PersonNew("test")
	if !reflect.DeepEqual(Link{}, m.Link()) {
		t.Errorf("%#v should be an empty GetLink object", m.Link())
	}
}
