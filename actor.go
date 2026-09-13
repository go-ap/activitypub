package activitypub

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/valyala/fastjson"
)

// CanReceiveActivities Types
const (
	ApplicationType  ActivityVocabularyType = "Application"
	GroupType        ActivityVocabularyType = "Group"
	OrganizationType ActivityVocabularyType = "Organization"
	PersonType       ActivityVocabularyType = "Person"
	ServiceType      ActivityVocabularyType = "Service"
)

// ActorTypes represent the valid Actor types.
var ActorTypes = ActivityVocabularyTypes{
	ApplicationType,
	GroupType,
	OrganizationType,
	PersonType,
	ServiceType,
}

// CanReceiveActivities is generally one of the ActivityStreams Actor Types, but they don't have to be.
// For example, a Profile object might be used as an actor, or a type from an ActivityStreams extension.
// Actors are retrieved like any other Object in ActivityPub.
// Like other ActivityStreams objects, actors have an id, which is a URI.
type CanReceiveActivities Item

type Actors interface {
	Actor
}

// Actor is generally one of the ActivityStreams actor Types, but they don't have to be.
// For example, a Profile object might be used as an actor, or a type from an ActivityStreams extension.
// Actors are retrieved like any other Object in ActivityPub.
// Like other ActivityStreams objects, actors have an id, which is a URI.
//
// Actor objects are specializations of the base Object type that represent entities capable of carrying
// out an Activity. The Activity Vocabulary provides the normative definition of five specific types of
// Actors: Application | Group | Organization | Person | Service.
//
// This specification intentionally defines Actors in only the most generalized way, stopping short of
// defining semantically specific properties for each. All Actor objects are specializations of Object and
// inherit all of the core properties common to all Objects. External vocabularies can be used to express
// additional detail not covered by the Activity Vocabulary. VCard SHOULD be used to provide additional
// metadata for Person, Group, and Organization instances.
//
// https://www.w3.org/TR/activitystreams-vocabulary/#actor-types
type Actor struct {
	// ID provides the globally unique identifier for anActivity Pub Object or Link.
	ID ID `jsonld:"id,omitempty"`
	// Type identifies the Activity Pub Object or Link type. Multiple values may be specified.
	Type Typer `jsonld:"type,omitempty"`
	// Name a simple, human-readable, plain-text name for the object.
	// HTML markup MUST NOT be included. The name MAY be expressed using multiple language-tagged values.
	Name NaturalLanguageValues `jsonld:"name,omitempty,collapsible"`
	// Attachment identifies a resource attached or related to an object that potentially requires special handling.
	// The intent is to provide a model that is at least semantically similar to attachments in email.
	Attachment Item `jsonld:"attachment,omitempty"`
	// AttributedTo identifies one or more entities to which this object is attributed. The attributed entities might not be Actors.
	// For instance, an object might be attributed to the completion of another activity.
	AttributedTo Item `jsonld:"attributedTo,omitempty"`
	// Audience identifies one or more entities that represent the total population of entities
	// for which the object can considered to be relevant.
	Audience ItemCollection `jsonld:"audience,omitempty"`
	// Content or textual representation of the Activity Pub Object encoded as a JSON string.
	// By default, the value of content is HTML.
	// The mediaType property can be used in the object to indicate a different content type.
	// (The content MAY be expressed using multiple language-tagged values.)
	Content NaturalLanguageValues `jsonld:"content,omitempty,collapsible"`
	// Context identifies the context within which the object exists or an activity was performed.
	// The notion of "context" used is intentionally vague.
	// The intended function is to serve as a means of grouping objects and activities that share a
	// common originating context or purpose. An example could be all activities relating to a common project or event.
	Context Item `jsonld:"context,omitempty"`
	// MediaType when used on an Object, identifies the MIME media type of the value of the content property.
	// If not specified, the content property is assumed to contain text/html content.
	MediaType MimeType `jsonld:"mediaType,omitempty"`
	// EndTime the date and time describing the actual or expected ending time of the object.
	// When used with an Activity object, for instance, the endTime property specifies the moment
	// the activity concluded or is expected to conclude.
	EndTime time.Time `jsonld:"endTime,omitempty"`
	// Generator identifies the entity (e.g. an application) that generated the object.
	Generator Item `jsonld:"generator,omitempty"`
	// Icon indicates an entity that describes an icon for this object.
	// The image should have an aspect ratio of one (horizontal) to one (vertical)
	// and should be suitable for presentation at a small size.
	Icon Item `jsonld:"icon,omitempty"`
	// Image indicates an entity that describes an image for this object.
	// Unlike the icon property, there are no aspect ratio or display size limitations assumed.
	Image Item `jsonld:"image,omitempty"`
	// InReplyTo indicates one or more entities for which this object is considered a response.
	InReplyTo Item `jsonld:"inReplyTo,omitempty"`
	// Location indicates one or more physical or logical locations associated with the object.
	Location Item `jsonld:"location,omitempty"`
	// Preview identifies an entity that provides a preview of this object.
	Preview Item `jsonld:"preview,omitempty"`
	// Published the date and time at which the object was published
	Published time.Time `jsonld:"published,omitempty"`
	// Replies identifies a Collection containing objects considered to be responses to this object.
	Replies Item `jsonld:"replies,omitempty"`
	// StartTime the date and time describing the actual or expected starting time of the object.
	// When used with an Activity object, for instance, the startTime property specifies
	// the moment the activity began or is scheduled to begin.
	StartTime time.Time `jsonld:"startTime,omitempty"`
	// Summary a natural language summarization of the object encoded as HTML.
	// *Multiple language tagged summaries may be provided.)
	Summary NaturalLanguageValues `jsonld:"summary,omitempty,collapsible"`
	// Tag one or more "tags" that have been associated with an objects. A tag can be any kind of Activity Pub Object.
	// The key difference between attachment and tag is that the former implies association by inclusion,
	// while the latter implies associated by reference.
	Tag ItemCollection `jsonld:"tag,omitempty"`
	// Updated the date and time at which the object was updated
	Updated time.Time `jsonld:"updated,omitempty"`
	// URL identifies one or more links to representations of the object
	URL Item `jsonld:"url,omitempty"`
	// To identifies an entity considered to be part of the public primary audience of an Activity Pub Object
	To ItemCollection `jsonld:"to,omitempty"`
	// Bto identifies anActivity Pub Object that is part of the private primary audience of this Activity Pub Object.
	Bto ItemCollection `jsonld:"bto,omitempty"`
	// CC identifies anActivity Pub Object that is part of the public secondary audience of this Activity Pub Object.
	CC ItemCollection `jsonld:"cc,omitempty"`
	// BCC identifies one or more Objects that are part of the private secondary audience of this Activity Pub Object.
	BCC ItemCollection `jsonld:"bcc,omitempty"`
	// Duration when the object describes a time-bound resource, such as an audio or video, a meeting, etc,
	// the duration property indicates the object's approximate duration.
	// The value must be expressed as an xsd:duration as defined by [ xmlschema11-2],
	// section 3.3.6 (e.g. a period of 5 seconds is represented as "PT5S").
	Duration time.Duration `jsonld:"duration,omitempty"`
	// This is a list of all Like activities with this object as the object property, added as a side effect.
	// The likes collection MUST be either an OrderedCollection or a Collection and MAY be filtered on privileges
	// of an authenticated user or as appropriate when no authentication is given.
	Likes Item `jsonld:"likes,omitempty"`
	// This is a list of all Announce activities with this object as the object property, added as a side effect.
	// The shares collection MUST be either an OrderedCollection or a Collection and MAY be filtered on privileges
	// of an authenticated user or as appropriate when no authentication is given.
	Shares Item `jsonld:"shares,omitempty"`
	// Source property is intended to convey some sort of source from which the content markup was derived,
	// as a form of provenance, or to support future editing by clients.
	// In general, clients do the conversion from source to content, not the other way around.
	Source Source `jsonld:"source,omitempty"`
	// A reference to an [ActivityStreams] OrderedCollection comprised of all the messages received by the actor;
	// see 5.2 Inbox.
	Inbox Item `jsonld:"inbox,omitempty"`
	// An [ActivityStreams] OrderedCollection comprised of all the messages produced by the actor;
	// see 5.1 outbox.
	Outbox Item `jsonld:"outbox,omitempty"`
	// A link to an [ActivityStreams] collection of the actors that this actor is following;
	// see 5.4 Following Collection
	Following Item `jsonld:"following,omitempty"`
	// A link to an [ActivityStreams] collection of the actors that follow this actor;
	// see 5.3 Followers Collection.
	Followers Item `jsonld:"followers,omitempty"`
	// A link to an [ActivityStreams] collection of objects this actor has liked;
	// see 5.5 Liked Collection.
	Liked Item `jsonld:"liked,omitempty"`
	// A short username which may be used to refer to the actor, with no uniqueness guarantees.
	PreferredUsername NaturalLanguageValues `jsonld:"preferredUsername,omitempty,collapsible"`
	// A json object which maps additional (typically server/domain-wide) endpoints which may be useful either
	// for this actor or someone referencing this actor.
	// This mapping may be nested inside the actor document as the value or may be a link
	// to a JSON-LD document with these properties.
	Endpoints *Endpoints `jsonld:"endpoints,omitempty"`
	// A list of supplementary Collections which may be of interest.
	Streams   ItemCollection `jsonld:"streams,omitempty"`
	PublicKey PublicKey      `jsonld:"publicKey,omitempty"`
}

