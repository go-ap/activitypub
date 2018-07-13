package activitypub

import "time"

// Actor Types
const (
	ApplicationType  ActivityVocabularyType = "Application"
	GroupType        ActivityVocabularyType = "Group"
	OrganizationType ActivityVocabularyType = "Organization"
	PersonType       ActivityVocabularyType = "Person"
	ServiceType      ActivityVocabularyType = "Service"
)

var validActorTypes = [...]ActivityVocabularyType{
	ApplicationType,
	GroupType,
	OrganizationType,
	PersonType,
	ServiceType,
}

// Endpoints a json object which maps additional (typically server/domain-wide)
// endpoints which may be useful either for this actor or someone referencing this actor.
// This mapping may be nested inside the actor document as the value or may be a link to
// a JSON-LD document with these properties.
type Endpoints struct {
	// Upload endpoint URI for this user for binary data.
	UploadMedia ObjectOrLink `jsonld:"uploadMedia,omitempty"`
	// Endpoint URI so this actor's clients may access remote ActivityStreams objects which require authentication
	//  to access. To use this endpoint, the client posts an x-www-form-urlencoded id parameter with the value being
	//  the id of the requested ActivityStreams object.
	OauthAuthorizationEndpoint ObjectOrLink `jsonld:"oauthAuthorizationEndpoint,omitempty"`
	// If OAuth 2.0 bearer tokens [RFC6749] [RFC6750] are being used for authenticating client to server interactions,
	//  this endpoint specifies a URI at which a browser-authenticated user may obtain a new authorization grant.
	OauthTokenEndpoint ObjectOrLink `jsonld:"oauthTokenEndpoint,omitempty"`
	// If OAuth 2.0 bearer tokens [RFC6749] [RFC6750] are being used for authenticating client to server interactions,
	//  this endpoint specifies a URI at which a client may acquire an access token.
	ProvideClientKey ObjectOrLink `jsonld:"provideClientKey,omitempty"`
	// If Linked Data Signatures and HTTP Signatures are being used for authentication and authorization,
	//  this endpoint specifies a URI at which browser-authenticated users may authorize a client's public
	//  key for client to server interactions.
	SignClientKey ObjectOrLink `jsonld:"signClientKey,omitempty"`
	// If Linked Data Signatures and HTTP Signatures are being used for authentication and authorization,
	//  this endpoint specifies a URI at which a client key may be signed by the actor's key for a time window to
	//  act on behalf of the actor in interacting with foreign servers.
	SharedInbox ObjectOrLink `jsonld:"sharedInbox,omitempty"`
}

