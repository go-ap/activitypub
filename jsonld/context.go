package jsonld

import (
	"encoding/json"
	"fmt"
	"strings"

	ap "github.com/mariusor/activitypub.go/activitypub"
)

// Ref basic type
type Ref string

// Context is the basic JSON-LD element. It is used to map terms to IRIs.
// Terms are case sensitive and any valid string that is not a reserved JSON-LD
// keyword can be used as a term.
type Context struct {
	URL      Ref        `jsonld:"@url"`
	Language ap.LangRef `jsonld:"@language,omitempty,collapsible"`
}

// Collapsible is an interface used by the JSON-LD marshaller to collapse a struct to one single value
type Collapsible interface {
	Collapse() []byte
}

// Ref returns a new Ref object based on Context URL
func (c *Context) Ref() Ref {
	return Ref(c.URL)
}

// Collapse returns the plain text collapsed value of the current Context object
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

// UnmarshalJSON tries to load the Context from the incoming json value
func (c *Context) UnmarshalJSON(data []byte) error {
	if data[0] == '"' {
		// a quoted string - loading it to c.URL
		if data[len(data)-1] != '"' {
			return fmt.Errorf("invalid string value when unmarshalling Context value")
		}
		c.URL = Ref(data[1 : len(data)-1])
	}
	if data[0] == '{' {
		// an object - trying to load it to the struct
		if data[len(data)-1] != '}' {
			return fmt.Errorf("invalid object value when unmarshalling Context value")
		}

		var s scanner
		s.reset()
		var d decodeState
		d.scan = s
		d.init(data)

		d.scan.reset()

		t := d.valueInterface()

		for key, value := range t.(map[string]interface{}) {
			switch strings.ToLower(key) {
			case "@url":
				fallthrough
			case "url":
				c.URL = Ref(value.(string))
			case "@language":
				fallthrough
			case "language":
				c.Language = ap.LangRef(value.(string))
			default:
				return fmt.Errorf("unkown Context field %q", key)
			}
		}
	}
	return nil
}
