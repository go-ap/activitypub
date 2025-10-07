package activitypub

import "fmt"

func CopyOrderedCollectionPageProperties(to, from *OrderedCollectionPage) (*OrderedCollectionPage, error) {
	to.PartOf = replaceIfItem(to.PartOf, from.PartOf)
	to.Next = replaceIfItem(to.Next, from.Next)
	to.Prev = replaceIfItem(to.Prev, from.Prev)
	oldCol, _ := ToOrderedCollection(to)
	newCol, _ := ToOrderedCollection(from)
	_, err := CopyOrderedCollectionProperties(oldCol, newCol)
	if err != nil {
		return to, err
	}
	return to, nil
}

func CopyCollectionPageProperties(to, from *CollectionPage) (*CollectionPage, error) {
	to.PartOf = replaceIfItem(to.PartOf, from.PartOf)
	to.Next = replaceIfItem(to.Next, from.Next)
	to.Prev = replaceIfItem(to.Prev, from.Prev)
	toCol, _ := ToCollection(to)
	fromCol, _ := ToCollection(from)
	_, err := CopyCollectionProperties(toCol, fromCol)
	return to, err
}

func CopyOrderedCollectionProperties(to, from *OrderedCollection) (*OrderedCollection, error) {
	to.First = replaceIfItem(to.First, from.First)
	to.Last = replaceIfItem(to.Last, from.Last)
	to.OrderedItems = replaceIfItemCollection(to.OrderedItems, from.OrderedItems)
	if to.TotalItems == 0 {
		to.TotalItems = from.TotalItems
	}
	oldOb, _ := ToObject(to)
	newOb, _ := ToObject(from)
	_, err := CopyObjectProperties(oldOb, newOb)
	return to, err
}

func CopyCollectionProperties(to, from *Collection) (*Collection, error) {
	to.First = replaceIfItem(to.First, from.First)
	to.Last = replaceIfItem(to.Last, from.Last)
	to.Items = replaceIfItemCollection(to.Items, from.Items)
	if to.TotalItems == 0 {
		to.TotalItems = from.TotalItems
	}
	oldOb, _ := ToObject(to)
	newOb, _ := ToObject(from)
	_, err := CopyObjectProperties(oldOb, newOb)
	return to, err
}

// CopyObjectProperties updates the "old" object properties with the "new's"
// Including ID and Type
func CopyObjectProperties(to, from *Object) (*Object, error) {
	to.ID = from.ID
	to.Type = from.Type
	to.Name = replaceIfNaturalLanguageValues(to.Name, from.Name)
	to.Attachment = replaceIfItem(to.Attachment, from.Attachment)
	to.AttributedTo = replaceIfItem(to.AttributedTo, from.AttributedTo)
	to.Audience = replaceIfItemCollection(to.Audience, from.Audience)
	to.Content = replaceIfNaturalLanguageValues(to.Content, from.Content)
	to.Context = replaceIfItem(to.Context, from.Context)
	if len(from.MediaType) > 0 {
		to.MediaType = from.MediaType
	}
	if !from.EndTime.IsZero() {
		to.EndTime = from.EndTime
	}
	to.Generator = replaceIfItem(to.Generator, from.Generator)
	to.Icon = replaceIfItem(to.Icon, from.Icon)
	to.Image = replaceIfItem(to.Image, from.Image)
	to.InReplyTo = replaceIfItem(to.InReplyTo, from.InReplyTo)
	to.Location = replaceIfItem(to.Location, from.Location)
	to.Preview = replaceIfItem(to.Preview, from.Preview)
	if to.Published.IsZero() && !from.Published.IsZero() {
		to.Published = from.Published
	}
	if to.Updated.IsZero() && !from.Updated.IsZero() {
		to.Updated = from.Updated
	}
	to.Replies = replaceIfItem(to.Replies, from.Replies)
	if !from.StartTime.IsZero() {
		to.StartTime = from.StartTime
	}
	to.Summary = replaceIfNaturalLanguageValues(to.Summary, from.Summary)
	to.Tag = replaceIfItemCollection(to.Tag, from.Tag)
	if from.URL != nil {
		to.URL = from.URL
	}
	to.To = replaceIfItemCollection(to.To, from.To)
	to.Bto = replaceIfItemCollection(to.Bto, from.Bto)
	to.CC = replaceIfItemCollection(to.CC, from.CC)
	to.BCC = replaceIfItemCollection(to.BCC, from.BCC)
	if from.Duration == 0 {
		to.Duration = from.Duration
	}
	to.Source = replaceIfSource(to.Source, from.Source)
	return to, nil
}

