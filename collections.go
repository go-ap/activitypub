package activitypub

type Collection struct {
	*BaseObject
	Summary    string
	TotalItems int
	Items      ItemCollection
	Current    Page
	First      Page
	Last       Page
	Next       Page
	Prev       Page
}

type OrderedCollection struct {
	*BaseObject
	*CollectionPage
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

//
//
//func Test() {
//	o := &BaseObject{
//		Id: "test",
//	}
//	p := &CollectionPage{Current: "http://localhost"}
//	t := &OrderedCollection{
//		BaseObject:         o,
//		CollectionPage: p,
//		Summary:        "tes",
//	}
//
//	fmt.Print(t)
//}
