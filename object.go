package activitypub

import (
	"github.com/buger/jsonparser"
	as "github.com/go-ap/activitystreams"
)

// Source is intended to convey some sort of source from which the content markup was derived,
// as a form of provenance, or to support future editing by clients.
type Source struct {
	// Content
	Content as.NaturalLanguageValue `jsonld:"content"`
	// MediaType
	MediaType as.MimeType `jsonld:"mediaType"`
}

// Object
type Object struct {
	as.Object
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
	o.Object.UnmarshalJSON(data)
	o.Source = GetAPSource(data)
	return nil
}
