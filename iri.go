package activitystreams

import (
	"net/url"
	"strings"
)

const PublicNS = IRI("https://www.w3.org/ns/activitystreams#Public")

type (
	// IRI is a Internationalized Resource Identifiers (IRIs) RFC3987
	IRI  string
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

// GetID
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
	if it != nil && it.IsObject() && len(it.GetLink()) > 0 {
		return it.GetLink()
	}
	return it
}

// Contains verifies if IRIs array contains the received one
func (i IRIs) Contains(r IRI) bool {
	if len(i) == 0 {
		return false
	}
	for _, iri := range i {
		if strings.ToLower(r.String()) == strings.ToLower(iri.String()) {
			return true
		}
	}
	return false
}

// Equals verifies if our receiver IRI is equals with the "with" IRI
// It ignores the protocol
// It tries to use the URL representation if possible and fallback to string comparison if unable to convert
// In URL representation checks hostname in a case insensitive way and the path and
func (i IRI) Equals(with IRI, checkScheme bool) bool {
	u, e := i.URL()
	uw, ew := with.URL()
	if e != nil || ew != nil {
		return strings.ToLower(i.String()) == strings.ToLower(with.String())
	}
	if checkScheme {
		if strings.ToLower(u.Scheme) != strings.ToLower(uw.Scheme) {
			return false
		}
	}
	if strings.ToLower(u.Host) != strings.ToLower(uw.Host) {
		return false
	}
	if u.Path != uw.Path {
		return false
	}
	uq := u.Query()
	uwq := uw.Query()
	if len(uq) != len(uwq) {
		return false
	}
	for k, uqv := range uq {
		uwqv, ok := uwq[k]
		if !ok {
			return false
		}
		if len(uqv) != len(uwqv) {
			return false
		}
		for _, uqvv := range uqv {
			eq := false
			for _, uwqvv := range uwqv {
				if uwqvv == uqvv {
					eq = true
					continue
				}
			}
			if !eq {
				return false
			}
		}
	}
	return true
}
