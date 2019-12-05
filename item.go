package activitypub

// Item struct
type Item ObjectOrLink

const EmptyID = ID("")
const EmptyIRI = IRI("")

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
