package jsonld

import (
	"encoding/json"

	"activitypub"
)

// Ref basic type
type Ref string

// Context is the basic JSON-LD element. It is used to map terms to IRIs.
// Terms are case sensitive and any valid string that is not a reserved JSON-LD
// keyword can be used as a term.
type Context struct {
	URL      Ref                              `jsonld:"@url"`
	Language activitypub.NaturalLanguageValue `jsonld:"@language,omitempty,collapsible"`
}

type Collapsible interface {
	Collapse() []byte
}

// Ref returns a new Ref object based on Context URL
func (c *Context) Ref() Ref {
	return Ref(c.URL)
}

func (c *Context) Collapse() []byte {
	return []byte(c.URL)
}

// MarshalText basic stringify function
func (r *Ref) MarshalText() ([]byte, error) {
	return []byte(*r), nil
}

// MarshalJSON returns the JSON document represented by the current Context
func (c *Context) MarshalJSON() ([]byte, error) {
	a := reflectToJSONValue(c)
	if a.isScalar {
		return json.Marshal(a.scalar)
	}
	return json.Marshal(a.object)
}
