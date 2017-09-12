package activitypub

const (
	// Actor Types
	ApplicationType  string = "Application"
	GroupType        string = "Group"
	OrganizationType string = "Organization"
	PersonType       string = "Person"
	ServiceType      string = "Service"
)

var validActorTypes = [...]string{
	ApplicationType,
	GroupType,
	OrganizationType,
	PersonType,
	ServiceType,
}

type Endpoints struct {
	// Upload endpoint URI for this user for binary data.
	UploadMedia					ObjectOrLink	`jsonld:"uploadMedia,omitempty"`
	// Endpoint URI so this actor's clients may access remote ActivityStreams objects which require authentication
	//  to access. To use this endpoint, the client posts an x-www-form-urlencoded id parameter with the value being
	//  the id of the requested ActivityStreams object.
	OauthAuthorizationEndpoint	ObjectOrLink	`jsonld:"oauthAuthorizationEndpoint,omitempty"`
	// If OAuth 2.0 bearer tokens [RFC6749] [RFC6750] are being used for authenticating client to server interactions,
	//  this endpoint specifies a URI at which a browser-authenticated user may obtain a new authorization grant.
	OauthTokenEndpoint			ObjectOrLink	`jsonld:"oauthTokenEndpoint,omitempty"`
	// If OAuth 2.0 bearer tokens [RFC6749] [RFC6750] are being used for authenticating client to server interactions,
	//  this endpoint specifies a URI at which a client may acquire an access token.
	ProvideClientKey			ObjectOrLink	`jsonld:"provideClientKey,omitempty"`
	// If Linked Data Signatures and HTTP Signatures are being used for authentication and authorization,
	//  this endpoint specifies a URI at which browser-authenticated users may authorize a client's public
	//  key for client to server interactions.
	SignClientKey				ObjectOrLink	`jsonld:"signClientKey,omitempty"`
	// If Linked Data Signatures and HTTP Signatures are being used for authentication and authorization,
	//  this endpoint specifies a URI at which a client key may be signed by the actor's key for a time window to
	//  act on behalf of the actor in interacting with foreign servers.
	SharedInbox					ObjectOrLink	`jsonld:"sharedInbox,omitempty"`
}

// Actor types are Object types that are capable of performing activities.
type Actor struct {
	*BaseObject
	// A reference to an [ActivityStreams] OrderedCollection comprised of all the messages received by the actor;
	//  see 5.2 Inbox.
	Inbox     			InboxStream				`jsonld:"inbox,omitempty"`
	// An [ActivityStreams] OrderedCollection comprised of all the messages produced by the actor;
	//  see 5.1 Outbox.
	Outbox    			OutboxStream			`jsonld:"outbox,omitempty"`
	// A link to an [ActivityStreams] collection of the actors that this actor is following;
	//  see 5.4 Following Collection
	Following 			FollowingCollection		`jsonld:"following,omitempty"`
	// A link to an [ActivityStreams] collection of the actors that follow this actor;
	//  see 5.3 Followers Collection.
	Followers 			FollowersCollection		`jsonld:"followers,omitempty"`
	// A link to an [ActivityStreams] collection of the actors that follow this actor;
	//  see 5.3 Followers Collection.
	Liked     			LikedCollection			`jsonld:"liked,omitempty"`
	// A short username which may be used to refer to the actor, with no uniqueness guarantees.
	PreferredUsername 	NaturalLanguageValue	`jsonld:"preferredUsername,omitempty"`
	// A json object which maps additional (typically server/domain-wide) endpoints which may be useful either
	//  for this actor or someone referencing this actor.
	// This mapping may be nested inside the actor document as the value or may be a link
	//  to a JSON-LD document with these properties.
	Endpoints			Endpoints				`jsonld:"endpoints,omitempty"`
	// A list of supplementary Collections which may be of interest.
	Streams 			[]Collection			`jsonld:"streams,omitempty"`
}

// Describes a software application.
type Application Actor

// Represents a formal or informal collective of Actors.
type Group Actor

// Represents an organization.
type Organization Actor

// Represents an individual person.
type Person Actor

// Represents a service of any kind.
type Service Actor

func ValidActorType(_type string) bool {
	for _, v := range validActorTypes {
		if v == _type {
			return true
		}
	}
	return false
}

func ActorNew(id ObjectId, _type string) *Actor {
	if !ValidActorType(_type) {
		_type = ActorType
	}
	o := BaseObject{Id: id, Type: _type}

	return &Actor{BaseObject: &o}
}

func ApplicationNew(id ObjectId) *Application {
	a := ActorNew(id, ApplicationType)
	o := Application(*a)
	return &o
}

func GroupNew(id ObjectId) *Group {
	a := ActorNew(id, GroupType)
	o := Group(*a)
	return &o
}

func OrganizationNew(id ObjectId) *Organization {
	a := ActorNew(id, OrganizationType)
	o := Organization(*a)
	return &o
}

func PersonNew(id ObjectId) *Person {
	a := ActorNew(id, PersonType)
	o := Person(*a)
	return &o
}

func ServiceNew(id ObjectId) *Service {
	a := ActorNew(id, ServiceType)
	o := Service(*a)
	return &o
}
