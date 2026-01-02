package activitypub

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
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
type Actor struct {
	// ID provides the globally unique identifier for anActivity Pub Object or Link.
	ID ID `jsonld:"id,omitempty"`
	// Type identifies the Activity Pub Object or Link type. Multiple values may be specified.
	Type ActivityVocabularyType `jsonld:"type,omitempty"`
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
func (a Actor) GetType() ActivityVocabularyType {
	return a.Type
}

// IsLink validates if currentActivity Pub Actor is a Link
func (a Actor) IsLink() bool {
	return false
}

// IsObject validates if currentActivity Pub Actor is an Object
func (a Actor) IsObject() bool {
	return true
}

// IsCollection returns false for Actor Objects
func (a Actor) IsCollection() bool {
	return false
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
	b := make([]byte, 0)
	notEmpty := true
	JSONWrite(&b, '{')
	if v, err := p.ID.MarshalJSON(); err == nil && len(v) > 0 {
		notEmpty = !JSONWriteProp(&b, "id", v)
	}
	if len(p.Owner) > 0 {
		notEmpty = JSONWriteIRIProp(&b, "owner", p.Owner) || notEmpty
	}
	if len(p.PublicKeyPem) > 0 {
		if pem, err := json.Marshal(p.PublicKeyPem); err == nil {
			notEmpty = JSONWriteProp(&b, "publicKeyPem", pem) || notEmpty
		}
	}

	if notEmpty {
		JSONWrite(&b, '}')
		return b, nil
	}
	return nil, nil
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
	Application = Actor

	// Group represents a formal or informal collective of Actors.
	Group = Actor

	// Organization represents an organization.
	Organization = Actor

	// Person represents an individual person.
	Person = Actor

	// Service represents a service of any kind.
	Service = Actor
)

// ActorNew initializes an CanReceiveActivities type actor
func ActorNew(id ID, typ ActivityVocabularyType) *Actor {
	if !ActorTypes.Contains(typ) {
		typ = ActorType
	}

	a := Actor{ID: id, Type: typ}
	a.Name = NaturalLanguageValuesNew()
	a.Content = NaturalLanguageValuesNew()
	a.Summary = NaturalLanguageValuesNew()
	a.PreferredUsername = NaturalLanguageValuesNew()

	return &a
}

// ApplicationNew initializes an Application type actor
func ApplicationNew(id ID) *Application {
	a := ActorNew(id, ApplicationType)
	o := Application(*a)
	return &o
}

// GroupNew initializes a Group type actor
func GroupNew(id ID) *Group {
	a := ActorNew(id, GroupType)
	o := Group(*a)
	return &o
}

// OrganizationNew initializes an Organization type actor
func OrganizationNew(id ID) *Organization {
	a := ActorNew(id, OrganizationType)
	o := Organization(*a)
	return &o
}

// PersonNew initializes a Person type actor
func PersonNew(id ID) *Person {
	a := ActorNew(id, PersonType)
	o := Person(*a)
	return &o
}

// ServiceNew initializes a Service type actor
func ServiceNew(id ID) *Service {
	a := ActorNew(id, ServiceType)
	o := Service(*a)
	return &o
}

func (a *Actor) Recipients() ItemCollection {
	aud := a.Audience
	return ItemCollectionDeduplication(&a.To, &a.CC, &a.Bto, &a.BCC, &aud)
}

func (a *Actor) Clean() {
	_ = OnObject(a, func(o *Object) error {
		o.Clean()
		return nil
	})
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
	b := make([]byte, 0)
	notEmpty := false
	JSONWrite(&b, '{')

	_ = OnObject(a, func(o *Object) error {
		notEmpty = JSONWriteObjectValue(&b, *o)
		return nil
	})
	if a.Inbox != nil {
		notEmpty = JSONWriteItemProp(&b, "inbox", a.Inbox) || notEmpty
	}
	if a.Outbox != nil {
		notEmpty = JSONWriteItemProp(&b, "outbox", a.Outbox) || notEmpty
	}
	if a.Following != nil {
		notEmpty = JSONWriteItemProp(&b, "following", a.Following) || notEmpty
	}
	if a.Followers != nil {
		notEmpty = JSONWriteItemProp(&b, "followers", a.Followers) || notEmpty
	}
	if a.Liked != nil {
		notEmpty = JSONWriteItemProp(&b, "liked", a.Liked) || notEmpty
	}
	if a.PreferredUsername != nil {
		notEmpty = JSONWriteNaturalLanguageProp(&b, "preferredUsername", a.PreferredUsername) || notEmpty
	}
	if a.Endpoints != nil {
		if v, err := a.Endpoints.MarshalJSON(); err == nil && len(v) > 0 {
			notEmpty = JSONWriteProp(&b, "endpoints", v) || notEmpty
		}
	}
	if len(a.Streams) > 0 {
		notEmpty = JSONWriteItemCollectionProp(&b, "streams", a.Streams, false)
	}
	if len(a.PublicKey.PublicKeyPem)+len(a.PublicKey.ID) > 0 {
		if v, err := a.PublicKey.MarshalJSON(); err == nil && len(v) > 0 {
			notEmpty = JSONWriteProp(&b, "publicKey", v) || notEmpty
		}
	}

	if notEmpty {
		JSONWrite(&b, '}')
		return b, nil
	}
	return nil, nil
}

func (a Actor) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		if a.Type != "" && a.ID != "" {
			_, _ = fmt.Fprintf(s, "%T[%s]( %s )", a, a.Type, a.ID)
		} else if a.ID != "" {
			_, _ = fmt.Fprintf(s, "%T( %s )", a, a.ID)
		} else {
			_, _ = fmt.Fprintf(s, "%T[%p]", a, &a)
		}
	case 'v':
		_, _ = fmt.Fprintf(s, "%T[%s] { }", a, a.Type)
	}
}