// GetID returns the ID corresponding to the current Actor
func (a Actor) GetID() ID {
	return a.ID
}

// GetLink returns the IRI corresponding to the current Actor
func (a Actor) GetLink() IRI {
	return IRI(a.ID)
}

// GetType returns the type of the current Actor
func (a Actor) GetType() Typer {
	return a.Type
}

// Match returns whether the receiver matches the ActivityVocabularyType arguments.
func (a Actor) Match(tt ...ActivityVocabularyType) bool {
	return ActivityVocabularyTypes(tt).Match(a.Type)
}

// PublicKey holds the ActivityPub compatible public key data
// The document reference can be found at:
// https://w3c-ccg.github.io/security-vocab/#publicKey
type PublicKey struct {
	ID           ID     `jsonld:"id,omitempty"`
	Owner        IRI    `jsonld:"owner,omitempty"`
	PublicKeyPem string `jsonld:"publicKeyPem,omitempty"`
}

func (p *PublicKey) UnmarshalJSON(data []byte) error {
	par := fastjson.Parser{}
	val, err := par.ParseBytes(data)
	if err != nil {
		return err
	}

	return JSONLoadPublicKey(val, p)
}

func (p PublicKey) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	notEmpty := false
	JSONWrite(&b, '{')
	if v, err := p.ID.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = JSONWriteProp(&b, "id", v, false)
	}
	if len(p.Owner) > 0 {
		notEmpty = JSONWriteIRIProp(&b, "owner", p.Owner, notEmpty) || notEmpty
	}
	if len(p.PublicKeyPem) > 0 {
		if pem, err := json.Marshal(p.PublicKeyPem); err == nil {
			notEmpty = JSONWriteProp(&b, "publicKeyPem", pem, notEmpty) || notEmpty
		}
	}

	if !notEmpty {
		return nil, nil
	}
	JSONWrite(&b, '}')
	return b.Bytes(), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (a *Actor) UnmarshalBinary(data []byte) error {
	return a.GobDecode(data)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (a Actor) MarshalBinary() ([]byte, error) {
	return a.GobEncode()
}

func (a Actor) GobEncode() ([]byte, error) {
	mm := make(map[string][]byte)
	hasData, err := mapActorProperties(mm, &a)
	if err != nil {
		return nil, err
	}
	if !hasData {
		return []byte{}, nil
	}
	bb := bytes.Buffer{}
	g := gob.NewEncoder(&bb)
	if err := g.Encode(mm); err != nil {
		return nil, err
	}
	return bb.Bytes(), nil
}

func (a *Actor) GobDecode(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	mm, err := gobDecodeObjectAsMap(data)
	if err != nil {
		return err
	}
	return unmapActorProperties(mm, a)
}

type (
	// Application describes a software application.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-application
	Application = Actor

	// Group represents a formal or informal collective of Actors.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-group
	Group = Actor

	// Organization represents an organization.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-organization
	Organization = Actor

	// Person represents an individual person.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-person
	Person = Actor

	// Service represents a service of any kind.
	//
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-service
	Service = Actor
)

func (a *Actor) Recipients() ItemCollection {
	aud := a.Audience
	return ItemCollectionDeduplication(&a.To, &a.CC, &a.Bto, &a.BCC, &aud)
}

func cleanActorProperties(aa *Actor) {
	_ = OnObject(aa, func(ob *Object) error {
		cleanObjectProperties(ob)
		return nil
	})
}

func (a *Actor) Clean() Item {
	aa := *a
	cleanActorProperties(&aa)
	return &aa
}

func (a *Actor) UnmarshalJSON(data []byte) error {
	p := fastjson.Parser{}
	val, err := p.ParseBytes(data)
	if err != nil {
		return err
	}
	return JSONLoadActor(val, a)
}

func (a Actor) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	notEmpty := false
	JSONWrite(&b, '{')

	_ = OnObject(a, func(o *Object) error {
		notEmpty = JSONWriteObjectValue(&b, *o)
		return nil
	})
	if a.Inbox != nil {
		notEmpty = JSONWriteItemProp(&b, "inbox", a.Inbox, notEmpty) || notEmpty
	}
	if a.Outbox != nil {
		notEmpty = JSONWriteItemProp(&b, "outbox", a.Outbox, notEmpty) || notEmpty
	}
	if a.Following != nil {
		notEmpty = JSONWriteItemProp(&b, "following", a.Following, notEmpty) || notEmpty
	}
	if a.Followers != nil {
		notEmpty = JSONWriteItemProp(&b, "followers", a.Followers, notEmpty) || notEmpty
	}
	if a.Liked != nil {
		notEmpty = JSONWriteItemProp(&b, "liked", a.Liked, notEmpty) || notEmpty
	}
	if a.PreferredUsername != nil {
		notEmpty = JSONWriteNaturalLanguageProp(&b, "preferredUsername", a.PreferredUsername, notEmpty) || notEmpty
	}
	if a.Endpoints != nil {
		if v, err := a.Endpoints.MarshalJSON(); err == nil && len(v) > 0 {
			notEmpty = JSONWriteProp(&b, "endpoints", v, notEmpty) || notEmpty
		}
	}
	if len(a.Streams) > 0 {
		notEmpty = JSONWriteItemCollectionProp(&b, "streams", a.Streams, false, notEmpty)
	}
	if len(a.PublicKey.PublicKeyPem)+len(a.PublicKey.ID) > 0 {
		if v, err := a.PublicKey.MarshalJSON(); err == nil && len(v) > 0 {
			notEmpty = JSONWriteProp(&b, "publicKey", v, notEmpty) || notEmpty
		}
	}

	if !notEmpty {
		return nil, nil
	}
	JSONWrite(&b, '}')
	return b.Bytes(), nil
}

