package activitypub

import (
	as "github.com/go-ap/activitystreams"
)

// Activity is a subtype of Object that describes some form of action that may happen,
// is currently happening, or has already happened.
// The Activity type itself serves as an abstract base type for all types of activities.
// It is important to note that the Activity type itself does not carry any specific semantics
// about the kind of action being taken.
type Activity struct {
	as.Activity
}

// IntransitiveActivity Instances of IntransitiveActivity are a subtype of Activity representing intransitive actions.
// The object property is therefore inappropriate for these activities.
type IntransitiveActivity struct {
	as.IntransitiveActivity
}

type (
	// Accept indicates that the actor accepts the object. The target property can be used in certain circumstances to indicate
	// the context into which the object has been accepted.
	Accept = Activity

	// Add indicates that the actor has added the object to the target. If the target property is not explicitly specified,
	// the target would need to be determined implicitly by context.
	// The origin can be used to identify the context from which the object originated.
	Add = Activity

	// Announce indicates that the actor is calling the target's attention the object.
	// The origin typically has no defined meaning.
	Announce = Activity

	// Arrive is an IntransitiveActivity that indicates that the actor has arrived at the location.
	// The origin can be used to identify the context from which the actor originated.
	// The target typically has no defined meaning.
	Arrive = IntransitiveActivity

	// Block indicates that the actor is blocking the object. Blocking is a stronger form of Ignore.
	// The typical use is to support social systems that allow one user to block activities or content of other users.
	// The target and origin typically have no defined meaning.
	Block = Ignore

	// Create indicates that the actor has created the object.
	Create = Activity

	// Delete indicates that the actor has deleted the object.
	// If specified, the origin indicates the context from which the object was deleted.
	Delete = Activity

	// Dislike indicates that the actor dislikes the object.
	Dislike = Activity

	// Flag indicates that the actor is "flagging" the object.
	// Flagging is defined in the sense common to many social platforms as reporting content as being
	// inappropriate for any number of reasons.
	Flag = Activity

	// Follow indicates that the actor is "following" the object. Following is defined in the sense typically used within
	// Social systems in which the actor is interested in any activity performed by or on the object.
	// The target and origin typically have no defined meaning.
	Follow = Activity

	// Ignore indicates that the actor is ignoring the object. The target and origin typically have no defined meaning.
	Ignore = Activity

	// Invite is a specialization of Offer in which the actor is extending an invitation for the object to the target.
	Invite = Offer

	// Join indicates that the actor has joined the object. The target and origin typically have no defined meaning.
	Join = Activity

	// Leave indicates that the actor has left the object. The target and origin typically have no meaning.
	Leave = Activity

	// Like indicates that the actor likes, recommends or endorses the object.
	// The target and origin typically have no defined meaning.
	Like = Activity

	// Listen inherits all properties from Activity.
	Listen = Activity

	// Move indicates that the actor has moved object from origin to target.
	// If the origin or target are not specified, either can be determined by context.
	Move = Activity

	// Offer indicates that the actor is offering the object.
	// If specified, the target indicates the entity to which the object is being offered.
	Offer = Activity

	// Reject indicates that the actor is rejecting the object. The target and origin typically have no defined meaning.
	Reject = Activity

	// Read indicates that the actor has read the object.
	Read = Activity

	// Remove indicates that the actor is removing the object. If specified,
	// the origin indicates the context from which the object is being removed.
	Remove = Activity

	// TentativeReject is a specialization of Reject in which the rejection is considered tentative.
	TentativeReject = Reject

	// TentativeAccept is a specialization of Accept indicating that the acceptance is tentative.
	TentativeAccept = Accept

	// Travel indicates that the actor is traveling to target from origin.
	// Travel is an IntransitiveObject whose actor specifies the direct object.
	// If the target or origin are not specified, either can be determined by context.
	Travel = IntransitiveActivity

	// Undo indicates that the actor is undoing the object. In most cases, the object will be an Activity describing
	// some previously performed action (for instance, a person may have previously "liked" an article but,
	// for whatever reason, might choose to undo that like at some later point in time).
	// The target and origin typically have no defined meaning.
	Undo = Activity

	// Update indicates that the actor has updated the object. Note, however, that this vocabulary does not define a mechanism
	// for describing the actual set of modifications made to object.
	// The target and origin typically have no defined meaning.
	Update = Activity

	// View indicates that the actor has viewed the object.
	View = Activity
)