// Actor is generally one of the ActivityStreams Actor Types, but they don't have to be.
// For example, a Profile object might be used as an actor, or a type from an ActivityStreams extension.
// Actors are retrieved like any other GetID in ActivityPub.
// Like other ActivityStreams objects, actors have an id, which is a URI.
type Actor struct {
	// Provides the globally unique identifier for an Activity Pub GetID or GetLink.
	ID ObjectID `jsonld:"id,omitempty"`
	//  Identifies the Activity Pub GetID or GetLink type. Multiple values may be specified.
	Type ActivityVocabularyType `jsonld:"type,omitempty"`
	// A simple, human-readable, plain-text name for the object.
	// HTML markup MUST NOT be included. The name MAY be expressed using multiple language-tagged values.
	Name NaturalLanguageValue `jsonld:"name,omitempty,collapsible"`
	// Identifies a resource attached or related to an object that potentially requires special handling.
	// The intent is to provide a model that is at least semantically similar to attachments in email.
	Attachment ObjectOrLink `jsonld:"attachment,omitempty"`
	// Identifies one or more entities to which this object is attributed. The attributed entities might not be Actors.
	// For instance, an object might be attributed to the completion of another activity.
	AttributedTo ObjectOrLink `jsonld:"attributedTo,omitempty"`
	// Identifies one or more entities that represent the total population of entities
	//  for which the object can considered to be relevant.
	Audience ObjectOrLink `jsonld:"audience,omitempty"`
	// The content or textual representation of the Activity Pub GetID encoded as a JSON string.
	// By default, the value of content is HTML.
	// The mediaType property can be used in the object to indicate a different content type.
	// (The content MAY be expressed using multiple language-tagged values.)
	Content NaturalLanguageValue `jsonld:"content,omitempty,collapsible"`
	// Identifies the context within which the object exists or an activity was performed.
	// The notion of "context" used is intentionally vague.
	// The intended function is to serve as a means of grouping objects and activities that share a
	//  common originating context or purpose. An example could be all activities relating to a common project or event.
	//Context ObjectOrLink `jsonld:"_"`
	// The date and time describing the actual or expected ending time of the object.
	// When used with an Activity object, for instance, the endTime property specifies the moment
	//  the activity concluded or is expected to conclude.
	EndTime time.Time `jsonld:"endTime,omitempty"`
	// Identifies the entity (e.g. an application) that generated the object.
	Generator ObjectOrLink `jsonld:"generator,omitempty"`
	// Indicates an entity that describes an icon for this object.
	// The image should have an aspect ratio of one (horizontal) to one (vertical)
	//  and should be suitable for presentation at a small size.
	Icon ImageOrLink `jsonld:"icon,omitempty"`
	// Indicates an entity that describes an image for this object.
	// Unlike the icon property, there are no aspect ratio or display size limitations assumed.
	Image ImageOrLink `jsonld:"image,omitempty"`
	// Indicates one or more entities for which this object is considered a response.
	InReplyTo ObjectOrLink `jsonld:"inReplyTo,omitempty"`
	// Indicates one or more physical or logical locations associated with the object.
	Location ObjectOrLink `jsonld:"location,omitempty"`
	// Identifies an entity that provides a preview of this object.
	Preview ObjectOrLink `jsonld:"preview,omitempty"`
	// The date and time at which the object was published
	Published time.Time `jsonld:"published,omitempty"`
	// Identifies a Collection containing objects considered to be responses to this object.
	Replies ObjectOrLink `jsonld:"replies,omitempty"`
	// The date and time describing the actual or expected starting time of the object.
	// When used with an Activity object, for instance, the startTime property specifies
	//  the moment the activity began or is scheduled to begin.
	StartTime time.Time `jsonld:"startTime,omitempty"`
	// A natural language summarization of the object encoded as HTML.
	// *Multiple language tagged summaries may be provided.)
	Summary NaturalLanguageValue `jsonld:"summary,omitempty,collapsible"`
	// One or more "tags" that have been associated with an objects. A tag can be any kind of Activity Pub GetID.
	// The key difference between attachment and tag is that the former implies association by inclusion,
	//  while the latter implies associated by reference.
	Tag ObjectOrLink `jsonld:"tag,omitempty"`
	// The date and time at which the object was updated
	Updated time.Time `jsonld:"updated,omitempty"`
	// Identifies one or more links to representations of the object
	URL LinkOrURI `jsonld:"url,omitempty"`
	// Identifies an entity considered to be part of the public primary audience of an Activity Pub GetID
	To ObjectsArr `jsonld:"to,omitempty"`
	// Identifies an Activity Pub GetID that is part of the private primary audience of this Activity Pub GetID.
	Bto ObjectsArr `jsonld:"bto,omitempty"`
	// Identifies an Activity Pub GetID that is part of the public secondary audience of this Activity Pub GetID.
	CC ObjectsArr `jsonld:"cc,omitempty"`
	// Identifies one or more Objects that are part of the private secondary audience of this Activity Pub GetID.
	BCC ObjectsArr `jsonld:"bcc,omitempty"`
	// When the object describes a time-bound resource, such as an audio or video, a meeting, etc,
	//  the duration property indicates the object's approximate duration.
	// The value must be expressed as an xsd:duration as defined by [ xmlschema11-2],
	//  section 3.3.6 (e.g. a period of 5 seconds is represented as "PT5S").
	Duration time.Duration `jsonld:"duration,omitempty"`
	// A reference to an [ActivityStreams] OrderedCollection comprised of all the messages received by the actor;
	//  see 5.2 Inbox.
	Inbox InboxStream `jsonld:"inbox,omitempty"`
	// An [ActivityStreams] OrderedCollection comprised of all the messages produced by the actor;
	//  see 5.1 Outbox.
	Outbox OutboxStream `jsonld:"outbox,omitempty"`
	// A link to an [ActivityStreams] collection of the actors that this actor is following;
	//  see 5.4 Following Collection
	Following FollowingCollection `jsonld:"following,omitempty"`
	// A link to an [ActivityStreams] collection of the actors that follow this actor;
	//  see 5.3 Followers Collection.
	Followers FollowersCollection `jsonld:"followers,omitempty"`
	// A link to an [ActivityStreams] collection of the actors that follow this actor;
	//  see 5.3 Followers Collection.
	Liked LikedCollection `jsonld:"liked,omitempty"`
	// A short username which may be used to refer to the actor, with no uniqueness guarantees.
	PreferredUsername NaturalLanguageValue `jsonld:"preferredUsername,omitempty,collapsible"`
	// A json object which maps additional (typically server/domain-wide) endpoints which may be useful either
	//  for this actor or someone referencing this actor.
	// This mapping may be nested inside the actor document as the value or may be a link
	//  to a JSON-LD document with these properties.
	Endpoints Endpoints `jsonld:"endpoints,omitempty"`
	// A list of supplementary Collections which may be of interest.
	Streams []Collection `jsonld:"streams,omitempty"`
}

type ActorInterface interface{}

type (
	// Application describes a software application.
	Application Actor

	// Group represents a formal or informal collective of Actors.
	Group Actor

	// Organization represents an organization.
	Organization Actor

	// Person represents an individual person.
	Person Actor

	// Service represents a service of any kind.
	Service Actor
)

// ValidActorType validates the passed type against the valid actor types
func ValidActorType(typ ActivityVocabularyType) bool {
	for _, v := range validActorTypes {
		if v == typ {
			return true
		}
	}
	return false
}