func fmtActorProps(w io.Writer, n *int) func(*Actor) error {
	return func(a *Actor) error {
		_ = OnObject(a, fmtObjectProps(w, n))
		comma := func() {
			if *n > 0 {
				_, _ = io.WriteString(w, ", ")
			}
		}

		if len(a.PreferredUsername) > 0 {
			comma()
			*n, _ = fmt.Fprintf(w, "preferredUsername: %s", a.PreferredUsername)
		}
		return nil
	}
}

func (a Actor) Format(s fmt.State, verb rune) {
	typ := a.Type
	switch verb {
	case 's':
		iri := a.ID
		if iri != "" {
			s.Write([]byte(iri))
		} else {
			_, _ = fmt.Fprintf(s, "%T[%v]", a, typ)
		}
	case 'v':
		n := 0
		if typ != nil {
			_, _ = fmt.Fprintf(s, "%T[%v] { ", a, typ)
			_ = fmtActorProps(s, &n)(&a)
			_, _ = io.WriteString(s, " }")
		} else {
			_, _ = fmt.Fprintf(s, "%T { ", a)
			_ = fmtActorProps(s, &n)(&a)
			_, _ = io.WriteString(s, " }")
		}
	}
}

// Endpoints maps additional (typically server/domain-wide) endpoints which may be
// useful either for this actor or someone referencing this actor.
// This mapping may be nested inside the actor document as the value, or may be a link to
// a JSON-LD document with these properties (this is not yet possible in the library).
type Endpoints struct {
	// UploadMedia Upload endpoint URI for this user for binary data.
	UploadMedia Item `jsonld:"uploadMedia,omitempty"`
	// OauthAuthorizationEndpoint Endpoint URI so this actor's clients may access remote ActivityStreams objects which require authentication
	// to access. To use this endpoint, the client posts an x-www-form-urlencoded id parameter with the value being
	// the id of the requested ActivityStreams object.
	OauthAuthorizationEndpoint Item `jsonld:"oauthAuthorizationEndpoint,omitempty"`
	// OauthTokenEndpoint If OAuth 2.0 bearer tokens [RFC6749] [RFC6750] are being used for authenticating client to server interactions,
	// this endpoint specifies a URI at which a browser-authenticated user may obtain a new authorization grant.
	OauthTokenEndpoint Item `jsonld:"oauthTokenEndpoint,omitempty"`
	// ProvideClientKey  If OAuth 2.0 bearer tokens [RFC6749] [RFC6750] are being used for authenticating client to server interactions,
	// this endpoint specifies a URI at which a client may acquire an access token.
	ProvideClientKey Item `jsonld:"provideClientKey,omitempty"`
	// SignClientKey If Linked Data Signatures and HTTP Signatures are being used for authentication and authorization,
	// this endpoint specifies a URI at which browser-authenticated users may authorize a client's public
	// key for client to server interactions.
	SignClientKey Item `jsonld:"signClientKey,omitempty"`
	// SharedInbox An optional endpoint used for wide delivery of publicly addressed activities and activities sent to followers.
	// SharedInbox endpoints SHOULD also be publicly readable OrderedCollection objects containing objects addressed to the
	// Public special collection. Reading from the sharedInbox endpoint MUST NOT present objects which are not addressed to the Public endpoint.
	SharedInbox Item `jsonld:"sharedInbox,omitempty"`
	// ProxyURL
	// Endpoint URI so this actor's clients may access remote ActivityStreams objects which require authentication to
	// access. To use this endpoint, the client posts an x-www-form-urlencoded id parameter with the value being the id
	// of the requested ActivityStreams object.
	ProxyURL IRI `jsonld:"proxyUrl,omitempty"`
}

