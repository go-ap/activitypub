package activitypub

import (
	as "github.com/go-ap/activitystreams"
)

func JSONGetItemByType(typ as.ActivityVocabularyType) (as.Item, error) {
	var ret as.Item
	var err error

	switch typ {
	case as.ObjectType:
		o := Object{}
		o.Type = typ
		ret = &o
	case as.ActivityType:
		ret = &Activity{}
		o := ret.(*Activity)
		o.Type = typ
	case as.IntransitiveActivityType:
		ret = &IntransitiveActivity{}
		o := ret.(*IntransitiveActivity)
		o.Type = typ
	case as.ActorType:
		ret = &Object{}
		o := ret.(*Object)
		o.Type = typ
	case as.ApplicationType:
		ret = &Application{}
		o := ret.(*Application)
		o.Type = typ
	case as.GroupType:
		ret = &Group{}
		o := ret.(*Group)
		o.Type = typ
	case as.OrganizationType:
		ret = &Organization{}
		o := ret.(*Organization)
		o.Type = typ
	case as.PersonType:
		ret = &Person{}
		o := ret.(*Person)
		o.Type = typ
	case as.ServiceType:
		ret = &Service{}
		o := ret.(*Service)
		o.Type = typ
	case as.AcceptType:
		ret = &Accept{}
		o := ret.(*Accept)
		o.Type = typ
	case as.AddType:
		ret = &Add{}
		o := ret.(*Add)
		o.Type = typ
	case as.AnnounceType:
		ret = &Announce{}
		o := ret.(*Announce)
		o.Type = typ
	case as.ArriveType:
		ret = &Arrive{}
		o := ret.(*Arrive)
		o.Type = typ
	case as.BlockType:
		ret = &Block{}
		o := ret.(*Block)
		o.Type = typ
	case as.CreateType:
		ret = &Create{}
		o := ret.(*Create)
		o.Type = typ
	case as.DeleteType:
		ret = &Delete{}
		o := ret.(*Delete)
		o.Type = typ
	case as.DislikeType:
		ret = &Dislike{}
		o := ret.(*Dislike)
		o.Type = typ
	case as.FlagType:
		ret = &Flag{}
		o := ret.(*Flag)
		o.Type = typ
	case as.FollowType:
		ret = &Follow{}
		o := ret.(*Follow)
		o.Type = typ
	case as.IgnoreType:
		ret = &Ignore{}
		o := ret.(*Ignore)
		o.Type = typ
	case as.InviteType:
		ret = &Invite{}
		o := ret.(*Invite)
		o.Type = typ
	case as.JoinType:
		ret = &Join{}
		o := ret.(*Join)
		o.Type = typ
	case as.LeaveType:
		ret = &Leave{}
		o := ret.(*Leave)
		o.Type = typ
	case as.LikeType:
		ret = &Like{}
		o := ret.(*Like)
		o.Type = typ
	case as.ListenType:
		ret = &Listen{}
		o := ret.(*Listen)
		o.Type = typ
	case as.MoveType:
		ret = &Move{}
		o := ret.(*Move)
		o.Type = typ
	case as.OfferType:
		ret = &Offer{}
		o := ret.(*Offer)
		o.Type = typ
	case as.QuestionType:
		ret = &Question{}
		o := ret.(*Question)
		o.Type = typ
	case as.RejectType:
		ret = &Reject{}
		o := ret.(*Reject)
		o.Type = typ
	case as.ReadType:
		ret = &Read{}
		o := ret.(*Read)
		o.Type = typ
	case as.RemoveType:
		ret = &Remove{}
		o := ret.(*Remove)
		o.Type = typ
	case as.TentativeRejectType:
		ret = &TentativeReject{}
		o := ret.(*TentativeReject)
		o.Type = typ
	case as.TentativeAcceptType:
		ret = &TentativeAccept{}
		o := ret.(*TentativeAccept)
		o.Type = typ
	case as.TravelType:
		ret = &Travel{}
		o := ret.(*Travel)
		o.Type = typ
	case as.UndoType:
		ret = &Undo{}
		o := ret.(*Undo)
		o.Type = typ
	case as.UpdateType:
		ret = &Update{}
		o := ret.(*Update)
		o.Type = typ
	case as.ViewType:
		ret = &View{}
		o := ret.(*View)
		o.Type = typ
	case "":
		// when no type is available use a plain Object
		ret = &Object{}
	default:
		return as.JSONGetItemByType(typ)
	}
	return ret, err
}
