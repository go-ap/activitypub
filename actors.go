package activitypub

import as "github.com/go-ap/activitystreams"

// Endpoints a json object which maps additional (typically server/domain-wide)
// endpoints which may be useful either for this actor or someone referencing this actor.
// This mapping may be nested inside the actor document as the value or may be a link to
// a JSON-LD document with these properties.
type Endpoints struct {
	// UploadMedia Upload endpoint URI for this user for binary data.
	UploadMedia as.Item `jsonld:"uploadMedia,omitempty"`
	// OauthAuthorizationEndpoint Endpoint URI so this actor's clients may access remote ActivityStreams objects which require authentication
	// to access. To use this endpoint, the client posts an x-www-form-urlencoded id parameter with the value being
	// the id of the requested ActivityStreams object.
	OauthAuthorizationEndpoint as.Item `jsonld:"oauthAuthorizationEndpoint,omitempty"`
	// OauthTokenEndpoint If OAuth 2.0 bearer tokens [RFC6749] [RFC6750] are being used for authenticating client to server interactions,
	// this endpoint specifies a URI at which a browser-authenticated user may obtain a new authorization grant.
	OauthTokenEndpoint as.Item `jsonld:"oauthTokenEndpoint,omitempty"`
	// ProvideClientKey  If OAuth 2.0 bearer tokens [RFC6749] [RFC6750] are being used for authenticating client to server interactions,
	// this endpoint specifies a URI at which a client may acquire an access token.
	ProvideClientKey as.Item `jsonld:"provideClientKey,omitempty"`
	// SignClientKey If Linked Data Signatures and HTTP Signatures are being used for authentication and authorization,
	// this endpoint specifies a URI at which browser-authenticated users may authorize a client's public
	// key for client to server interactions.
	SignClientKey as.Item `jsonld:"signClientKey,omitempty"`
	// SharedInbox If Linked Data Signatures and HTTP Signatures are being used for authentication and authorization,
	// this endpoint specifies a URI at which a client key may be signed by the actor's key for a time window to
	// act on behalf of the actor in interacting with foreign servers.
	SharedInbox as.Item `jsonld:"sharedInbox,omitempty"`
}

// Actor is generally one of the ActivityStreams Actor Types, but they don't have to be.
// For example, a Profile object might be used as an actor, or a type from an ActivityStreams extension.
// Actors are retrieved like any other Object in ActivityPub.
// Like other ActivityStreams objects, actors have an id, which is a URI.
type Actor struct {
	as.Parent
	// A reference to an [ActivityStreams] OrderedCollection comprised of all the messages received by the actor;
	// see 5.2 Inbox.
	Inbox as.Item `jsonld:"inbox,omitempty"`
	// An [ActivityStreams] OrderedCollection comprised of all the messages produced by the actor;
	// see 5.1 Outbox.
	Outbox as.Item `jsonld:"outbox,omitempty"`
	// A link to an [ActivityStreams] collection of the actors that this actor is following;
	// see 5.4 Following Collection
	Following as.Item `jsonld:"following,omitempty"`
	// A link to an [ActivityStreams] collection of the actors that follow this actor;
	// see 5.3 Followers Collection.
	Followers as.Item `jsonld:"followers,omitempty"`
	// A link to an [ActivityStreams] collection of the actors that follow this actor;
	// see 5.3 Followers Collection.
	Liked as.Item `jsonld:"liked,omitempty"`
	// A short username which may be used to refer to the actor, with no uniqueness guarantees.
	PreferredUsername as.NaturalLanguageValue `jsonld:"preferredUsername,omitempty,collapsible"`
	// A json object which maps additional (typically server/domain-wide) endpoints which may be useful either
	// for this actor or someone referencing this actor.
	// This mapping may be nested inside the actor document as the value or may be a link
	// to a JSON-LD document with these properties.
	Endpoints Endpoints `jsonld:"endpoints,omitempty"`
	// A list of supplementary Collections which may be of interest.
	Streams []as.ItemCollection `jsonld:"streams,omitempty"`
}

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

