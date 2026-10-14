package build

import vocab "github.com/go-ap/activitypub"

func OrderedCollection(initFn ...InitFn) *vocab.OrderedCollection {
	initFn = append([]InitFn{WithType(vocab.OrderedCollectionType)}, initFn...)
	ob := new(vocab.OrderedCollection)
	for _, maybeFn := range initFn {
		switch fn := maybeFn.(type) {
		case vocab.WithObjectFn:
			_ = vocab.OnObject(ob, fn)
		case vocab.WithOrderedCollectionFn:
			_ = vocab.OnOrderedCollection(ob, fn)
		}
	}
	return ob
}

func OrderedCollectionPage(initFn ...InitFn) *vocab.OrderedCollectionPage {
	initFn = append([]InitFn{WithType(vocab.OrderedCollectionPageType)}, initFn...)
	ob := new(vocab.OrderedCollectionPage)
	for _, maybeFn := range initFn {
		switch fn := maybeFn.(type) {
		case vocab.WithObjectFn:
			_ = vocab.OnObject(ob, fn)
		case vocab.WithOrderedCollectionFn:
			_ = vocab.OnOrderedCollection(ob, fn)
		case vocab.WithOrderedCollectionPageFn:
			_ = vocab.OnOrderedCollectionPage(ob, fn)
		}
	}
	return ob
}

func Collection(initFn ...InitFn) *vocab.Collection {
	initFn = append([]InitFn{WithType(vocab.CollectionType)}, initFn...)

	ob := new(vocab.Collection)
	for _, maybeFn := range initFn {
		switch fn := maybeFn.(type) {
		case vocab.WithObjectFn:
			_ = vocab.OnObject(ob, fn)
		case vocab.WithCollectionFn:
			_ = vocab.OnCollection(ob, fn)
		}
	}
	return ob
}

func CollectionPage(initFn ...InitFn) *vocab.CollectionPage {
	initFn = append([]InitFn{WithType(vocab.CollectionPageType)}, initFn...)

	ob := new(vocab.CollectionPage)
	for _, maybeFn := range initFn {
		switch fn := maybeFn.(type) {
		case vocab.IRI:
			ob.ID = fn
		case vocab.WithObjectFn:
			_ = vocab.OnObject(ob, fn)
		case vocab.WithCollectionFn:
			_ = vocab.OnCollection(ob, fn)
		case vocab.WithCollectionPageFn:
			_ = vocab.OnCollectionPage(ob, fn)
		}
	}
	return ob
}
