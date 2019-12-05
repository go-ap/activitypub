package activitypub

type (
	// SharesCollection is a list of all Announce activities with this object as the object property,
	// added as a side effect. The shares collection MUST be either an OrderedCollection or a Collection
	// and MAY be filtered on privileges of an authenticated user or as appropriate when no authentication
	// is given.
	SharesCollection = Shares

	// Shares is a type alias for an Ordered Collection
	Shares = OrderedCollection
)

// SharesNew initializes a new Shares
func SharesNew() *Shares {
	id := ID("Shares")

	i := Shares{ID: id, Type: CollectionType}
	i.Name = NaturalLanguageValuesNew()
	i.Content = NaturalLanguageValuesNew()

	i.TotalItems = 0

	return &i
}