func copyAllItemProperties(to, from Item) (Item, error) {
	if CollectionType == to.GetType() {
		o, err := ToCollection(to)
		if err != nil {
			return o, err
		}
		n, err := ToCollection(from)
		if err != nil {
			return o, err
		}
		return CopyCollectionProperties(o, n)
	}
	if CollectionPageType == to.GetType() {
		o, err := ToCollectionPage(to)
		if err != nil {
			return o, err
		}
		n, err := ToCollectionPage(from)
		if err != nil {
			return o, err
		}
		return CopyCollectionPageProperties(o, n)
	}
	if OrderedCollectionType == to.GetType() {
		o, err := ToOrderedCollection(to)
		if err != nil {
			return o, err
		}
		n, err := ToOrderedCollection(from)
		if err != nil {
			return o, err
		}
		return CopyOrderedCollectionProperties(o, n)
	}
	if OrderedCollectionPageType == to.GetType() {
		o, err := ToOrderedCollectionPage(to)
		if err != nil {
			return o, err
		}
		n, err := ToOrderedCollectionPage(from)
		if err != nil {
			return o, err
		}
		return CopyOrderedCollectionPageProperties(o, n)
	}
	if ActorTypes.Contains(to.GetType()) {
		o, err := ToActor(to)
		if err != nil {
			return o, err
		}
		n, err := ToActor(from)
		if err != nil {
			return o, err
		}
		return UpdatePersonProperties(o, n)
	}
	if ObjectTypes.Contains(to.GetType()) || to.GetType() == "" {
		o, err := ToObject(to)
		if err != nil {
			return o, err
		}
		n, err := ToObject(from)
		if err != nil {
			return o, err
		}
		return CopyObjectProperties(o, n)
	}
	return to, fmt.Errorf("could not process objects with type %s", to.GetType())
}

// CopyItemProperties delegates to the correct per type functions for copying
// properties between matching Activity Objects
func CopyItemProperties(to, from Item) (Item, error) {
	if to == nil {
		return to, fmt.Errorf("nil object to update")
	}
	if from == nil {
		return to, fmt.Errorf("nil object for update")
	}
	if !to.GetLink().Equals(from.GetLink(), false) {
		return to, fmt.Errorf("object IDs don't match")
	}
	return copyAllItemProperties(to, from)
}

func CopyUnsafeItemProperties(to, from Item) (Item, error) {
	if to == nil {
		return to, fmt.Errorf("nil object to update")
	}
	if from == nil {
		return to, fmt.Errorf("nil object for update")
	}
	return copyAllItemProperties(to, from)
}

// UpdatePersonProperties
func UpdatePersonProperties(to, from *Actor) (*Actor, error) {
	oldOb, _ := ToObject(to)
	newOb, _ := ToObject(from)
	_, err := CopyObjectProperties(oldOb, newOb)
	if err != nil {
		return to, err
	}
	to.Inbox = replaceIfItem(to.Inbox, from.Inbox)
	to.Outbox = replaceIfItem(to.Outbox, from.Outbox)
	to.Following = replaceIfItem(to.Following, from.Following)
	to.Followers = replaceIfItem(to.Followers, from.Followers)
	to.Liked = replaceIfItem(to.Liked, from.Liked)
	to.PreferredUsername = replaceIfNaturalLanguageValues(to.PreferredUsername, from.PreferredUsername)
	to.PublicKey = replaceIfPublicKey(to.PublicKey, from.PublicKey)
	return to, nil
}

func replaceIfItem(old, new Item) Item {
	if new == nil {
		return old
	}
	return new
}

func replaceIfItemCollection(old, new ItemCollection) ItemCollection {
	if new == nil {
		return old
	}
	return new
}

func replaceIfNaturalLanguageValues(old, new NaturalLanguageValues) NaturalLanguageValues {
	if new == nil {
		return old
	}
	return new
}

func replaceIfSource(to, from Source) Source {
	if from.MediaType != to.MediaType {
		return from
	}
	to.Content = replaceIfNaturalLanguageValues(to.Content, from.Content)
	return to
}

func replaceIfPublicKey(to, from PublicKey) PublicKey {
	if from.ID != to.ID {
		return from
	}
	to.Owner = from.Owner
	to.PublicKeyPem = from.PublicKeyPem
	return to
}
