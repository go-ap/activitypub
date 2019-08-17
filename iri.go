package activitystreams

import (
	"net/url"
	"strings"
)

type (
	// IRI is a Internationalized Resource Identifiers (IRIs) RFC3987
	IRI string
	IRIs []IRI
)

// String returns the String value of the IRI object
func (i IRI) String() string {
	return string(i)
}

// GetLink
func (i IRI) GetLink() IRI {
	return i
}

// URL
func (i IRI) URL() (*url.URL, error) {
	return url.Parse(i.String())
}

// UnmarshalJSON
func (i *IRI) UnmarshalJSON(s []byte) error {
	*i = IRI(strings.Trim(string(s), "\""))
	return nil
}

// IsObject
func (i IRI) GetID() *ObjectID {
	o := ObjectID(i)
	return &o
}

// GetType
func (i IRI) GetType() ActivityVocabularyType {
	return LinkType
}

// IsLink
func (i IRI) IsLink() bool {
	return true
}

// IsObject
func (i IRI) IsObject() bool {
	return false
}

// FlattenToIRI checks if Item can be flatten to an IRI and returns it if so
func FlattenToIRI(it Item) Item {
	if it!= nil && it.IsObject() && len(it.GetLink()) > 0 {
		return it.GetLink()
	}
	return it
}

// Contains verifies if IRIs array contains the received one
func(i IRIs) Contains(r IRI) bool {
	if len(i) == 0 {
		return true
	}
	for _, iri := range i {
		if strings.ToLower(r.String()) == strings.ToLower(iri.String()) {
			return true
		}
	}
	return false
}
