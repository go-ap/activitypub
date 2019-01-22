package activitypub

import as "github.com/go-ap/activitystreams"

// Actor is the ActivityPub version of an Activity Streams vocabulary Actor
type Actor struct {
	as.Actor

	Inbox     InboxStream
	Outbox    OutboxStream
	Followers FollowersCollection
	Following FollowingCollection
}