// Question represents a question being asked. Question objects are an extension of IntransitiveActivity.
// That is, the Question object is an Activity, but the direct object is the question
// itself and therefore it would not contain an object property.
// Either of the anyOf and oneOf properties may be used to express possible answers,
// but a Question object must not have both properties.
type Question = as.Question

// AcceptNew initializes an Accept activity
func AcceptNew(id as.ObjectID, ob as.Item) *Accept {
	return ActivityNew(id, as.AcceptType, ob)
}

// AddNew initializes an Add activity
func AddNew(id as.ObjectID, ob as.Item, trgt as.Item) *Add {
	a := ActivityNew(id, as.AddType, ob)
	a.Target = trgt
	return a
}

// AnnounceNew initializes an Announce activity
func AnnounceNew(id as.ObjectID, ob as.Item) *Announce {
	return ActivityNew(id, as.AnnounceType, ob)
}

// ArriveNew initializes an Arrive activity
func ArriveNew(id as.ObjectID) *Arrive {
	return IntransitiveActivityNew(id, as.ArriveType)
}

// BlockNew initializes a Block activity
func BlockNew(id as.ObjectID, ob as.Item) *Block {
	return ActivityNew(id, as.BlockType, ob)
}

// CreateNew initializes a Create activity
func CreateNew(id as.ObjectID, ob as.Item) *Create {
	return ActivityNew(id, as.CreateType, ob)
}

// DeleteNew initializes a Delete activity
func DeleteNew(id as.ObjectID, ob as.Item) *Delete {
	return ActivityNew(id, as.DeleteType, ob)
}

// DislikeNew initializes a Dislike activity
func DislikeNew(id as.ObjectID, ob as.Item) *Dislike {
	return ActivityNew(id, as.DislikeType, ob)
}

// FlagNew initializes a Flag activity
func FlagNew(id as.ObjectID, ob as.Item) *Flag {
	return ActivityNew(id, as.FlagType, ob)
}

// FollowNew initializes a Follow activity
func FollowNew(id as.ObjectID, ob as.Item) *Follow {
	return ActivityNew(id, as.FollowType, ob)
}

// IgnoreNew initializes an Ignore activity
func IgnoreNew(id as.ObjectID, ob as.Item) *Ignore {
	return ActivityNew(id, as.IgnoreType, ob)
}

// InviteNew initializes an Invite activity
func InviteNew(id as.ObjectID, ob as.Item) *Invite {
	return ActivityNew(id, as.InviteType, ob)
}

// JoinNew initializes a Join activity
func JoinNew(id as.ObjectID, ob as.Item) *Join {
	return ActivityNew(id, as.JoinType, ob)
}

// LeaveNew initializes a Leave activity
func LeaveNew(id as.ObjectID, ob as.Item) *Leave {
	return ActivityNew(id, as.LeaveType, ob)
}

// LikeNew initializes a Like activity
func LikeNew(id as.ObjectID, ob as.Item) *Like {
	return ActivityNew(id, as.LikeType, ob)
}

// ListenNew initializes a Listen activity
func ListenNew(id as.ObjectID, ob as.Item) *Listen {
	return ActivityNew(id, as.ListenType, ob)
}

// MoveNew initializes a Move activity
func MoveNew(id as.ObjectID, ob as.Item) *Move {
	return ActivityNew(id, as.MoveType, ob)
}

// OfferNew initializes an Offer activity
func OfferNew(id as.ObjectID, ob as.Item) *Offer {
	return ActivityNew(id, as.OfferType, ob)
}

// RejectNew initializes a Reject activity
func RejectNew(id as.ObjectID, ob as.Item) *Reject {
	return ActivityNew(id, as.RejectType, ob)
}

// ReadNew initializes a Read activity
func ReadNew(id as.ObjectID, ob as.Item) *Read {
	return ActivityNew(id, as.ReadType, ob)
}

// RemoveNew initializes a Remove activity
func RemoveNew(id as.ObjectID, ob as.Item, trgt as.Item) *Remove {
	a := ActivityNew(id, as.RemoveType, ob)
	a.Target = trgt
	return a
}

