package activitystreams

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

func TestActor_IsLink(t *testing.T) {
	m := ActorNew("test", ActorType)
	if m.IsLink() {
		t.Errorf("%#v should not be a valid Link", m.Type)
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
	if reflect.DeepEqual(ObjectID(""), m.GetID()) {
		t.Errorf("%#v should not be an empty activity pub object", m.GetID())
	}
}

func TestActor_Type(t *testing.T) {
	m := ActorNew("test", ActorType)
	if m.GetType() != ActorType {
		t.Errorf("%#v should be an empty Link object", m.GetType())
	}
}

func TestPerson_IsLink(t *testing.T) {
	m := PersonNew("test")
	if m.IsLink() {
		t.Errorf("%T should not be a valid Link", m)
	}
}

func TestPerson_IsObject(t *testing.T) {
	m := PersonNew("test")
	if !m.IsObject() {
		t.Errorf("%T should be a valid object", m)
	}
}

func TestActor_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestActor_GetActor(t *testing.T) {
	t.Skipf("TODO")
}

func TestActor_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestActor_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestActor_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestApplication_GetActor(t *testing.T) {
	t.Skipf("TODO")
}

func TestApplication_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestApplication_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestApplication_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestApplication_IsLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestApplication_IsObject(t *testing.T) {
	t.Skipf("TODO")
}

func TestGroup_GetActor(t *testing.T) {
	t.Skipf("TODO")
}

func TestGroup_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestGroup_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestGroup_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestGroup_IsLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestGroup_IsObject(t *testing.T) {
	t.Skipf("TODO")
}

func TestOrganization_GetActor(t *testing.T) {
	t.Skipf("TODO")
}

func TestOrganization_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestOrganization_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestOrganization_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestOrganization_IsLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestOrganization_IsObject(t *testing.T) {
	t.Skipf("TODO")
}

func TestPerson_GetActor(t *testing.T) {
	t.Skipf("TODO")
}

func TestPerson_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestPerson_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestPerson_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func validateEmptyPerson(p Person, t *testing.T) {
	if p.ID != "" {
		t.Errorf("Unmarshalled object %T should have empty ID, received %q", p, p.ID)
	}
	if p.Type != "" {
		t.Errorf("Unmarshalled object %T should have empty Type, received %q", p, p.Type)
	}
	if p.AttributedTo != nil {
		t.Errorf("Unmarshalled object %T should have empty AttributedTo, received %q", p, p.AttributedTo)
	}
	if len(p.Name) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Name, received %q", p, p.Name)
	}
	if len(p.Summary) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Summary, received %q", p, p.Summary)
	}
	if len(p.Content) != 0 {
		t.Errorf("Unmarshalled object %T should have empty Content, received %q", p, p.Content)
	}
	if p.URL != nil {
		t.Errorf("Unmarshalled object %T should have empty URL, received %v", p, p.URL)
	}
	if !p.Published.IsZero() {
		t.Errorf("Unmarshalled object %T should have empty Published, received %q", p, p.Published)
	}
	if !p.StartTime.IsZero() {
		t.Errorf("Unmarshalled object %T  should have empty StartTime, received %q", p, p.StartTime)
	}
	if !p.Updated.IsZero() {
		t.Errorf("Unmarshalled object %T  should have empty Updated, received %q", p, p.Updated)
	}
}

func TestPerson_UnmarshalJSON(t *testing.T) {
	p := Person{}

	dataEmpty := []byte("{}")
	p.UnmarshalJSON(dataEmpty)
	validateEmptyPerson(p, t)
}

func TestApplication_UnmarshalJSON(t *testing.T) {
	a := Application{}

	dataEmpty := []byte("{}")
	a.UnmarshalJSON(dataEmpty)
	validateEmptyPerson(Person(a), t)
}

func TestGroup_UnmarshalJSON(t *testing.T) {
	g := Group{}

	dataEmpty := []byte("{}")
	g.UnmarshalJSON(dataEmpty)
	validateEmptyPerson(Person(g), t)
}

func TestOrganization_UnmarshalJSON(t *testing.T) {
	o := Organization{}

	dataEmpty := []byte("{}")
	o.UnmarshalJSON(dataEmpty)
	validateEmptyPerson(Person(o), t)
}

func TestService_UnmarshalJSON(t *testing.T) {
	s := Service{}

	dataEmpty := []byte("{}")
	s.UnmarshalJSON(dataEmpty)
	validateEmptyPerson(Person(s), t)
}

func TestService_GetActor(t *testing.T) {
	t.Skipf("TODO")
}

func TestService_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestService_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestService_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestService_IsLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestService_IsObject(t *testing.T) {
	t.Skipf("TODO")
}