// ActorNew initializes an Actor type actor
func ActorNew(id as.ObjectID, typ as.ActivityVocabularyType) *Actor {
	if !as.ValidActorType(typ) {
		typ = as.ActorType
	}

	a := Actor{Parent: as.Object {ID: id, Type: typ}}
	a.Name = as.NaturalLanguageValueNew()
	a.Content = as.NaturalLanguageValueNew()
	a.Summary = as.NaturalLanguageValueNew()
	in := as.OrderedCollectionNew(as.ObjectID("test-inbox"))
	out := as.OrderedCollectionNew(as.ObjectID("test-outbox"))
	liked := as.OrderedCollectionNew(as.ObjectID("test-liked"))

	a.Inbox = in
	a.Outbox = out
	a.Liked = liked
	a.PreferredUsername = as.NaturalLanguageValueNew()

	return &a
}

// ApplicationNew initializes an Application type actor
func ApplicationNew(id as.ObjectID) *Application {
	a := ActorNew(id, as.ApplicationType)
	o := Application(*a)
	return &o
}

// GroupNew initializes a Group type actor
func GroupNew(id as.ObjectID) *Group {
	a := ActorNew(id, as.GroupType)
	o := Group(*a)
	return &o
}

// OrganizationNew initializes an Organization type actor
func OrganizationNew(id as.ObjectID) *Organization {
	a := ActorNew(id, as.OrganizationType)
	o := Organization(*a)
	return &o
}

// PersonNew initializes a Person type actor
func PersonNew(id as.ObjectID) *Person {
	a := ActorNew(id, as.PersonType)
	o := Person(*a)
	return &o
}

// ServiceNew initializes a Service type actor
func ServiceNew(id as.ObjectID) *Service {
	a := ActorNew(id, as.ServiceType)
	o := Service(*a)
	return &o
}

// IsLink validates if current Actor is a Link
func (a Actor) IsLink() bool {
	return a.Type == as.LinkType || as.ValidLinkType(a.Type)
}

// IsObject validates if current Actor is an Object
func (a Actor) IsObject() bool {
	return a.Type == as.ObjectType || as.ValidObjectType(a.Type)
}

// GetID returns the ObjectID corresponding to the Actor object
func (a Actor) GetID() *as.ObjectID {
	return &a.ID
}

// GetLink returns the IRI corresponding to the Actor object
func (a Actor) GetLink() as.IRI {
	return as.IRI(a.ID)
}

// GetType returns the type corresponding to the Actor object
func (a Actor) GetType() as.ActivityVocabularyType {
	return a.Type
}

// IsLink validates if current Application is a Link
func (a Application) IsLink() bool {
	return a.Type == as.LinkType || as.ValidLinkType(a.Type)
}

// IsObject validates if current Application is an Object
func (a Application) IsObject() bool {
	return a.Type == as.ObjectType || as.ValidObjectType(a.Type)
}

// GetID returns the ObjectID corresponding to the  Application object
func (a Application) GetID() *as.ObjectID {
	return a.GetActor().GetID()
}

// GetLink returns the IRI corresponding to the Application object
func (a Application) GetLink() as.IRI {
	return as.IRI(a.ID)
}

// GetType returns the type corresponding to the Application object
func (a Application) GetType() as.ActivityVocabularyType {
	return a.Type
}

// IsLink validates if current Group is a Link
func (g Group) IsLink() bool {
	return g.Type == as.LinkType || as.ValidLinkType(g.Type)
}

// IsObject validates if current Group is an Object
func (g Group) IsObject() bool {
	return g.Type == as.ObjectType || as.ValidObjectType(g.Type)
}

// GetID returns the ObjectID corresponding to the  Group object
func (g Group) GetID() *as.ObjectID {
	return g.GetActor().GetID()
}

// GetLink returns the IRI corresponding to the Group object
func (g Group) GetLink() as.IRI {
	return as.IRI(g.ID)
}

// GetType returns the type corresponding to the Group object
func (g Group) GetType() as.ActivityVocabularyType {
	return g.Type
}

// IsLink validates if current Organization is a Link
func (o Organization) IsLink() bool {
	return o.Type == as.LinkType || as.ValidLinkType(o.Type)
}

// IsObject validates if current Organization is an Object
func (o Organization) IsObject() bool {
	return o.Type == as.ObjectType || as.ValidObjectType(o.Type)
}

// GetID returns the ObjectID corresponding to the  Organization object
func (o Organization) GetID() *as.ObjectID {
	return o.GetActor().GetID()
}