// TentativeRejectNew initializes a TentativeReject activity
func TentativeRejectNew(id as.ObjectID, ob as.Item) *TentativeReject {
	a := ActivityNew(id, as.TentativeRejectType, ob)
	o := TentativeReject(*a)
	return &o
}

// TentativeAcceptNew initializes a TentativeAccept activity
func TentativeAcceptNew(id as.ObjectID, ob as.Item) *TentativeAccept {
	return ActivityNew(id, as.TentativeAcceptType, ob)
}

// TravelNew initializes a Travel activity
func TravelNew(id as.ObjectID) *Travel {
	return IntransitiveActivityNew(id, as.TravelType)
}

// UndoNew initializes an Undo activity
func UndoNew(id as.ObjectID, ob as.Item) *Undo {
	return ActivityNew(id, as.UndoType, ob)
}

// UpdateNew initializes an Update activity
func UpdateNew(id as.ObjectID, ob as.Item) *Update {
	return ActivityNew(id, as.UpdateType, ob)
}

// ViewNew initializes a View activity
func ViewNew(id as.ObjectID, ob as.Item) *View {
	return ActivityNew(id, as.ViewType, ob)
}

// QuestionNew initializes a Question activity
func QuestionNew(id as.ObjectID) *Question {
	q := Question{ ID: id, Type: as.QuestionType}
	q.Name = as.NaturalLanguageValueNew()
	q.Content = as.NaturalLanguageValueNew()
	return &q
}

// ActivityNew initializes a basic activity
func ActivityNew(id as.ObjectID, typ as.ActivityVocabularyType, ob as.Item) *Activity {
	a := as.ActivityNew(id, typ, ob)
	return &Activity{ Activity: *a}
}

// IntransitiveActivityNew initializes a intransitive activity
func IntransitiveActivityNew(id as.ObjectID, typ as.ActivityVocabularyType) *IntransitiveActivity {
	i := as.IntransitiveActivityNew(id, typ)
	return &IntransitiveActivity{ IntransitiveActivity: *i}
}

// UnmarshalJSON
func (a *Activity) UnmarshalJSON(data []byte) error {
	as.ItemTyperFunc = JSONGetItemByType

	a.Parent.UnmarshalJSON(data)
	a.Actor = as.JSONGetItem(data, "actor")
	a.Object = as.JSONGetItem(data, "object")
	return nil
}

func loadActorWithInboxObject(a as.Item, o as.Item) as.Item {
	typ := a.GetType()
	switch typ {
	case as.ApplicationType:
		var app Application
		app, _ = a.(Application)
		var inbox *as.OrderedCollection
		if app.Inbox == nil {
			inbox = InboxNew()
		} else {
			inbox = app.Inbox.(*as.OrderedCollection)
		}
		inbox.Append(o)
		app.Inbox = inbox
		return app
	case as.GroupType:
		var grp Group
		grp, _ = a.(Group)
		var inbox *as.OrderedCollection
		if grp.Inbox == nil {
			inbox = InboxNew()
		} else {
			inbox = grp.Inbox.(*as.OrderedCollection)
		}
		inbox.Append(o)
		grp.Inbox = inbox
		return grp
	case as.OrganizationType:
		var org Organization
		org, _ = a.(Organization)
		var inbox *as.OrderedCollection
		if org.Inbox == nil {
			inbox = InboxNew()
		} else {
			inbox = org.Inbox.(*as.OrderedCollection)
		}
		inbox.Append(o)
		org.Inbox = inbox
		return org
	case as.PersonType:
		var pers Person
		pers, _ = a.(Person)
		var inbox *as.OrderedCollection
		if pers.Inbox == nil {
			inbox = InboxNew()
		} else {
			inbox = pers.Inbox.(*as.OrderedCollection)
		}
		inbox.Append(o)
		pers.Inbox = inbox
		return pers
	case as.ServiceType:
		var serv Service
		serv, _ = a.(Service)
		var inbox *as.OrderedCollection
		if serv.Inbox == nil {
			inbox = InboxNew()
		} else {
			inbox = serv.Inbox.(*as.OrderedCollection)
		}
		inbox.Append(o)
		serv.Inbox = inbox
		return serv
	default:
		o, _ := a.(as.Object)
		return Object{Parent: o}
	}
}
