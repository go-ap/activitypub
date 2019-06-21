package activitypub

import (
	as "github.com/go-ap/activitystreams"
	"testing"
)

func TestActorNew(t *testing.T) {
	var testValue = as.ObjectID("test")
	var testType = as.ApplicationType

	o := actorNew(testValue, testType)

	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != testType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, testType)
	}

	n := actorNew(testValue, "")
	if n.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", n.ID, testValue)
	}
	if n.Type != as.ActorType {
		t.Errorf("APObject Type '%v' different than expected '%v'", n.Type, as.ActorType)
	}
}

func TestPersonNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	o := PersonNew(testValue)
	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != as.PersonType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, as.PersonType)
	}
}

func TestApplicationNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	o := ApplicationNew(testValue)
	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != as.ApplicationType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, as.ApplicationType)
	}
}

func TestGroupNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	o := GroupNew(testValue)
	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != as.GroupType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, as.GroupType)
	}
}

func TestOrganizationNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	o := OrganizationNew(testValue)
	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != as.OrganizationType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, as.OrganizationType)
	}
}

func TestServiceNew(t *testing.T) {
	var testValue = as.ObjectID("test")

	o := ServiceNew(testValue)
	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != as.ServiceType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, as.ServiceType)
	}
}

func TestActor_UnmarshalJSON(t *testing.T) {
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

func TestToPerson(t *testing.T) {
	t.Skipf("TODO")
}
