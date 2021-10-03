package activitypub

import (
	"fmt"
)

func CopyOrderedCollectionPageProperties(old, new *OrderedCollectionPage) (*OrderedCollectionPage, error) {
	old.PartOf = replaceIfItem(old.PartOf, new.PartOf)
	old.Next = replaceIfItem(old.Next, new.Next)
	old.Prev = replaceIfItem(old.Prev, new.Prev)
	oldCol, _ := ToOrderedCollection(old)
	newCol, _ := ToOrderedCollection(new)
	_, err := CopyOrderedCollectionProperties(oldCol, newCol)
	if err != nil {
		return old, err
	}
	return old, nil
}

func CopyCollectionPageProperties(old, new *CollectionPage) (*CollectionPage, error) {
	old.PartOf = replaceIfItem(old.PartOf, new.PartOf)
	old.Next = replaceIfItem(old.Next, new.Next)
	old.Prev = replaceIfItem(old.Prev, new.Prev)
	oldCol, _ := ToCollection(old)
	newCol, _ := ToCollection(new)
	_, err := CopyCollectionProperties(oldCol, newCol)
	return old, err
}

func CopyOrderedCollectionProperties(old, new *OrderedCollection) (*OrderedCollection, error) {
	old.First = replaceIfItem(old.First, new.First)
	old.Last = replaceIfItem(old.Last, new.Last)
	old.OrderedItems = replaceIfItemCollection(old.OrderedItems, new.OrderedItems)
	if old.TotalItems == 0 {
		old.TotalItems = new.TotalItems
	}
	oldOb, _ := ToObject(old)
	newOb, _ := ToObject(new)
	_, err := CopyObjectProperties(oldOb, newOb)
	return old, err
}

func CopyCollectionProperties(old, new *Collection) (*Collection, error) {
	old.First = replaceIfItem(old.First, new.First)
	old.Last = replaceIfItem(old.Last, new.Last)
	old.Items = replaceIfItemCollection(old.Items, new.Items)
	if old.TotalItems == 0 {
		old.TotalItems = new.TotalItems
	}
	oldOb, _ := ToObject(old)
	newOb, _ := ToObject(new)
	_, err := CopyObjectProperties(oldOb, newOb)
	return old, err
}

// CopyObjectProperties updates the "old" object properties with "new's"
func CopyObjectProperties(old, new *Object) (*Object, error) {
	old.Name = replaceIfNaturalLanguageValues(old.Name, new.Name)
	old.Attachment = replaceIfItem(old.Attachment, new.Attachment)
	old.AttributedTo = replaceIfItem(old.AttributedTo, new.AttributedTo)
	old.Audience = replaceIfItemCollection(old.Audience, new.Audience)
	old.Content = replaceIfNaturalLanguageValues(old.Content, new.Content)
	old.Context = replaceIfItem(old.Context, new.Context)
	if len(new.MediaType) > 0 {
		old.MediaType = new.MediaType
	}
	if !new.EndTime.IsZero() {
		old.EndTime = new.EndTime
	}
	old.Generator = replaceIfItem(old.Generator, new.Generator)
	old.Icon = replaceIfItem(old.Icon, new.Icon)
	old.Image = replaceIfItem(old.Image, new.Image)
	old.InReplyTo = replaceIfItem(old.InReplyTo, new.InReplyTo)
	old.Location = replaceIfItem(old.Location, new.Location)
	old.Preview = replaceIfItem(old.Preview, new.Preview)
	if old.Published.IsZero() && !new.Published.IsZero() {
		old.Published = new.Published
	}
	old.Replies = replaceIfItem(old.Replies, new.Replies)
	if !new.StartTime.IsZero() {
		old.StartTime = new.StartTime
	}
	old.Summary = replaceIfNaturalLanguageValues(old.Summary, new.Summary)
	old.Tag = replaceIfItemCollection(old.Tag, new.Tag)
	if !new.Updated.IsZero() {
		old.Updated = new.Updated
	}
	if new.URL != nil {
		old.URL = new.URL
	}
	old.To = replaceIfItemCollection(old.To, new.To)
	old.Bto = replaceIfItemCollection(old.Bto, new.Bto)
	old.CC = replaceIfItemCollection(old.CC, new.CC)
	old.BCC = replaceIfItemCollection(old.BCC, new.BCC)
	if new.Duration == 0 {
		old.Duration = new.Duration
	}
	old.Source = replaceIfSource(old.Source, new.Source)
	return old, nil
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
	if to.GetType() != from.GetType() {
		return to, fmt.Errorf("invalid object types for update %s(old) and %s(new)", from.GetType(), to.GetType())
	}
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
	if ObjectTypes.Contains(to.GetType()) {
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

// UpdatePersonProperties
func UpdatePersonProperties(old, new *Actor) (*Actor, error) {
	old.Inbox = replaceIfItem(old.Inbox, new.Inbox)
	old.Outbox = replaceIfItem(old.Outbox, new.Outbox)
	old.Following = replaceIfItem(old.Following, new.Following)
	old.Followers = replaceIfItem(old.Followers, new.Followers)
	old.Liked = replaceIfItem(old.Liked, new.Liked)
	old.PreferredUsername = replaceIfNaturalLanguageValues(old.PreferredUsername, new.PreferredUsername)
	oldOb, _ := ToObject(old)
	newOb, _ := ToObject(new)
	_, err := CopyObjectProperties(oldOb, newOb)
	return old, err
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

func replaceIfSource(old, new Source) Source {
	if new.MediaType != old.MediaType {
		return new
	}
	old.Content = replaceIfNaturalLanguageValues(old.Content, new.Content)
	return old
}