// GetLink returns the IRI corresponding to the Organization object
func (o Organization) GetLink() as.IRI {
	return as.IRI(o.ID)
}

// GetType returns the type corresponding to the Organization object
func (o Organization) GetType() as.ActivityVocabularyType {
	return o.Type
}

// IsLink validates if current Service is a Link
func (s Service) IsLink() bool {
	return s.Type == as.LinkType || as.ValidLinkType(s.Type)
}

// IsObject validates if current Service is an Object
func (s Service) IsObject() bool {
	return s.Type == as.ObjectType || as.ValidObjectType(s.Type)
}

// GetID returns the ObjectID corresponding to the Service object
func (s Service) GetID() *as.ObjectID {
	return s.GetActor().GetID()
}

// GetLink returns the IRI corresponding to the Service object
func (s Service) GetLink() as.IRI {
	return as.IRI(s.ID)
}

// GetType returns the type corresponding to the Service object
func (s Service) GetType() as.ActivityVocabularyType {
	return s.Type
}

// IsLink validates if current Person is a Link
func (p Person) IsLink() bool {
	return p.Type == as.LinkType || as.ValidLinkType(p.Type)
}

// IsObject validates if current Person is an Object
func (p Person) IsObject() bool {
	return p.Type == as.ObjectType || as.ValidObjectType(p.Type)
}

// GetID returns the ObjectID corresponding to the Person object
func (p Person) GetID() *as.ObjectID {
	return p.GetActor().GetID()
}

// GetType returns the object type for the current Person object
func (p Person) GetType() as.ActivityVocabularyType {
	return p.Type
}

// GetLink returns the IRI corresponding to the Person object
func (p Person) GetLink() as.IRI {
	return as.IRI(p.ID)
}

// UnmarshalJSON
func (a *Actor) UnmarshalJSON(data []byte) error {
	a.Parent.UnmarshalJSON(data)

	a.PreferredUsername = as.JSONGetNaturalLanguageField(data, "preferredUsername")

	out := as.JSONGetItem(data, "outbox")
	if out != nil {
		a.Outbox = out
	}
	inb := as.JSONGetItem(data, "inbox")
	if inb != nil {
		a.Inbox = inb
	}
	followers := as.JSONGetItem(data, "followers")
	if followers != nil {
		a.Followers = followers
	}
	following := as.JSONGetItem(data, "following")
	if following != nil {
		a.Following = following
	}
	liked := as.JSONGetItem(data, "liked")
	if liked != nil {
		a.Liked = liked
	}
	//streams := as.JSONGetItems(data, "streams")
	//if streams != nil {
	//	a.Streams = streams
	//}
	// @todo(marius) : Add as.JSONGetIEndPoints

	return nil
}

func (p *Person) UnmarshalJSON(data []byte) error {
	a := p.GetActor()
	err := a.UnmarshalJSON(data)

	*p = Person(a)

	return err
}

// UnmarshalJSON
func (a *Application) UnmarshalJSON(data []byte) error {
	act := a.GetActor()
	err := act.UnmarshalJSON(data)

	*a = Application(act)

	return err
}

// UnmarshalJSON
func (g *Group) UnmarshalJSON(data []byte) error {
	a := g.GetActor()
	err := a.UnmarshalJSON(data)

	*g = Group(a)

	return err
}

// UnmarshalJSON
func (o *Organization) UnmarshalJSON(data []byte) error {
	a := o.GetActor()
	err := a.UnmarshalJSON(data)

	*o = Organization(a)

	return err
}

// UnmarshalJSON
func (s *Service) UnmarshalJSON(data []byte) error {
	a := s.GetActor()
	err := a.UnmarshalJSON(data)

	*s = Service(a)

	return err
}

// GetActor returns the underlying Actor type
func (a Actor) GetActor() Actor {
	return a
}

// GetActor returns the underlying Actor type
func (a Application) GetActor() Actor {
	return Actor(a)
}

// GetActor returns the underlying Actor type
func (g Group) GetActor() Actor {
	return Actor(g)
}

// GetActor returns the underlying Actor type
func (o Organization) GetActor() Actor {
	return Actor(o)
}

// GetActor returns the underlying Actor type
func (p Person) GetActor() Actor {
	return Actor(p)
}

// GetActor returns the underlying Actor type
func (s Service) GetActor() Actor {
	return Actor(s)
}