// UnmarshalJSON decodes an incoming JSON document into the receiver object.
func (e *Endpoints) UnmarshalJSON(data []byte) error {
	p := fastjson.Parser{}
	val, err := p.ParseBytes(data)
	if err != nil {
		return err
	}
	e.OauthAuthorizationEndpoint = JSONGetItem(val, "oauthAuthorizationEndpoint")
	e.OauthTokenEndpoint = JSONGetItem(val, "oauthTokenEndpoint")
	e.UploadMedia = JSONGetItem(val, "uploadMedia")
	e.ProvideClientKey = JSONGetItem(val, "provideClientKey")
	e.SignClientKey = JSONGetItem(val, "signClientKey")
	e.SharedInbox = JSONGetItem(val, "sharedInbox")
	e.ProxyURL = JSONGetIRI(val, "proxyUrl")
	return nil
}

// MarshalJSON encodes the receiver object to a JSON document.
func (e Endpoints) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	notEmpty := false

	JSONWrite(&b, '{')
	if e.OauthAuthorizationEndpoint != nil {
		notEmpty = JSONWriteItemProp(&b, "oauthAuthorizationEndpoint", e.OauthAuthorizationEndpoint, notEmpty)
	}
	if e.OauthTokenEndpoint != nil {
		notEmpty = JSONWriteItemProp(&b, "oauthTokenEndpoint", e.OauthTokenEndpoint, notEmpty) || notEmpty
	}
	if e.ProvideClientKey != nil {
		notEmpty = JSONWriteItemProp(&b, "provideClientKey", e.ProvideClientKey, notEmpty) || notEmpty
	}
	if e.SignClientKey != nil {
		notEmpty = JSONWriteItemProp(&b, "signClientKey", e.SignClientKey, notEmpty) || notEmpty
	}
	if e.SharedInbox != nil {
		notEmpty = JSONWriteItemProp(&b, "sharedInbox", e.SharedInbox, notEmpty) || notEmpty
	}
	if e.UploadMedia != nil {
		notEmpty = JSONWriteItemProp(&b, "uploadMedia", e.UploadMedia, notEmpty) || notEmpty
	}
	if e.ProxyURL != NilID {
		notEmpty = JSONWriteItemProp(&b, "proxyUrl", e.ProxyURL, notEmpty) || notEmpty
	}
	if !notEmpty {
		return nil, nil
	}
	JSONWrite(&b, '}')
	return b.Bytes(), nil
}

