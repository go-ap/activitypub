package activitystreams

import (
	"strings"
)

type (
	// IRI is a Internationalized Resource Identifiers (IRIs) RFC3987
	IRI string
)

// String returns the String value of the IRI object
func (i IRI) String() string {
	return string(i)
}

// GetLink
func (i IRI) GetLink() IRI {
	return i
}

// UnmarshalJSON
func (i *IRI) UnmarshalJSON(s []byte) error {
	*i = IRI(strings.Trim(string(s), "\""))
	return nil
}

// IsObject
func (i IRI) GetID() *ObjectID {
	return nil
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
