package activitypub

type (
	// FollowingCollection is a list of everybody that the actor has followed, added as a side effect.
	// The following collection MUST be either an OrderedCollection or a Collection and MAY
	// be filtered on privileges of an authenticated user or as appropriate when no authentication is given.
	FollowingCollection Following

	// Following is a type alias for a simple Collection
	Following Collection
)
