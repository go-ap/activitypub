package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/valyala/fastjson"
)

const (
	// ActivityBaseURI the URI for the ActivityStreams namespace
	ActivityBaseURI = IRI("https://www.w3.org/ns/activitystreams")
	// SecurityContextURI the URI for the security namespace (for an Actor's PublicKey)
	SecurityContextURI = IRI("https://w3id.org/security/v1")
	// PublicNS is the reference to the Public entity in the ActivityStreams namespace
	PublicNS = ActivityBaseURI + "#Public"
)

// JsonLDContext is a slice of IRIs that form the default context for the objects in the
// GoActivitypub vocabulary.
// It does not represent just the default ActivityStreams public namespace, but it also
// has the W3 Permanent Identifier Community Group's Security namespace, which appears
// in the Actor type objects, which contain public key related data.
var JsonLDContext = []IRI{
	ActivityBaseURI,
	SecurityContextURI,
}

type (
	// IRI is a Internationalized Resource Identifiers (IRIs) RFC3987
	IRI  string
	IRIs []IRI
)

func (i IRI) Format(s fmt.State, verb rune) {
	switch verb {
	case 's', 'v':
		u, _ := i.URL()
		u.RawQuery, _ = url.QueryUnescape(u.RawQuery)
		io.WriteString(s, u.String())
	}
}

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

// UnmarshalJSON decodes an incoming JSON document into the receiver object.
func (i *IRI) UnmarshalJSON(s []byte) error {
	*i = IRI(strings.Trim(string(s), "\""))
	return nil
}

// MarshalJSON encodes the receiver object to a JSON document.
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

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (i *IRI) UnmarshalBinary(data []byte) error {
	return i.GobDecode(data)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (i IRI) MarshalBinary() ([]byte, error) {
	return i.GobEncode()
}

// GobEncode
func (i IRI) GobEncode() ([]byte, error) {
	return []byte(i), nil
}

// GobEncode
func (i IRIs) GobEncode() ([]byte, error) {
	if len(i) == 0 {
		return []byte{}, nil
	}
	b := bytes.Buffer{}
	gg := gob.NewEncoder(&b)
	bb := make([][]byte, 0)
	for _, iri := range i {
		bb = append(bb, []byte(iri))
	}
	if err := gg.Encode(bb); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// GobDecode
func (i *IRI) GobDecode(data []byte) error {
	*i = IRI(data)
	return nil
}

func (i *IRIs) GobDecode(data []byte) error {
	if len(data) == 0 {
		// NOTE(marius): this behaviour diverges from vanilla gob package
		return nil
	}
	bb := make([][]byte, 0)
	if err := gob.NewDecoder(bytes.NewReader(data)).Decode(&bb); err != nil {
		return err
	}
	for _, b := range bb {
		*i = append(*i, IRI(b))
	}
	return nil
}

// AddPath concatenates el elements as a path to i
func (i IRI) AddPath(el ...string) IRI {
	iri := strings.TrimRight(i.String(), "/")
	return IRI(iri + filepath.Clean(filepath.Join("/", filepath.Join(el...))))
}

// GetID
func (i IRI) GetID() ID {
	return i
}

// GetType
func (i IRI) GetType() ActivityVocabularyType {
	return IRIType
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
	if !IsNil(it) && it.IsObject() && len(it.GetLink()) > 0 {
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
	p := fastjson.Parser{}
	val, err := p.ParseBytes(data)
	if err != nil {
		return err
	}
	switch val.Type() {
	case fastjson.TypeString:
		if iri, ok := asIRI(val); ok && len(iri) > 0 {
			*i = append(*i, iri)
		}
	case fastjson.TypeArray:
		for _, v := range val.GetArray() {
			if iri, ok := asIRI(v); ok && len(iri) > 0 {
				*i = append(*i, iri)
			}
		}
	}
	return nil
}

// GetID returns the ID corresponding to ItemCollection
func (i IRIs) GetID() ID {
	return EmptyID
}

// GetLink returns the empty IRI
func (i IRIs) GetLink() IRI {
	return EmptyIRI
}

// GetType returns the ItemCollection's type
func (i IRIs) GetType() ActivityVocabularyType {
	return CollectionOfItems
}

// IsLink returns false for an ItemCollection object
func (i IRIs) IsLink() bool {
	return false
}

// IsObject returns true for a ItemCollection object
func (i IRIs) IsObject() bool {
	return false
}

// IsCollection returns true for IRI slices
func (i IRIs) IsCollection() bool {
	return true
}

// Append facilitates adding elements to IRI slices
// and ensures IRIs implements the Collection interface
func (i *IRIs) Append(r IRI) error {
	*i = append(*i, r)
	return nil
}

func (i *IRIs) Collection() ItemCollection {
	res := make(ItemCollection, len(*i))
	for k, iri := range *i {
		res[k] = iri
	}
	return res
}

func (i *IRIs) Count() uint {
	return uint(len(*i))
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
	if filepath.Clean(u.Path) != filepath.Clean(uw.Path) {
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

func hostSplit(h string) (string, string) {
	pieces := strings.Split(h, ":")
	if len(pieces) == 0 {
		return "", ""
	}
	if len(pieces) == 1 {
		return pieces[0], ""
	}
	return pieces[0], pieces[1]
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
	uHost, _ := hostSplit(u.Host)
	uwHost, _ := hostSplit(uw.Host)
	if uHost != uwHost {
		return false
	}
	p := u.Path
	if p != "" {
		p = filepath.Clean(p)
	}
	pw := uw.Path
	if pw != "" {
		pw = filepath.Clean(pw)
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