func tombstoneAsActor(t *Tombstone) (*Actor, error) {
	a := new(Actor)
	err := OnObject(a, func(aob *Object) error {
		return OnObject(t, func(tob *Object) error {
			_, err := CopyObjectProperties(aob, tob)
			return err
		})
	})
	return a, err
}

// ToActor
func ToActor(it LinkOrIRI) (*Actor, error) {
	switch i := it.(type) {
	case *Actor:
		return i, nil
	case Actor:
		return &i, nil
	case *Tombstone:
		return tombstoneAsActor(i)
	case Tombstone:
		return tombstoneAsActor(&i)
	default:
		return reflectItemToType[Actor](it)
	}
}

// Equals verifies if our receiver Object is equals with the "with" Item
func (a Actor) Equals(with Item) bool {
	withActor, err := ToActor(with)
	if err != nil {
		return false
	}
	return a.equal(*withActor)
}

// equal verifies if our receiver Object is equals with the "with" Object
func (a Actor) equal(with Actor) bool {
	result := true
	_ = OnObject(a, func(oa *Object) error {
		result = oa.Equals(with)
		return nil
	})
	if !result {
		return false
	}
	if !ItemsEqual(a.Inbox, with.Inbox) {
		return false
	}
	if !ItemsEqual(a.Outbox, with.Outbox) {
		return false
	}
	if !ItemsEqual(a.Liked, with.Liked) {
		return false
	}
	if !a.PreferredUsername.Equal(with.PreferredUsername) {
		return false
	}
	return true
}