// ActorNew initializes an Actor type actor
func ActorNew(id ObjectID, typ ActivityVocabularyType) *Actor {
	if !ValidActorType(typ) {
		typ = ActorType
	}

	a := Actor{ID: id, Type: typ}
	a.Name = make(NaturalLanguageValue)
	a.Content = make(NaturalLanguageValue)
	a.Summary = make(NaturalLanguageValue)
	in := InboxNew()
	out := OutboxNew()
	liked := LikedNew()

	a.Inbox = InboxStream(*in)
	a.Outbox = OutboxStream(*out)
	a.Liked = LikedCollection(*liked)
	a.PreferredUsername = make(NaturalLanguageValue, 0)

	return &a
}

// ApplicationNew initializes an Application type actor
func ApplicationNew(id ObjectID) *Application {
	a := ActorNew(id, ApplicationType)
	o := Application(*a)
	return &o
}

// GroupNew initializes a Group type actor
func GroupNew(id ObjectID) *Group {
	a := ActorNew(id, GroupType)
	o := Group(*a)
	return &o
}

// OrganizationNew initializes an Organization type actor
func OrganizationNew(id ObjectID) *Organization {
	a := ActorNew(id, OrganizationType)
	o := Organization(*a)
	return &o
}

// PersonNew initializes a Person type actor
func PersonNew(id ObjectID) *Person {
	a := ActorNew(id, PersonType)
	o := Person(*a)
	return &o
}

// ServiceNew initializes a Service type actor
func ServiceNew(id ObjectID) *Service {
	a := ActorNew(id, ServiceType)
	o := Service(*a)
	return &o
}

// IsLink validates if current Actor is a GetLink
func (a Actor) IsLink() bool {
	return a.Type == LinkType || ValidLinkType(a.Type)
}

// IsObject validates if current Actor is an GetID
func (a Actor) IsObject() bool {
	return a.Type == ObjectType || ValidObjectType(a.Type)
}

// GetID returns the GetID corresponding to the Actor object
func (a Actor) GetID() ObjectID {
	return a.ID
}

// GetLink returns the GetLink corresponding to the Actor object
func (a Actor) GetType() ActivityVocabularyType {
	return a.Type
}

// IsLink validates if current Application is a GetLink
func (a Application) IsLink() bool {
	return a.Type == LinkType || ValidLinkType(a.Type)
}

// IsObject validates if current Application is an GetID
func (a Application) IsObject() bool {
	return a.Type == ObjectType || ValidObjectType(a.Type)
}

// GetID returns the GetID corresponding to the Application object
func (a Application) GetID() ObjectID {
	return a.ID
}

// GetLink returns the GetLink corresponding to the Application object
func (a Application) GetType() ActivityVocabularyType {
	return a.Type
}

// IsLink validates if current Group is a GetLink
func (a Group) IsLink() bool {
	return a.Type == LinkType || ValidLinkType(a.Type)
}

// IsObject validates if current Group is an GetID
func (a Group) IsObject() bool {
	return a.Type == ObjectType || ValidObjectType(a.Type)
}

// GetID returns the GetID corresponding to the Group object
func (a Group) GetID() ObjectID {
	return a.ID
}

// GetLink returns the GetLink corresponding to the Group object
func (a Group) GetType() ActivityVocabularyType {
	return a.Type
}

// IsLink validates if current Organization is a GetLink
func (a Organization) IsLink() bool {
	return a.Type == LinkType || ValidLinkType(a.Type)
}

// IsObject validates if current Organization is an GetID
func (a Organization) IsObject() bool {
	return a.Type == ObjectType || ValidObjectType(a.Type)
}

// GetID returns the GetID corresponding to the Organization object
func (a Organization) GetID() ObjectID {
	return a.ID
}

// GetLink returns the GetLink corresponding to the Organization object
func (a Organization) GetType() ActivityVocabularyType {
	return a.Type
}

// IsLink validates if current Service is a GetLink
func (a Service) IsLink() bool {
	return a.Type == LinkType || ValidLinkType(a.Type)
}

// IsObject validates if current Service is an GetID
func (a Service) IsObject() bool {
	return a.Type == ObjectType || ValidObjectType(a.Type)
}

// GetID returns the GetID corresponding to the Service object
func (a Service) GetID() ObjectID {
	return a.ID
}

// GetLink returns the GetLink corresponding to the Service object
func (a Service) GetType() ActivityVocabularyType {
	return a.Type
}
// IsLink validates if current Person is a GetLink
func (p Person) IsLink() bool {
	return p.Type == LinkType || ValidLinkType(p.Type)
}

// IsObject validates if current Person is an GetID
func (p Person) IsObject() bool {
	return p.Type == ObjectType || ValidObjectType(p.Type)
}

// GetID returns the GetID corresponding to the Person object
func (p Person) GetID() ObjectID {
	return p.ID
}

// GetLink returns the GetLink corresponding to the Person object
func (p Person) GetType() ActivityVocabularyType {
	return p.Type
}

func (p Person) GetLink() URI {
	return p.URL.(URI)
}
