package activitypub

type (
	// SharesCollection is a list of all Announce activities with this object as the object property,
	// added as a side effect. The shares collection MUST be either an OrderedCollection or a Collection
	// and MAY be filtered on privileges of an authenticated user or as appropriate when no authentication
	// is given.
	SharesCollection Shares

	// Shares is a type alias for an Ordered Collection
	Shares OrderedCollection
)
