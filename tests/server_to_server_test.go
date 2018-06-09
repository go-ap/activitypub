package tests

// Server to Server tests from: https://test.activitypub.rocks/

import (
	"activitypub"
	"fmt"
	"testing"
)

// S2S Server: Activities requiring the object property
// The distribution of the following activities require that they contain the object property:
// Create, Update, Delete, Follow, Add, Remove, Like, Block, Undo.
// Implementation always includes object property for each of the above supported activities
func TestObjectPropertyExists(t *testing.T) {
	desc := `
S2S Server: Activities requiring the object property
  The distribution of the following activities require that they contain the object property:
   Create, Update, Delete, Follow, Add, Remove, Like, Block, Undo.

  Implementation always includes object property for each of the above supported activities
`
	t.Log(desc)

	obj := activitypub.MentionNew("gigel")

	add := activitypub.AddNew("https://localhost/myactivity", obj, nil)
	if add.Object == nil {
		t.Errorf("Missing GetID in Add activity %#v", add.Object)
	}
	if add.Object != obj {
		t.Errorf("Add.GetID different than what we initialized %#v %#v", add.Object, obj)
	}

	block := activitypub.BlockNew("https://localhost/myactivity", obj)
	if block.Object == nil {
		t.Errorf("Missing GetID in Add activity %#v", block.Object)
	}
	if block.Object != obj {
		t.Errorf("Block.GetID different than what we initialized %#v %#v", block.Object, obj)
	}

	create := activitypub.CreateNew("https://localhost/myactivity", obj)
	if create.Object == nil {
		t.Errorf("Missing GetID in Add activity %#v", create.Object)
	}
	if create.Object != obj {
		t.Errorf("Create.GetID different than what we initialized %#v %#v", create.Object, obj)
	}

	delete := activitypub.DeleteNew("https://localhost/myactivity", obj)
	if delete.Object == nil {
		t.Errorf("Missing GetID in Delete activity %#v", delete.Object)
	}
	if delete.Object != obj {
		t.Errorf("Delete.GetID different than what we initialized %#v %#v", delete.Object, obj)
	}

	follow := activitypub.FollowNew("https://localhost/myactivity", obj)
	if follow.Object == nil {
		t.Errorf("Missing GetID in Follow activity %#v", follow.Object)
	}
	if follow.Object != obj {
		t.Errorf("Follow.GetID different than what we initialized %#v %#v", follow.Object, obj)
	}

	like := activitypub.LikeNew("https://localhost/myactivity", obj)
	if like.Object == nil {
		t.Errorf("Missing GetID in Like activity %#v", like.Object)
	}
	if like.Object != obj {
		t.Errorf("Like.GetID different than what we initialized %#v %#v", add.Object, obj)
	}

	update := activitypub.UpdateNew("https://localhost/myactivity", obj)
	if update.Object == nil {
		t.Errorf("Missing GetID in Update activity %#v", update.Object)
	}
	if update.Object != obj {
		t.Errorf("Update.GetID different than what we initialized %#v %#v", update.Object, obj)
	}

	undo := activitypub.UndoNew("https://localhost/myactivity", obj)
	if undo.Object == nil {
		t.Errorf("Missing GetID in Undo activity %#v", undo.Object)
	}
	if undo.Object != obj {
		t.Errorf("Undo.GetID different than what we initialized %#v %#v", undo.Object, obj)
	}
}

// S2S Server: Activities requiring the target property
// The distribution of the following activities require that they contain the target property:
//   Add, Remove.
// Implementation always includes target property for each of the above supported activities.
func TestTargetPropertyExists(t *testing.T) {
	desc := `
S2S Server: Activities requiring the target property
  The distribution of the following activities require that they contain the target
property: Add, Remove.

  Implementation always includes target property for each of the above supported activities.
`
	t.Log(desc)

	obj := activitypub.MentionNew("foo")
	target := activitypub.MentionNew("bar")

	add := activitypub.AddNew("https://localhost/myactivity", obj, target)
	if add.Target == nil {
		t.Errorf("Missing Target in Add activity %#v", add.Target)
	}
	if add.Target != target {
		t.Errorf("Add.Target different than what we initialized %#v %#v", add.Target, target)
	}

	remove := activitypub.RemoveNew("https://localhost/myactivity", obj, target)
	if remove.Target == nil {
		t.Errorf("Missing Target in Remove activity %#v", remove.Target)
	}
	if remove.Target != target {
		t.Errorf("Remove.Target different than what we initialized %#v %#v", remove.Target, target)
	}
}

