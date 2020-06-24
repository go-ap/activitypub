package activitypub

// Item struct
type Item = ObjectOrLink

const (
	EmptyIRI IRI = ""
	NilIRI   IRI = "-"

	EmptyID = EmptyIRI
	NilID   = NilIRI
)

// Flatten checks if Item can be flatten to an IRI or array of IRIs and returns it if so
func Flatten(it Item) Item {
	if it.IsCollection() {
		if c, ok := it.(CollectionInterface); ok {
			it = FlattenItemCollection(c.Collection())
		}
	}
	if it != nil && len(it.GetLink()) > 0 {
		return it.GetLink()
	}
	return it
}

// ItemsEqual checks if it and with Items are equal
func ItemsEqual(it, with Item) bool {
	if it == nil || with == nil{
		return with == it
	}
	result := true
	if it.IsCollection() {
		if it.GetType() == CollectionOfItems {
			OnItemCollection(it, func(c *ItemCollection) error {
				result = c.Equals(with)
				return nil
			})
		}
		if it.GetType() == CollectionType {
			OnCollection(it, func(c *Collection) error {
				result = c.Equals(with)
				return nil
			})
		}
		if it.GetType() == OrderedCollectionType {
			OnOrderedCollection(it, func(c *OrderedCollection) error {
				result = c.Equals(with)
				return nil
			})
		}
		if it.GetType() == CollectionPageType {
			OnCollectionPage(it, func(c *CollectionPage) error {
				result = c.Equals(with)
				return nil
			})
		}
		if it.GetType() == OrderedCollectionPageType {
			OnOrderedCollectionPage(it, func(c *OrderedCollectionPage) error {
				result = c.Equals(with)
				return nil
			})
		}
	}
	if it.IsObject() {
		if ActivityTypes.Contains(with.GetType()) {
			OnActivity(it, func(i*Activity) error {
				result = i.Equals(with)
				return nil
			})
		} else if ActorTypes.Contains(with.GetType()) {
			OnActor(it, func(i *Actor) error {
				result = i.Equals(with)
				return nil
			})
		} else {
			OnObject(it, func(i *Object) error {
				result = i.Equals(with)
				return nil
			})
		}
	}
	if with.IsLink() {
		result = with.GetLink().Equals(it.GetLink(), false)
	}
	return result
}
