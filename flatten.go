package activitypub

// FlattenActivityProperties flattens the Activity's properties from Object type to IRI
func FlattenActivityProperties(act *Activity) *Activity {
	if act == nil {
		return nil
	}
	_ = OnIntransitiveActivity(act, func(in *IntransitiveActivity) error {
		FlattenIntransitiveActivityProperties(in)
		return nil
	})
	act.Object = FlattenToIRI(act.Object)
	return act
}

// FlattenIntransitiveActivityProperties flattens the Activity's properties from Object type to IRI
func FlattenIntransitiveActivityProperties(act *IntransitiveActivity) *IntransitiveActivity {
	if act == nil {
		return nil
	}
	act.Actor = FlattenToIRI(act.Actor)
	act.Target = FlattenToIRI(act.Target)
	act.Result = FlattenToIRI(act.Result)
	act.Origin = FlattenToIRI(act.Origin)
	act.Result = FlattenToIRI(act.Result)
	act.Instrument = FlattenToIRI(act.Instrument)
	_ = OnObject(act, func(o *Object) error {
		FlattenObjectProperties(o)
		return nil
	})
	return act
}

// FlattenItemCollection flattens an Item Collection to their respective IRIs
func FlattenItemCollection(col ItemCollection) ItemCollection {
	if col == nil {
		return col
	}
	for k, it := range ItemCollectionDeduplication(&col) {
		if iri := it.GetLink(); iri != "" {
			col[k] = iri
		}
	}
	return col
}

// FlattenCollection flattens a Collection's objects to their respective IRIs
func FlattenCollection(col *Collection) *Collection {
	if col == nil {
		return col
	}
	col.Items = FlattenItemCollection(col.Items)

	return col
}

// FlattenOrderedCollection flattens an OrderedCollection's objects to their respective IRIs
func FlattenOrderedCollection(col *OrderedCollection) *OrderedCollection {
	if col == nil {
		return col
	}
	col.OrderedItems = FlattenItemCollection(col.OrderedItems)

	return col
}

// FlattenActorProperties flattens the Actor's properties from Object types to IRI
func FlattenActorProperties(a *Actor) *Actor {
	if a == nil {
		return nil
	}
	OnObject(a, func(o *Object) error {
		FlattenObjectProperties(o)
		return nil
	})
	return a
}

// FlattenObjectProperties flattens the Object's properties from Object types to IRI
func FlattenObjectProperties(o *Object) *Object {
	if o == nil {
		return nil
	}
	o.Replies = Flatten(o.Replies)
	o.Shares = Flatten(o.Shares)
	o.Likes = Flatten(o.Likes)
	o.AttributedTo = Flatten(o.AttributedTo)
	o.To = FlattenItemCollection(o.To)
	o.Bto = FlattenItemCollection(o.Bto)
	o.CC = FlattenItemCollection(o.CC)
	o.BCC = FlattenItemCollection(o.BCC)
	o.Audience = FlattenItemCollection(o.Audience)
	// o.Tag = FlattenItemCollection(o.Tag)
	return o
}

// FlattenProperties flattens the Item's properties from Object types to IRI
func FlattenProperties(it Item) Item {
	if IsNil(it) {
		return nil
	}
	typ := it.GetType()
	if IntransitiveActivityTypes.Contains(typ) {
		_ = OnIntransitiveActivity(it, func(a *IntransitiveActivity) error {
			FlattenIntransitiveActivityProperties(a)
			return nil
		})
	} else if ActivityTypes.Contains(typ) {
		_ = OnActivity(it, func(a *Activity) error {
			FlattenActivityProperties(a)
			return nil
		})
	}
	if ActorTypes.Contains(typ) {
		OnActor(it, func(a *Actor) error {
			FlattenActorProperties(a)
			return nil
		})
	}
	if ObjectTypes.Contains(typ) {
		OnObject(it, func(o *Object) error {
			FlattenObjectProperties(o)
			return nil
		})
	}
	return it
}

// Flatten checks if Item can be flattened to an IRI or array of IRIs and returns it if so
func Flatten(it Item) Item {
	if IsNil(it) {
		return nil
	}
	if it.IsCollection() {
		OnCollectionIntf(it, func(c CollectionInterface) error {
			it = FlattenItemCollection(c.Collection()).Normalize()
			return nil
		})
		return it
	}
	return it.GetLink()
}
