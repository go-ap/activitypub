package client

import (
	"strings"
	"testing"

	as "github.com/go-ap/activitystreams"
)

func TestNewClient(t *testing.T) {
	c := client{}

	if c != *NewClient() {
		t.Errorf("NewClient didn't return a valid client")
	}
}

func TestErr_Error(t *testing.T) {
	e := err{
		msg: "test",
		iri: as.IRI(""),
	}

	if len(e.Error()) == 0 {
		t.Errorf("error message should not be empty")
	}
	if !strings.Contains(e.Error(), "test") {
		t.Errorf("error message should contain the 'test' string")
	}
}

func TestClient_LoadIRI(t *testing.T) {
	empty := as.IRI("")
	c := NewClient()

	var err error
	_, err = c.LoadIRI(empty)
	if err == nil {
		t.Errorf("LoadIRI should have failed when using empty IRI value")
	} else {
		t.Logf("Valid error received: %s", err)
	}

	inv := as.IRI("example.com")
	_, err = c.LoadIRI(inv)
	if err == nil {
		t.Errorf("LoadIRI should have failed when using invalid http url")
	} else {
		t.Logf("Valid error received: %s", err)
	}
}

func TestClient_Get(t *testing.T) {

}

func TestClient_Head(t *testing.T) {

}

func TestClient_Post(t *testing.T) {

}

func TestClient_Put(t *testing.T) {

}

func TestClient_Delete(t *testing.T) {

}
