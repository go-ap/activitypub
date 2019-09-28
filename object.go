package activitypub

import (
	"errors"
	"github.com/buger/jsonparser"
	as "github.com/go-ap/activitystreams"
)

// Source is intended to convey some sort of source from which the content markup was derived,
// as a form of provenance, or to support future editing by clients.
type Source struct {
	// Content
	Content as.NaturalLanguageValues `jsonld:"content"`
	// MediaType
	MediaType as.MimeType `jsonld:"mediaType"`
}

type Parent = Object

// Object
type Object struct {
	as.Parent
	// This is a list of all Like activities with this object as the object property, added as a side effect.
	// The likes collection MUST be either an OrderedCollection or a Collection and MAY be filtered on privileges
	// of an authenticated user or as appropriate when no authentication is given.
	Likes as.Item `jsonld:"likes,omitempty"`
	// This is a list of all Announce activities with this object as the object property, added as a side effect.
	// The shares collection MUST be either an OrderedCollection or a Collection and MAY be filtered on privileges
	// of an authenticated user or as appropriate when no authentication is given.
	Shares as.Item `jsonld:"shares,omitempty"`
	// Source property is intended to convey some sort of source from which the content markup was derived,
	// as a form of provenance, or to support future editing by clients.
	// In general, clients do the conversion from source to content, not the other way around.
	Source Source `jsonld:"source,omitempty"`
}

// GetAPSource
func GetAPSource(data []byte) Source {
	s := Source{}

	if contBytes, _, _, err := jsonparser.Get(data, "source", "content"); err == nil {
		s.Content.UnmarshalJSON(contBytes)
	}
	if mimeBytes, _, _, err := jsonparser.Get(data, "source", "mediaType"); err == nil {
		s.MediaType.UnmarshalJSON(mimeBytes)
	}

	return s
}

// UnmarshalJSON
func (s *Source) UnmarshalJSON(data []byte) error {
	*s = GetAPSource(data)
	return nil
}

// UnmarshalJSON
func (o *Object) UnmarshalJSON(data []byte) error {
	if as.ItemTyperFunc == nil {
		as.ItemTyperFunc = JSONGetItemByType
	}
	o.Parent.UnmarshalJSON(data)
	o.Likes = as.JSONGetItem(data, "likes")
	o.Shares = as.JSONGetItem(data, "shares")
	o.Source = GetAPSource(data)
	return nil
}

// ToObject
func ToObject(it as.Item) (*Object, error) {
	switch i := it.(type) {
	case *as.Object:
		return &Object{Parent: *i}, nil
	case as.Object:
		return &Object{Parent: i}, nil
	case *Object:
		return i, nil
	case Object:
		return &i, nil
	}
	return nil, errors.New("unable to convert object")
}
