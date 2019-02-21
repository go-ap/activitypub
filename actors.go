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

// Actor is generally one of the ActivityStreams Actor Types, but they don't have to be.
// For example, a Profile object might be used as an actor, or a type from an ActivityStreams extension.
// Actors are retrieved like any other Object in ActivityPub.
// Like other ActivityStreams objects, actors have an id, which is a URI.
type Actor Item

type (
	// Application describes a software application.
	Application = Object

	// Group represents a formal or informal collective of Actors.
	Group = Object

	// Organization represents an organization.
	Organization = Object

	// Person represents an individual person.
	Person = Object

	// Service represents a service of any kind.
	Service = Object
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
	a.Name = NaturalLanguageValuesNew()
	a.Content = NaturalLanguageValuesNew()
	a.Summary = NaturalLanguageValuesNew()

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
