package activitypub

import (
	"fmt"
	"github.com/buger/jsonparser"
	"net/url"
	"path"
	"strings"
)

const (
	// ActivityBaseURI the URI for the activity streams namespace
	ActivityBaseURI = IRI("https://www.w3.org/ns/activitystreams")
	// SecurityContextURI the URI for the secruity namespace (for an Actor's PublicKey)
	SecurityContextURI = IRI("https://w3id.org/security/v1")
	// PublicNs is the reference to the Public entity in the Acitivystreams namespace
	PublicNS = ActivityBaseURI + "#Public"
)

var JsonLDContext = []IRI{
	ActivityBaseURI,
	SecurityContextURI,
}

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

// MarshalJSON
func (i IRI) MarshalJSON() ([]byte, error) {
	if i == "" {
		return nil, nil
	}
	b := make([]byte, 0)
	write(&b, '"')
	writeS(&b, i.String())
	write(&b, '"')
	return b, nil
}

// AddPath concatenates el elements as a path to i
func (i IRI) AddPath(el ...string) IRI {
	return IRI(fmt.Sprintf("%s/%s", i, path.Join(el...)))
}

// GetID
func (i IRI) GetID() ID {
	return i
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

// IsCollection returns false for IRI objects
func (i IRI) IsCollection() bool {
	return false
}

// FlattenToIRI checks if Item can be flatten to an IRI and returns it if so
func FlattenToIRI(it Item) Item {
	if it != nil && it.IsObject() && len(it.GetLink()) > 0 {
		return it.GetLink()
	}
	return it
}

func (i IRIs) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	if len(i) == 0 {
		return nil, nil
	}
	writeCommaIfNotEmpty := func(notEmpty bool) {
		if notEmpty {
			writeS(&b, ",")
		}
	}
	write(&b, '[')
	for k, iri := range i {
		writeCommaIfNotEmpty(k > 0)
		write(&b, '"')
		writeS(&b, iri.String())
		write(&b, '"')
	}
	write(&b, ']')
	return b, nil
}

func (i *IRIs) UnmarshalJSON(data []byte) error {
	if i == nil {
		return nil
	}
	value, typ, _, err := jsonparser.Get(data)
	if err != nil {
		return err
	}
	switch typ {
	case jsonparser.String:
		if iri, ok := asIRI(value); ok {
			*i = append(*i, iri)
		}
	case jsonparser.Array:
		jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			if iri, ok := asIRI(value); ok {
				*i = append(*i, iri)
			}
		})
	}
	return nil
}

// Contains verifies if IRIs array contains the received one
func (i IRIs) Contains(r IRI) bool {
	if len(i) == 0 {
		return false
	}
	for _, iri := range i {
		if r.Equals(iri, false) {
			return true
		}
	}
	return false
}

func validURL(u *url.URL) bool {
	return len(u.Scheme) > 0 && len(u.Host) > 0
}

// Equals verifies if our receiver IRI is equals with the "with" IRI
// It ignores the protocol
// It tries to use the URL representation if possible and fallback to string comparison if unable to convert
// In URL representation checks hostname in a case insensitive way and the path and
func (i IRI) Equals(with IRI, checkScheme bool) bool {
	u, e := i.URL()
	uw, ew := with.URL()
	if e != nil || ew != nil || !validURL(u) || !validURL(uw) {
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
	if path.Clean(u.Path) != path.Clean(uw.Path) {
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

func (i IRI) Contains(what IRI, checkScheme bool) bool {
	u, e := i.URL()
	uw, ew := what.URL()
	if e != nil || ew != nil {
		return strings.Contains(i.String(), what.String())
	}
	if checkScheme {
		if u.Scheme != uw.Scheme {
			return false
		}
	}
	if u.Host != uw.Host {
		return false
	}
	p := u.Path
	if p != "" {
		p = path.Clean(p)
	}
	pw := uw.Path
	if pw != "" {
		pw = path.Clean(pw)
	}
	return strings.Contains(p, pw)
}

func (i IRI) ItemsMatch(col ...Item) bool {
	for _, it := range col {
		if match := it.GetLink().Contains(i, false); !match {
			return false
		}
	}
	return true
}
