package activitypub

type (
	// LikesCollection is a list of all Like activities with this object as the object property,
	// added as a side effect. The likes collection MUST be either an OrderedCollection or a Collection
	// and MAY be filtered on privileges of an authenticated user or as appropriate when
	// no authentication is given.
	LikesCollection Likes

	// Likes is a type alias for an Ordered Collection
	Likes OrderedCollection
)
