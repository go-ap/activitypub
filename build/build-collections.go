package build

import vocab "github.com/go-ap/activitypub"

func OrderedCollection(initFn ...InitFn) *vocab.OrderedCollection {
	initFn = append([]InitFn{WithType(vocab.OrderedCollectionType)}, initFn...)
	ob := new(vocab.OrderedCollection)
	for _, maybeFn := range initFn {
		switch fn := maybeFn.(type) {
		case vocab.WithObjectFn:
			_ = vocab.OnObject(ob, fn)
		case vocab.WithCollectionFn:
			_ = vocab.OnCollection(ob, fn)
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
		case vocab.WithCollectionFn:
			_ = vocab.OnCollection(ob, fn)
		case vocab.WithOrderedCollectionFn:
			_ = vocab.OnOrderedCollection(ob, fn)
		case vocab.WithCollectionPageFn:
			_ = vocab.OnCollectionPage(ob, fn)
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

func WithCurrent(i ...vocab.Item) vocab.WithCollectionPageFn {
	c := vocab.ItemCollection(i)
	return func(col *vocab.CollectionPage) error {
		col.Current = c.Normalize()
		return nil
	}
}

func WithPartOf(i ...vocab.Item) vocab.WithCollectionPageFn {
	p := vocab.ItemCollection(i)
	return func(col *vocab.CollectionPage) error {
		col.PartOf = p.Normalize()
		return nil
	}
}

func WithPrev(i ...vocab.Item) vocab.WithCollectionPageFn {
	n := vocab.ItemCollection(i)
	return func(col *vocab.CollectionPage) error {
		col.Prev = n.Normalize()
		return nil
	}
}

func WithNext(i ...vocab.Item) vocab.WithCollectionPageFn {
	n := vocab.ItemCollection(i)
	return func(col *vocab.CollectionPage) error {
		col.Next = n.Normalize()
		return nil
	}
}

func WithLast(i ...vocab.Item) vocab.WithCollectionFn {
	l := vocab.ItemCollection(i)
	return func(col *vocab.Collection) error {
		col.Last = l.Normalize()
		return nil
	}
}

func WithFirst(i ...vocab.Item) vocab.WithCollectionFn {
	f := vocab.ItemCollection(i)
	return func(col *vocab.Collection) error {
		col.First = f.Normalize()
		return nil
	}
}

func WithTotalItems(count uint) vocab.WithCollectionFn {
	return func(col *vocab.Collection) error {
		col.TotalItems = count
		return nil
	}
}

func WithItems(it ...vocab.Item) vocab.WithCollectionFn {
	return func(col *vocab.Collection) error {
		return col.Items.Append(it...)
	}
}

func WithOrderedItems(it ...vocab.Item) vocab.WithOrderedCollectionFn {
	return func(col *vocab.OrderedCollection) error {
		return col.OrderedItems.Append(it...)
	}
}
