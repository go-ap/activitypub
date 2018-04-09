package activitypub

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
// Actors are retrieved like any other Object in ActivityPub.
// Like other ActivityStreams objects, actors have an id, which is a URI.
type Actor struct {
	apObject
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
func ValidActorType(_type ActivityVocabularyType) bool {
	for _, v := range validActorTypes {
		if v == _type {
			return true
		}
	}
	return false
}

// ActorNew initializes an Actor type actor
func ActorNew(id ObjectID, _type ActivityVocabularyType) *Actor {
	if !ValidActorType(_type) {
		_type = ActorType
	}
	o := ObjectNew(id, _type)
	a := Actor{apObject: o}

	in := InboxNew()

	a.Inbox = InboxStream(*in)
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

// IsLink validates if current Actor is a Link
func (a Actor) IsLink() bool {
	return a.Type == LinkType || ValidLinkType(a.Type)
}

// IsObject validates if current Actor is an Object
func (a Actor) IsObject() bool {
	return a.Type == ObjectType || ValidObjectType(a.Type)
}

// Object returns the apObject corresponding to the Actor object
func (a Actor) Object() apObject {
	return a.apObject
}

// Link returns the Link corresponding to the Actor object
func (a Actor) Link() Link {
	return Link{}
}

// IsLink validates if current Person is a Link
func (p Person) IsLink() bool {
	return p.Type == LinkType || ValidLinkType(p.Type)
}

// IsObject validates if current Person is an Object
func (p Person) IsObject() bool {
	return p.Type == ObjectType || ValidObjectType(p.Type)
}

// Object returns the apObject corresponding to the Person object
func (p Person) Object() apObject {
	return p.apObject
}

// Link returns the Link corresponding to the Person object
func (p Person) Link() Link {
	return Link{}
}