func (e Endpoints) GobEncode() ([]byte, error) {
	return nil, nil
}

func (e *Endpoints) GobDecode(data []byte) error {
	return nil
}

func (p PublicKey) GobEncode() ([]byte, error) {
	var (
		mm      = make(map[string][]byte)
		err     error
		hasData bool
	)
	if len(p.ID) > 0 {
		if mm["id"], err = p.ID.GobEncode(); err != nil {
			return nil, err
		}
		hasData = true
	}
	if len(p.PublicKeyPem) > 0 {
		mm["publicKeyPem"] = []byte(p.PublicKeyPem)
		hasData = true
	}
	if len(p.Owner) > 0 {
		if mm["owner"], err = gobEncodeItem(p.Owner); err != nil {
			return nil, err
		}
		hasData = true
	}
	if !hasData {
		return []byte{}, nil
	}
	bb := bytes.Buffer{}
	g := gob.NewEncoder(&bb)
	if err := g.Encode(mm); err != nil {
		return nil, err
	}
	return bb.Bytes(), nil
}

func (p *PublicKey) GobDecode(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	mm, err := gobDecodeObjectAsMap(data)
	if err != nil {
		return err
	}
	if raw, ok := mm["id"]; ok {
		if err = p.ID.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["owner"]; ok {
		if err = p.Owner.GobDecode(raw); err != nil {
			return err
		}
	}
	if raw, ok := mm["publicKeyPem"]; ok {
		p.PublicKeyPem = string(raw)
	}
	return nil
}

// WithActorFn represents a function type that can be used as a parameter for OnActor helper function
type WithActorFn func(*Actor) error

// OnActor calls function fn on it Item if it can be asserted to type *Actor
//
// This function should be called if trying to access the Actor specific
// properties like "preferredName", "publicKey", etc. For the other properties
// OnObject should be used instead.
func OnActor(it LinkOrIRI, fn func(*Actor) error) error {
	if IsNil(it) {
		return nil
	}

	actFn := func(it LinkOrIRI) error {
		act, err := ToActor(it)
		if err != nil {
			return err
		}
		return fn(act)
	}

	if !IsItemCollection(it) {
		return actFn(it)
	}
	return OnItemCollection(it, func(col *ItemCollection) error {
		for _, ob := range *col {
			if err := actFn(ob); err != nil {
				return err
			}
		}
		return nil
	})
}

func notEmptyActor(a *Actor) bool {
	var notEmpty bool
	_ = OnObject(a, func(o *Object) error {
		notEmpty = notEmptyObject(o)
		return nil
	})
	return notEmpty ||
		a.Inbox != nil ||
		a.Outbox != nil ||
		a.Following != nil ||
		a.Followers != nil ||
		a.Liked != nil ||
		a.PreferredUsername != nil ||
		a.Endpoints != nil ||
		a.Streams != nil ||
		len(a.PublicKey.ID)+len(a.PublicKey.Owner)+len(a.PublicKey.PublicKeyPem) > 0
}

// CopyActorProperties
func CopyActorProperties(to, from *Actor) (*Actor, error) {
	oldOb, _ := ToObject(to)
	newOb, _ := ToObject(from)
	_, err := CopyObjectProperties(oldOb, newOb)
	if err != nil {
		return to, err
	}
	to.Inbox = replaceIfItem(to.Inbox, from.Inbox)
	to.Outbox = replaceIfItem(to.Outbox, from.Outbox)
	to.Following = replaceIfItem(to.Following, from.Following)
	to.Followers = replaceIfItem(to.Followers, from.Followers)
	to.Liked = replaceIfItem(to.Liked, from.Liked)
	to.PreferredUsername = replaceIfNaturalLanguageValues(to.PreferredUsername, from.PreferredUsername)
	to.PublicKey = replaceIfPublicKey(to.PublicKey, from.PublicKey)
	return to, nil
}
