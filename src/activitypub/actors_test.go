package activitypub

import (
	"testing"
)

func TestActorNew(t *testing.T) {
	var testValue = ObjectId("test")
	var testType = ApplicationType

	o := ActorNew(testValue, testType)

	if o.Id != testValue {
		t.Errorf("Object Id '%v' different than expected '%v'", o.Id, testValue)
	}
	if o.Type != testType {
		t.Errorf("Object Type '%v' different than expected '%v'", o.Type, testType)
	}

	n := ActorNew(testValue, "")
	if n.Id != testValue {
		t.Errorf("Object Id '%v' different than expected '%v'", n.Id, testValue)
	}
	if n.Type != ActorType {
		t.Errorf("Object Type '%v' different than expected '%v'", n.Type, ActorType)
	}
}

func TestPersonNew(t *testing.T) {
	var testValue = ObjectId("test")

	o := PersonNew(testValue)
	if o.Id != testValue {
		t.Errorf("Object Id '%v' different than expected '%v'", o.Id, testValue)
	}
	if o.Type != PersonType {
		t.Errorf("Object Type '%v' different than expected '%v'", o.Type, PersonType)
	}
}

func TestApplicationNew(t *testing.T) {
	var testValue = ObjectId("test")

	o := ApplicationNew(testValue)
	if o.Id != testValue {
		t.Errorf("Object Id '%v' different than expected '%v'", o.Id, testValue)
	}
	if o.Type != ApplicationType {
		t.Errorf("Object Type '%v' different than expected '%v'", o.Type, ApplicationType)
	}
}

func TestGroupNew(t *testing.T) {
	var testValue = ObjectId("test")

	o := GroupNew(testValue)
	if o.Id != testValue {
		t.Errorf("Object Id '%v' different than expected '%v'", o.Id, testValue)
	}
	if o.Type != GroupType {
		t.Errorf("Object Type '%v' different than expected '%v'", o.Type, GroupType)
	}
}

func TestOrganizationNew(t *testing.T) {
	var testValue = ObjectId("test")

	o := OrganizationNew(testValue)
	if o.Id != testValue {
		t.Errorf("Object Id '%v' different than expected '%v'", o.Id, testValue)
	}
	if o.Type != OrganizationType {
		t.Errorf("Object Type '%v' different than expected '%v'", o.Type, OrganizationType)
	}
}

func TestServiceNew(t *testing.T) {
	var testValue = ObjectId("test")

	o := ServiceNew(testValue)
	if o.Id != testValue {
		t.Errorf("Object Id '%v' different than expected '%v'", o.Id, testValue)
	}
	if o.Type != ServiceType {
		t.Errorf("Object Type '%v' different than expected '%v'", o.Type, ServiceType)
	}
}
