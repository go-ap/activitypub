package activitypub

/*
func GobEncode(it Item) ([]byte, error) {
	b := new(bytes.Buffer)
	err := gob.NewEncoder(b).Encode(it)
	return b.Bytes(), err
}

type hasType struct {
	Type ActivityVocabularyType
}

func GobUnmarshalToItem(data []byte) Item {
	if len(data) == 0 {
		return nil
	}
	if ItemTyperFunc == nil {
		return nil
	}
	typer := new(hasType)
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(typer)
	if err != nil {
		return nil
	}

	var it Item
	switch typer.Type {

	}
	return it
}

// UnmarshalGob
func UnmarshalGob(data []byte) (Item, error) {
	if ItemTyperFunc == nil {
		ItemTyperFunc = GetItemByType
	}
	return GobUnmarshalToItem(data), nil
}
*/
