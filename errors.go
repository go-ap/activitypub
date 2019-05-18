package activitypub

import (
	"fmt"
	as "github.com/go-ap/activitystreams"
)

type causer interface {
	// Cause returns the parent error
	Cause() error
}

// Error is an local error type
// It should contain the ActivityPub object that generated the error
type Error struct {
	on     as.Item
	msg    string
	parent error
}

// Error implements the errors interface
func (e Error) Error() string {
	return fmt.Sprintf("%s %s", e.msg, *e.on.GetID())
}

// Cause returns the parent error
func (e Error) Cause() error {
	return e.parent
}

func (e Error) Is(err error) bool {
	for {
		if err == e {
			return true
		}
		if e, ok := err.(causer); ok {
			err = e.Cause()
		} else {
			return false
		}
	}
}
