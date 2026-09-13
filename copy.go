package activitypub

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
