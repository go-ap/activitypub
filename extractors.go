package activitypub

func PreferredNameOf(it Item) string {
	var cont string
	if IsObject(it) {
		_ = OnActor(it, func(act *Actor) error {
			if act.PreferredUsername != nil {
				cont = act.PreferredUsername.First().String()
			}
			return nil
		})
	}
	return cont
}

func ContentOf(it Item) string {
	var cont string
	if IsObject(it) {
		_ = OnObject(it, func(ob *Object) error {
			if ob.Content != nil {
				cont = ob.Content.First().String()
			}
			return nil
		})
	}
	return cont
}

func SummaryOf(it Item) string {
	var cont string
	if IsObject(it) {
		_ = OnObject(it, func(ob *Object) error {
			if ob.Summary != nil {
				cont = ob.Summary.First().String()
			}
			return nil
		})
	}
	return cont
}

func NameOf(it Item) string {
	var name string
	if IsLink(it) {
		_ = OnLink(it, func(lnk *Link) error {
			if lnk.Name != nil {
				name = lnk.Name.First().String()
			}
			return nil
		})
	} else {
		_ = OnObject(it, func(ob *Object) error {
			if ob.Name != nil {
				name = ob.Name.First().String()
			}
			return nil
		})
	}
	return name
}
