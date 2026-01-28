package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"net/url"
	"path/filepath"
	"slices"
	"strings"

	"github.com/go-ap/errors"
	"github.com/valyala/fastjson"
)

const (
	// ActivityBaseURI the URI for the ActivityStreams namespace
	ActivityBaseURI = IRI("https://www.w3.org/ns/activitystreams")
	// SecurityContextURI the URI for the security namespace (for an Actor's PublicKey)
	SecurityContextURI = IRI("https://w3id.org/security/v1")
	// PublicNS is the reference to the Public entity in the ActivityStreams namespace.
	//
	// Public Addressing
	//
	// https://www.w3.org/TR/activitypub/#public-addressing
	//
	// In addition to [ActivityStreams] collections and objects, Activities may additionally be addressed to the
	// special "public" collection, with the identifier https://www.w3.org/ns/activitystreams#Public. For example:
	//
	// {
	//   "@context": "https://www.w3.org/ns/activitystreams",
	//   "id": "https://www.w3.org/ns/activitystreams#Public",
	//   "type": "Collection"
	// }
	// Activities addressed to this special URI shall be accessible to all users, without authentication.
	// Implementations MUST NOT deliver to the "public" special collection; it is not capable of receiving
	// actual activities. However, actors MAY have a sharedInbox endpoint which is available for efficient
	// shared delivery of public posts (as well as posts to followers-only); see 7.1.3 Shared Inbox Delivery.
	//
	// NOTE
	// Compacting an ActivityStreams object using the ActivityStreams JSON-LD context might result in
	// https://www.w3.org/ns/activitystreams#Public being represented as simply Public or as:Public which are valid
	// representations of the Public collection. Implementations which treat ActivityStreams objects as simply JSON
	// rather than converting an incoming activity over to a local context using JSON-LD tooling should be aware
	// of this and should be prepared to accept all three representations.
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
		_, _ = io.WriteString(s, i.String())
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
	if i == "" {
		return nil, errors.Newf("empty IRI")
	}
	return url.Parse(string(i))
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
	JSONWrite(&b, '"')
	JSONWriteS(&b, i.String())
	JSONWrite(&b, '"')
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
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(i)
	if err == nil {
		return nil
	}
	bb := make([][]byte, 0)
	err = gob.NewDecoder(bytes.NewReader(data)).Decode(&bb)
	if err != nil {
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
func (i IRI) GetType() TypeMatcher {
	return IRIType
}

// Matches returns whether the receiver matches the ActivityVocabularyType arguments.
func (i IRI) Matches(tt ...ActivityVocabularyType) bool {
	return i.GetType().Matches(tt...)
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
	if len(i) == 0 {
		return []byte{'[', ']'}, nil
	}
	b := make([]byte, 0)
	writeCommaIfNotEmpty := func(notEmpty bool) {
		if notEmpty {
			JSONWriteS(&b, ",")
		}
	}
	JSONWrite(&b, '[')
	for k, iri := range i {
		writeCommaIfNotEmpty(k > 0)
		JSONWrite(&b, '"')
		JSONWriteS(&b, iri.String())
		JSONWrite(&b, '"')
	}
	JSONWrite(&b, ']')
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
func (i IRIs) GetType() TypeMatcher {
	return CollectionOfIRIs
}

// GetTypes returns the ItemCollection's types
func (i IRIs) GetTypes() ActivityVocabularyTypes {
	return ActivityVocabularyTypes{CollectionOfIRIs}
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

// Append facilitates adding elements to the IRIs slices
func (i *IRIs) Append(items ...Item) error {
	for _, it := range items {
		if IsNil(it) {
			continue
		}
		if (*i).Contains(it.GetLink()) {
			continue
		}
		*i = append(*i, it.GetLink())
	}
	return nil
}

// Remove removes the items from the IRIs slice
func (i *IRIs) Remove(it ...Item) {
	items := ItemCollection(it)
	*i = slices.DeleteFunc(*i, func(iri IRI) bool {
		return items.Contains(iri)
	})
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
func (i IRIs) Contains(r Item) bool {
	if len(i) == 0 {
		return false
	}
	for _, iri := range i {
		if r.GetLink().Equals(iri, false) {
			return true
		}
	}
	return false
}

func validURL(u *url.URL, checkScheme bool) bool {
	if u == nil {
		return false
	}
	if len(u.Host) == 0 {
		return false
	}
	if checkScheme {
		return len(u.Scheme) > 0
	}
	return true
}

func stripFragment(u string) string {
	p := strings.Index(u, "#")
	if p <= 0 {
		p = len(u)
	}
	return u[:p]
}

func stripScheme(u string) string {
	p := strings.Index(u, "://")
	if p < 0 {
		p = 0
	}
	return u[p:]
}

func irisEqual(i1, i2 IRI, checkScheme bool) bool {
	u, e := i1.URL()
	uw, ew := i2.URL()
	if e != nil || ew != nil || !validURL(u, checkScheme) || !validURL(uw, checkScheme) {
		return strings.EqualFold(i1.String(), i2.String())
	}
	if checkScheme {
		if !strings.EqualFold(u.Scheme, uw.Scheme) {
			return false
		}
	}
	if !strings.EqualFold(u.Host, uw.Host) {
		return false
	}
	if u.Path == "" {
		u.Path = "/"
	}
	if uw.Path == "" {
		uw.Path = "/"
	}
	if !strings.EqualFold(filepath.Clean(u.Path), filepath.Clean(uw.Path)) {
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

// Equal verifies if our receiver IRI is equals with the "with" IRI with strict protocol check
func (i IRI) Equal(with IRI) bool {
	return i.Equals(with, true)
}

// Equals verifies if our receiver IRI is equals with the "with" IRI
func (i IRI) Equals(with IRI, checkScheme bool) bool {
	is := stripFragment(string(i))
	ws := stripFragment(string(with))
	if !checkScheme {
		is = stripScheme(is)
		ws = stripScheme(ws)
	}
	if strings.EqualFold(is, ws) {
		return true
	}
	return irisEqual(i, with, checkScheme)
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
	} else {
		p = "/"
	}
	pw := uw.Path
	if pw != "" {
		pw = filepath.Clean(pw)
	} else {
		pw = "/"
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

// OnIRIs calls function fn on it Item if it can be asserted to type IRIs
//
// It should be used when Item represents an IRI slice.
func OnIRIs(it Item, fn WithIRIsFn) error {
	if it == nil {
		return nil
	}
	col, err := ToIRIs(it)
	if err != nil {
		return err
	}
	return fn(col)
}
