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

type Actor struct {
	BaseObject

	Inbox     InboxStream
	Outbox    OutboxStream
	Following FollowingCollection
	Followers FollowersCollection
	Liked     LikedCollection

	PreferredUsername NaturalLanguageValue
	Url               Url
	Summary           NaturalLanguageValue
	Icon              Url

	/*/
	streams
	preferredUsername
	endpoints {
		uploadMedia
		oauthAuthorizationEndpoint
		oauthTokenEndpoint
		provideClientKey
		signClientKey
		sharedInbox
	}
	/**/
}

func ValidActorType(_type string) bool {
	for _, v := range validActorTypes {
		if v == _type {
			return true
		}
	}
	return false
}

func ActorNew(id ObjectId, _type string) Actor {
	if !ValidActorType(_type) {
		_type = ActorType
	}
	o := BaseObject{Id: id, Type: _type}

	return Actor{BaseObject: o}
}

func ApplicationNew(id ObjectId) Actor {
	return ActorNew(id, ApplicationType)
}

func GroupNew(id ObjectId) Actor {
	return ActorNew(id, GroupType)
}

func OrganizationNew(id ObjectId) Actor {
	return ActorNew(id, OrganizationType)
}

func PersonNew(id ObjectId) Actor {
	return ActorNew(id, PersonType)
}

func ServiceNew(id ObjectId) Actor {
	return ActorNew(id, ServiceType)
}
