package activitypub

// ItemCollection is an array of items
type ItemCollection []Item

// Item struct
type Item ObjectOrLink

// GetID returns the ObjectID corresponding to ItemCollection
func (i ItemCollection) GetID() *ObjectID {
	return nil
}

// GetType returns the ItemCollection's type
func (i ItemCollection) GetType() ActivityVocabularyType {
	return ActivityVocabularyType("")
}

// IsLink returns false for an ItemCollection object
func (i ItemCollection) IsLink() bool {
	return false
}

// IsObject returns true for a ItemCollection object
func (i ItemCollection) IsObject() bool {
	return false
}