// S2S Server: Deduplication of recipient list
// Attempt to submit for delivery an activity that addresses the same actor
// (ie an actor with the same id) twice.
// (For example, the same actor could appear on both the to and cc fields,
//  or the actor could be explicitly addressed
// in to but could also be a member of the addressed followers collection of the sending actor.)
// The server should deduplicate the list of inboxes to deliver to before delivering.
// The final recipient list is deduplicated before delivery.
func TestDeduplication(t *testing.T) {
	desc := `
S2S Server: Deduplication of recipient list
  Attempt to submit for delivery an activity that addresses the same actor
(ie an actor with the same id) twice.

  The final recipient list is deduplicated before delivery.
`
	t.Log(desc)

	to := activitypub.PersonNew("bob")
	o := activitypub.ObjectNew("something", activitypub.ArticleType)
	cc := activitypub.PersonNew("alice")

	c := activitypub.CreateNew("create", o)
	c.To.Append(to)
	c.CC.Append(cc)
	c.BCC.Append(cc)

	c.RecipientsDeduplication()

	checkDedup := func(list activitypub.ObjectsArr, recIds *[]activitypub.ObjectID) error {
		for _, rec := range list {
			for _, id := range *recIds {
				if rec.GetID() == id {
					return fmt.Errorf("%T[%s] already stored in recipients list, Deduplication faild", rec, id)
				}
			}
			*recIds = append(*recIds, rec.GetID())
		}
		return nil
	}

	var err error
	recIds := make([]activitypub.ObjectID, 0)
	err = checkDedup(c.To, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(c.Bto, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(c.CC, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = checkDedup(c.BCC, &recIds)
	if err != nil {
		t.Error(err)
	}
}

// S2S Server: Do-not-deliver considerations
// Server does not deliver to recipients which are the same as the actor of the
// Activity being notified about
func TestDoNotDeliverToActor(t *testing.T) {
	desc := `
S2S Server: Do-not-deliver considerations

  Server does not deliver to recipients which are the same as the actor of the
Activity being notified about
`
	t.Log(desc)

	p := activitypub.PersonNew("main actor")

	to := activitypub.PersonNew("bob")
	o := activitypub.ObjectNew("something", activitypub.ArticleType)
	cc := activitypub.PersonNew("alice")

	c := activitypub.CreateNew("create", o)
	c.Actor = activitypub.Actor(*p)

	c.To.Append(p)
	c.To.Append(to)
	c.CC.Append(cc)
	c.CC.Append(p)
	c.BCC.Append(cc)
	c.BCC.Append(p)

	c.RecipientsDeduplication()

	checkActor := func(list activitypub.ObjectsArr, actor activitypub.Actor) error {
		for _, rec := range list {
			if rec.GetID() == actor.GetID() {
				return fmt.Errorf("%T[%s] Actor of activity should not be in the recipients list", rec, actor.GetID())
			}
		}
		return nil
	}

	var err error
	err = checkActor(c.To, c.Actor)
	if err != nil {
		t.Error(err)
	}
	err = checkActor(c.Bto, c.Actor)
	if err != nil {
		t.Error(err)
	}
	err = checkActor(c.CC, c.Actor)
	if err != nil {
		t.Error(err)
	}
	err = checkActor(c.BCC, c.Actor)
	if err != nil {
		t.Error(err)
	}
}

// S2S Server: Do-not-deliver considerations
// Server does not deliver Block activities to their object.
func TestDoNotDeliverBlockToObject(t *testing.T) {
	desc := `
S2S Server: Do-not-deliver considerations

  Server does not deliver Block activities to their object.
`
	t.Log(desc)

	p := activitypub.PersonNew("blocked")

	bob := activitypub.PersonNew("bob")
	jane := activitypub.PersonNew("jane doe")

	b := activitypub.BlockNew("block actor", p)
	b.Actor = activitypub.Actor(*bob)

	b.To.Append(jane)
	b.To.Append(p)
	b.To.Append(bob)

	b.RecipientsDeduplication()

	checkActor := func(list activitypub.ObjectsArr, ob activitypub.ObjectOrLink) error {
		for _, rec := range list {
			if rec.GetID() == ob.GetID() {
				return fmt.Errorf("%T[%s] of activity should not be in the recipients list", rec, ob.GetID())
			}
		}
		return nil
	}

	var err error
	err = checkActor(b.To, b.Object)
	if err != nil {
		t.Error(err)
	}
	err = checkActor(b.To, b.Actor)
	if err != nil {
		t.Error(err)
	}
}

// S2S Sever: Support for sharedInbox
// Delivers to sharedInbox endpoints to reduce the number of receiving actors delivered
// to by identifying all followers which share the same sharedInbox who would otherwise be
// individual recipients and instead deliver objects to said sharedInbox.
func TestSharedInboxIdentifySharedInbox(t *testing.T) {
	desc := `
S2S Sever: Support for sharedInbox

  Delivers to sharedInbox endpoints to reduce the number of receiving actors delivered 
to by identifying all followers which share the same sharedInbox who would otherwise be
individual recipients and instead deliver objects to said sharedInbox.
`
	t.Skip(desc)
}

// S2S Sever: Support for sharedInbox
//  Deliver to actor inboxes and collections otherwise addressed which do not have a sharedInbox.
func TestSharedInboxActorsWOSharedInbox(t *testing.T) {
	desc := `
S2S Server: Do-not-deliver considerations

  Server does not deliver Block activities to their object.
`
	t.Skip(desc)
}

// S2S Server: Deduplicating received activities
// Server deduplicates activities received in inbox by comparing activity ids
func TestInboxDeduplication(t *testing.T) {
	desc := `
S2S Server: Deduplicating received activities

  Server deduplicates activities received in inbox by comparing activity ids
`
	t.Skip(desc)
}

// S2S Server: Special forwarding mechanism
// ActivityPub contains a special mechanism for forwarding replies to avoid "ghost replies".
//  Forwards incoming activities to the values of to, bto, cc, bcc, audience if and only if criteria are met.
func TestForwardingMechanismsToRecipients(t *testing.T) {
	desc := `
S2S Server: Special forwarding mechanism
 ActivityPub contains a special mechanism for forwarding replies to avoid "ghost replies".

  Forwards incoming activities to the values of to, bto, cc, bcc, audience if and only if criteria are met.
`
	t.Skip(desc)
}

// S2S Server: Special forwarding mechanism
// ActivityPub contains a special mechanism for forwarding replies to avoid "ghost replies".
//  Recurse through to, bto, cc, bcc, audience object values to determine whether/where
// to forward according to criteria in 7.1.2
func TestForwardingMechanismsRecurseRecipients(t *testing.T) {
	desc := `
S2S Server: Special forwarding mechanism
ActivityPub contains a special mechanism for forwarding replies to avoid "ghost replies".

  Recurse through to, bto, cc, bcc, audience object values to determine whether/where
  to forward according to criteria in 7.1.2
`
	t.Skip(desc)
}

// S2S Server: Special forwarding mechanism
// ActivityPub contains a special mechanism for forwarding replies to avoid "ghost replies".
//  Limits depth of this recursion.
func TestForwardingMechanismsLimitsRecursion(t *testing.T) {
	desc := `
S2S Server: Special forwarding mechanism
 ActivityPub contains a special mechanism for forwarding replies to avoid "ghost replies".

  Limits depth of this recursion.
`
	t.Skip(desc)
}

// S2S Server: Verification of content authorship
// Before accepting activities delivered to an actor's inbox some sort of verification
// should be performed. (For example, if the delivering actor has a public key on their profile,
// the request delivering the activity may be signed with HTTP Signatures.)
// Don't trust content received from a server other than the content's origin without some form of verification.
func TestVerification(t *testing.T) {
	desc := `
S2S Server: Verification of content authorship
 Before accepting activities delivered to an actor's inbox some sort of verification
 should be performed. (For example, if the delivering actor has a public key on their profile,
 the request delivering the activity may be signed with HTTP Signatures.)

  Don't trust content received from a server other than the content's origin without some form of verification.
`
	t.Skip(desc)
}

// S2S Server: Update activity
// On receiving an Update activity to an actor's inbox, the server:
//  Takes care to be sure that the Update is authorized to modify its object
func TestUpdateIsAuthorized(t *testing.T) {
	desc := `
S2S Server: Update activity
 On receiving an Update activity to an actor's inbox, the server:

  Takes care to be sure that the Update is authorized to modify its object
`
	t.Skip(desc)
}

// S2S Server: Update activity
// On receiving an Update activity to an actor's inbox, the server:
//  Completely replaces its copy of the activity with the newly received value
func TestUpdateReplacesActivity(t *testing.T) {
	desc := `
S2S Server: Update activity
 On receiving an Update activity to an actor's inbox, the server:

  Completely replaces its copy of the activity with the newly received value
`
	t.Skip(desc)
}

// S2S Server: Delete activity
//  Delete removes object's representation, assuming object is owned by sending actor/server
func TestDeleteRemoves(t *testing.T) {
	desc := `
S2S Server: Delete activity

  Delete removes object's representation, assuming object is owned by sending actor/server
`
	t.Skip(desc)
}

// S2S Server: Delete activity
//  Replaces deleted object with a Tombstone object (optional)
func TestDeleteReplacesWithTombstone(t *testing.T) {
	desc := `
S2S Server: Delete activity

  Replaces deleted object with a Tombstone object (optional)
`
	t.Skip(desc)
}

//S2S Server: Following, and handling accept/reject of follows
//  Follow should add the activity's actor to the receiving actor's Followers Collection.
func TestFollowAddsToFollowers(t *testing.T) {
	desc := `
S2S Server: Following, and handling accept/reject of follows

  Follow should add the activity's actor to the receiving actor's Followers Collection.
`
	t.Skip(desc)
}

//S2S Server: Following, and handling accept/reject of follows
//  Generates either an Accept or Reject activity with Follow as object and deliver to actor of the Follow
func TestGeneratesAcceptOrReject(t *testing.T) {
	desc := `
S2S Server: Following, and handling accept/reject of follows

  Generates either an Accept or Reject activity with Follow as object and deliver to actor of the Follow
`
	t.Skip(desc)
}

//S2S Server: Following, and handling accept/reject of follows
//  If receiving an Accept in reply to a Follow activity, adds actor to receiver's Following Collection
func TestAddsFollowerIfAccept(t *testing.T) {
	desc := `
S2S Server: Following, and handling accept/reject of follows

  If receiving an Accept in reply to a Follow activity, adds actor to receiver's Following Collection
`
	t.Skip(desc)
}

//S2S Server: Following, and handling accept/reject of follows
//  If receiving a Reject in reply to a Follow activity, does not add actor to receiver's Following Collection
func TestDoesntAddFollowerIfReject(t *testing.T) {
	desc := `
S2S Server: Following, and handling accept/reject of follows

  If receiving a Reject in reply to a Follow activity, does not add actor to receiver's Following Collection
`
	t.Skip(desc)
}

//S2S Server: Activity acceptance side-effects
// Test accepting the following activities to an actor's inbox and observe the side effects:
//
//  Create makes record of the object existing
func TestCreateMakesRecord(t *testing.T) {
	desc := `
S2S Server: Activity acceptance side-effects
 Test accepting the following activities to an actor's inbox and observe the side effects:

  Create makes record of the object existing
`
	t.Skip(desc)
}

//S2S Server: Activity acceptance side-effects
// Test accepting the following activities to an actor's inbox and observe the side effects:
//
//  Add should add the activity's object to the Collection specified in the target property,
//  unless not allowed per requirements
func TestAddObjectToTarget(t *testing.T) {
	desc := `
S2S Server: Activity acceptance side-effects
 Test accepting the following activities to an actor's inbox and observe the side effects:

  Add should add the activity's object to the Collection specified in the target property,
  unless not allowed per requirements
`
	t.Skip(desc)
}

//S2S Server: Activity acceptance side-effects
// Test accepting the following activities to an actor's inbox and observe the side effects:
//
//  Remove should remove the object from the Collection specified in the target property,
//  unless not allowed per requirements
func TestRemoveObjectFromTarget(t *testing.T) {
	desc := `
S2S Server: Activity acceptance side-effects
 Test accepting the following activities to an actor's inbox and observe the side effects:

  Remove should remove the object from the Collection specified in the target property,
  unless not allowed per requirements
`
	t.Skip(desc)
}

//S2S Server: Activity acceptance side-effects
// Test accepting the following activities to an actor's inbox and observe the side effects:
//
//  Like increments the object's count of likes by adding the received activity to the likes
//  collection if this collection is present
func TestLikeIncrementsLikes(t *testing.T) {
	desc := `
S2S Server: Activity acceptance side-effects
 Test accepting the following activities to an actor's inbox and observe the side effects:

  Like increments the object's count of likes by adding the received activity to the likes
  collection if this collection is present
`
	t.Skip(desc)
}

//S2S Server: Activity acceptance side-effects
// Test accepting the following activities to an actor's inbox and observe the side effects:
//
//  Announce increments object's count of shares by adding the received activity to the
// 'shares' collection if this collection is present
func TestAnnounceIncrementsShares(t *testing.T) {
	desc := `
S2S Server: Activity acceptance side-effects
 Test accepting the following activities to an actor's inbox and observe the side effects:

  Announce increments object's count of shares by adding the received activity to the
 'shares' collection if this collection is present
`
	t.Skip(desc)
}

//S2S Server: Activity acceptance side-effects
// Test accepting the following activities to an actor's inbox and observe the side effects:
//
//  Undo performs Undo of object in federated context
func TestUndoPerformsUndoOnObject(t *testing.T) {
	desc := `
S2S Server: Activity acceptance side-effects
 Test accepting the following activities to an actor's inbox and observe the side effects:

  Undo performs Undo of object in federated context
`
	t.Skip(desc)
}