// Endpoints a json object which maps additional (typically server/domain-wide)
// endpoints which may be useful either for this actor or someone referencing this actor.
// This mapping may be nested inside the actor document as the value or may be a link to
// a JSON-LD document with these properties.
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
	return nil
}

// MarshalJSON encodes the receiver object to a JSON document.
func (e Endpoints) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	notEmpty := false

	JSONWrite(&b, '{')
	if e.OauthAuthorizationEndpoint != nil {
		notEmpty = JSONWriteItemProp(&b, "oauthAuthorizationEndpoint", e.OauthAuthorizationEndpoint) || notEmpty
	}
	if e.OauthTokenEndpoint != nil {
		notEmpty = JSONWriteItemProp(&b, "oauthTokenEndpoint", e.OauthTokenEndpoint) || notEmpty
	}
	if e.ProvideClientKey != nil {
		notEmpty = JSONWriteItemProp(&b, "provideClientKey", e.ProvideClientKey) || notEmpty
	}
	if e.SignClientKey != nil {
		notEmpty = JSONWriteItemProp(&b, "signClientKey", e.SignClientKey) || notEmpty
	}
	if e.SharedInbox != nil {
		notEmpty = JSONWriteItemProp(&b, "sharedInbox", e.SharedInbox) || notEmpty
	}
	if e.UploadMedia != nil {
		notEmpty = JSONWriteItemProp(&b, "uploadMedia", e.UploadMedia) || notEmpty
	}
	if notEmpty {
		JSONWrite(&b, '}')
		return b, nil
	}
	return nil, nil
}

// ToActor
func ToActor(it LinkOrIRI) (*Actor, error) {
	switch i := it.(type) {
	case *Actor:
		return i, nil
	case Actor:
		return &i, nil
	default:
		return reflectItemToType[Actor](it)
	}
}

// Equals verifies if our receiver Object is equals with the "with" Item
func (a *Actor) Equals(with Item) bool {
	if IsNil(with) {
		return a == nil
	}
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
	if with.Inbox != nil {
		if !ItemsEqual(a.Inbox, with.Inbox) {
			result = false
		}
	}
	if with.Outbox != nil {
		if !ItemsEqual(a.Outbox, with.Outbox) {
			result = false
		}
	}
	if with.Liked != nil {
		if !ItemsEqual(a.Liked, with.Liked) {
			result = false
		}
	}
	if with.PreferredUsername != nil {
		if !a.PreferredUsername.Equal(with.PreferredUsername) {
			result = false
		}
	}
	return result
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
	if it == nil {
		return nil
	}
	if IsItemCollection(it) {
		return callOnItemCollection(it, OnActor, fn)
	}
	act, err := ToActor(it)
	if err != nil {
		return err
	}
	return fn(act)
}
