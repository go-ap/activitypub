package activitypub

import (
	as "github.com/go-ap/activitypub.go/activitystreams"
)

// ValidationErrors is an aggregated error interface that allows
// a Validator implementation to return all possible errors.
type ValidationErrors interface {
	error
	Errors() []error
	Add(error)
}

// Validator is the interface that needs to be implemented by objects that
// provide a validation mechanism for incoming ActivityPub Objects or IRIs
// against an external set of rules.
type Validator interface {
	Validate(receiver as.IRI, incoming as.Item) (bool, ValidationErrors)
}

func (v defaultValidator) Validate(receiver as.IRI, incoming as.Item) (bool, ValidationErrors) {
	return true, nil
}

type defaultValidator struct{}
