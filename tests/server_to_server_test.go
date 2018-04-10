package tests

// Server to Server tests from: https://test.activitypub.rocks/

import (
	"activitypub"
	"fmt"
	"os"
	//	"reflect"
	"testing"
)

/*
func testActivityHasObject(activity interface{}, obj LinkOrObject) error {
	activityType := reflect.TypeOf(activity)
	if activity.Object == nil {
		return fmt.Errorf("Missing Object in %s activity %#v", activityType, activity.Object)
	}
	if activity.Object != obj {
		return fmt.Errorf("%s.Object different than what we initialized %#v %#v", activityType, activity.Object, obj)
	}
}
*/

func TestMain(m *testing.M) {
	fmt.Println("starting")
	status := m.Run()
	fmt.Println("ending")
	os.Exit(status)
}

// S2S Server: Activities requiring the object property
// The distribution of the following activities require that they contain the object property:
// Create, Update, Delete, Follow, Add, Remove, Like, Block, Undo.
//	Implementation always includes object property for each of the above supported activities
func TestObjectPropertyExistsForAdd(t *testing.T) {
	obj := activitypub.MentionNew("gigel")
	add := activitypub.AddNew("https://localhost/myactivity", obj, nil)

	if add.Object == nil {
		t.Errorf("Missing Object in Add activity %#v", add.Object)
	}
	if add.Object != obj {
		t.Errorf("Add.Object different than what we initialized %#v %#v", add.Object, obj)
	}
}

func TestObjectPropertyExistsForBlock(t *testing.T) {
	obj := activitypub.MentionNew("gigel")
	block := activitypub.BlockNew("https://localhost/myactivity", obj)

	if block.Object == nil {
		t.Errorf("Missing Object in Add activity %#v", block.Object)
	}
	if block.Object != obj {
		t.Errorf("Block.Object different than what we initialized %#v %#v", block.Object, obj)
	}
}

func TestObjectPropertyExistsForCreate(t *testing.T) {
	obj := activitypub.MentionNew("gigel")
	create := activitypub.CreateNew("https://localhost/myactivity", obj)

	if create.Object == nil {
		t.Errorf("Missing Object in Add activity %#v", create.Object)
	}
	if create.Object != obj {
		t.Errorf("Create.Object different than what we initialized %#v %#v", create.Object, obj)
	}
}

func TestObjectPropertyExistsForDelete(t *testing.T) {
	obj := activitypub.MentionNew("gigel")
	delete := activitypub.DeleteNew("https://localhost/myactivity", obj)

	if delete.Object == nil {
		t.Errorf("Missing Object in Delete activity %#v", delete.Object)
	}
	if delete.Object != obj {
		t.Errorf("Delete.Object different than what we initialized %#v %#v", delete.Object, obj)
	}
}

func TestObjectPropertyExistsForFollow(t *testing.T) {
	obj := activitypub.MentionNew("gigel")
	follow := activitypub.FollowNew("https://localhost/myactivity", obj)

	if follow.Object == nil {
		t.Errorf("Missing Object in Follow activity %#v", follow.Object)
	}
	if follow.Object != obj {
		t.Errorf("Follow.Object different than what we initialized %#v %#v", follow.Object, obj)
	}
}

func TestObjectPropertyExistsForLike(t *testing.T) {
	obj := activitypub.MentionNew("gigel")
	add := activitypub.LikeNew("https://localhost/myactivity", obj)

	if add.Object == nil {
		t.Errorf("Missing Object in Like activity %#v", add.Object)
	}
	if add.Object != obj {
		t.Errorf("Like.Object different than what we initialized %#v %#v", add.Object, obj)
	}
}

func TestObjectPropertyExistsForUpdate(t *testing.T) {
	obj := activitypub.MentionNew("gigel")
	update := activitypub.UpdateNew("https://localhost/myactivity", obj)

	if update.Object == nil {
		t.Errorf("Missing Object in Update activity %#v", update.Object)
	}
	if update.Object != obj {
		t.Errorf("Update.Object different than what we initialized %#v %#v", update.Object, obj)
	}
}

func TestObjectPropertyExistsForUndo(t *testing.T) {
	obj := activitypub.MentionNew("gigel")
	undo := activitypub.UndoNew("https://localhost/myactivity", obj)

	if undo.Object == nil {
		t.Errorf("Missing Object in Undo activity %#v", undo.Object)
	}
	if undo.Object != obj {
		t.Errorf("Undo.Object different than what we initialized %#v %#v", undo.Object, obj)
	}
}

// S2S Server: Activities requiring the target property
// The distribution of the following activities require that they contain the target property: Add, Remove.
func TestTargetPropertyExistsForAdd(t *testing.T) {
	obj := activitypub.MentionNew("foo")
	target := activitypub.MentionNew("bar")
	add := activitypub.AddNew("https://localhost/myactivity", obj, target)

	if add.Target == nil {
		t.Errorf("Missing Target in Add activity %#v", add.Target)
	}
	if add.Target != target {
		t.Errorf("Add.Target different than what we initialized %#v %#v", add.Target, target)
	}
}

func TestTargetPropertyExistsForRemove(t *testing.T) {
	obj := activitypub.MentionNew("foo")
	target := activitypub.MentionNew("bar")
	remove := activitypub.RemoveNew("https://localhost/myactivity", obj, target)

	if remove.Target == nil {
		t.Errorf("Missing Target in Remove activity %#v", remove.Target)
	}
	if remove.Target != target {
		t.Errorf("Remove.Target different than what we initialized %#v %#v", remove.Target, target)
	}
}

// S2S Server: Deduplication of recipient list
// Attempt to submit for delivery an activity that addresses the same actor (ie an actor with the same id) twice.
// (For example, the same actor could appear on both the to and cc fields, or the actor could be explicitly addressed
// in to but could also be a member of the addressed followers collection of the sending actor.)
// The server should deduplicate the list of inboxes to deliver to before delivering.
func TestDeduplication(t *testing.T) {
	to := activitypub.PersonNew("bob")
	o := activitypub.ObjectNew("something", activitypub.ArticleType)
	cc := activitypub.PersonNew("alice")

	c := activitypub.CreateNew("create", o)
	c.To.Append(to)
	c.CC.Append(cc)
	c.BCC.Append(cc)

	activitypub.RecipientsDeduplication(&c.To, &c.Bto, &c.CC, &c.BCC)

	iter := func(list activitypub.ObjectsArr, recIds *[]activitypub.ObjectID) error {
		for _, rec := range list {
			for _, id := range *recIds {
				if rec.Object().ID == id {
					return fmt.Errorf("%T[%s] already stored in recipients list, Deduplication faild", rec, id)
				}
			}
			*recIds = append(*recIds, rec.Object().ID)
		}
		return nil
	}

	var err error
	recIds := make([]activitypub.ObjectID, 0)
	err = iter(c.To, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = iter(c.Bto, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = iter(c.CC, &recIds)
	if err != nil {
		t.Error(err)
	}
	err = iter(c.BCC, &recIds)
	if err != nil {
		t.Error(err)
	}
}
