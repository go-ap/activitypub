package activitystreams

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
	// SharedInbox If Linked Data Signatures and HTTP Signatures are being used for authentication and authorization,
	// this endpoint specifies a URI at which a client key may be signed by the actor's key for a time window to
	// act on behalf of the actor in interacting with foreign servers.
	SharedInbox Item `jsonld:"sharedInbox,omitempty"`
}

type WillAct interface {
	GetActor() Actor
}

// Actor is generally one of the ActivityStreams Actor Types, but they don't have to be.
// For example, a Profile object might be used as an actor, or a type from an ActivityStreams extension.
// Actors are retrieved like any other Object in ActivityPub.
// Like other ActivityStreams objects, actors have an id, which is a URI.
type Actor Item

type (
	// Application describes a software application.
	Application Object

	// Group represents a formal or informal collective of Actors.
	Group Object

	// Organization represents an organization.
	Organization Object

	// Person represents an individual person.
	Person Object

	// Service represents a service of any kind.
	Service Object
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
func ActorNew(id ObjectID, typ ActivityVocabularyType) *Object {
	if !ValidActorType(typ) {
		typ = ActorType
	}

	a := Object{ID: id, Type: typ}
	a.Name = NaturalLanguageValueNew()
	a.Content = NaturalLanguageValueNew()
	a.Summary = NaturalLanguageValueNew()

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

// IsLink validates if current Application is a Link
func (a Application) IsLink() bool {
	return a.Type == LinkType || ValidLinkType(a.Type)
}

// IsObject validates if current Application is an Object
func (a Application) IsObject() bool {
	return a.Type == ObjectType || ValidObjectType(a.Type)
}

// GetID returns the ObjectID corresponding to the  Application object
func (a Application) GetID() *ObjectID {
	return &a.ID
}

// GetLink returns the IRI corresponding to the Application object
func (a Application) GetLink() IRI {
	return IRI(a.ID)
}

// GetType returns the type corresponding to the Application object
func (a Application) GetType() ActivityVocabularyType {
	return a.Type
}

// IsLink validates if current Group is a Link
func (g Group) IsLink() bool {
	return g.Type == LinkType || ValidLinkType(g.Type)
}

// IsObject validates if current Group is an Object
func (g Group) IsObject() bool {
	return g.Type == ObjectType || ValidObjectType(g.Type)
}

// GetID returns the ObjectID corresponding to the  Group object
func (g Group) GetID() *ObjectID {
	return &g.ID
}

// GetLink returns the IRI corresponding to the Group object
func (g Group) GetLink() IRI {
	return IRI(g.ID)
}

// GetType returns the type corresponding to the Group object
func (g Group) GetType() ActivityVocabularyType {
	return g.Type
}

// IsLink validates if current Organization is a Link
func (o Organization) IsLink() bool {
	return o.Type == LinkType || ValidLinkType(o.Type)
}

// IsObject validates if current Organization is an Object
func (o Organization) IsObject() bool {
	return o.Type == ObjectType || ValidObjectType(o.Type)
}

// GetID returns the ObjectID corresponding to the  Organization object
func (o Organization) GetID() *ObjectID {
	return &o.ID
}

// GetLink returns the IRI corresponding to the Organization object
func (o Organization) GetLink() IRI {
	return IRI(o.ID)
}

// GetType returns the type corresponding to the Organization object
func (o Organization) GetType() ActivityVocabularyType {
	return o.Type
}

// IsLink validates if current Service is a Link
func (s Service) IsLink() bool {
	return s.Type == LinkType || ValidLinkType(s.Type)
}

// IsObject validates if current Service is an Object
func (s Service) IsObject() bool {
	return s.Type == ObjectType || ValidObjectType(s.Type)
}

// GetID returns the ObjectID corresponding to the Service object
func (s Service) GetID() *ObjectID {
	return &s.ID
}

// GetLink returns the IRI corresponding to the Service object
func (s Service) GetLink() IRI {
	return IRI(s.ID)
}

// GetType returns the type corresponding to the Service object
func (s Service) GetType() ActivityVocabularyType {
	return s.Type
}

// IsLink validates if current Person is a Link
func (p Person) IsLink() bool {
	return p.Type == LinkType || ValidLinkType(p.Type)
}

// IsObject validates if current Person is an Object
func (p Person) IsObject() bool {
	return p.Type == ObjectType || ValidObjectType(p.Type)
}

// GetID returns the ObjectID corresponding to the Person object
func (p Person) GetID() *ObjectID {
	return &p.ID
}

// GetType returns the object type for the current Person object
func (p Person) GetType() ActivityVocabularyType {
	return p.Type
}

// GetLink returns the IRI corresponding to the Person object
func (p Person) GetLink() IRI {
	return IRI(p.ID)
}

func (p *Person) UnmarshalJSON(data []byte) error {
	a := Object(*p)
	err := a.UnmarshalJSON(data)

	*p = Person(a)

	return err
}

// UnmarshalJSON
func (a *Application) UnmarshalJSON(data []byte) error {
	act := Object(*a)
	err := act.UnmarshalJSON(data)

	*a = Application(act)

	return err
}

// UnmarshalJSON
func (g *Group) UnmarshalJSON(data []byte) error {
	a := Object(*g)
	err := a.UnmarshalJSON(data)

	*g = Group(a)

	return err
}

// UnmarshalJSON
func (o *Organization) UnmarshalJSON(data []byte) error {
	a := Object(*o)
	err := a.UnmarshalJSON(data)

	*o = Organization(a)

	return err
}

// UnmarshalJSON
func (s *Service) UnmarshalJSON(data []byte) error {
	a := Object(*s)
	err := a.UnmarshalJSON(data)

	*s = Service(a)
	return err
}

// GetActor returns the underlying Actor type
func (a Application) GetActor() Actor {
	return a
}

// GetActor returns the underlying Actor type
func (g Group) GetActor() Actor {
	return g
}

// GetActor returns the underlying Actor type
func (o Organization) GetActor() Actor {
	return o
}

// GetActor returns the underlying Actor type
func (p Person) GetActor() Actor {
	return p
}

// GetActor returns the underlying Actor type
func (s Service) GetActor() Actor {
	return s
}
