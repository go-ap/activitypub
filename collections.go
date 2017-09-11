package activitypub

type Collection struct {
	BaseObject
	CollectionPage
	Summary    string
	TotalItems int
	Items      ItemCollection
}

type OrderedCollection struct {
	BaseObject
	CollectionPage
	Summary      string
	TotalItems   int
	OrderedItems ItemCollection
}

type Page Url

type CollectionPage struct {
	Current Page
	First   Page
	Last    Page
	Next    Page
	Prev    Page
}

func CollectionNew(id ObjectId) Collection {
	o := BaseObject{Id: id, Type: CollectionType}

	return Collection{BaseObject:o}
}
func OrderedCollectionNew(id ObjectId) OrderedCollection {
	o := BaseObject{Id: id, Type: OrderedCollectionType}

	return OrderedCollection{BaseObject:o}
}