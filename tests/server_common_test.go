package tests

// Common server tests...

import (
	"testing"
)

// Server: Fetching the inbox
// Try retrieving the actor's inbox of an actor.
// Server responds to GET request at inbox URL
func TestInboxGETRequest(t *testing.T) {
	desc := `
Server: Fetching the inbox
 Try retrieving the actor's inbox of an actor.

  Server responds to GET request at inbox URL
`
	t.Skip(desc)
}

// Server: Fetching the inbox
// Try retrieving the actor's inbox of an actor.
// inbox is an OrderedCollection
func TestInboxIsOrderedCollection(t *testing.T) {
	desc := `
Server: Fetching the inbox
 Try retrieving the actor's inbox of an actor.

  inbox is an OrderedCollection
`
	t.Skip(desc)
}

// Server: Fetching the inbox
// Try retrieving the actor's inbox of an actor.
// Server filters inbox content according to the requester's permission
func TestInboxFilteringBasedOnPermissions(t *testing.T) {
	desc := `
Server: Fetching the inbox
 Try retrieving the actor's inbox of an actor.

  Server filters inbox content according to the requester's permission
`
	t.Skip(desc)
}

/*
func Test_(t *testing.T) {
	desc := `
`
	t.Skip(desc)
}
*/
