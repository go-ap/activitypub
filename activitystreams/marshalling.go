package activitystreams

import (
	"fmt"
)

func getAPObjectByType(typ ActivityVocabularyType) (Item, error) {
	var ret Item
	var err error

	switch typ {
	case ObjectType:
		ret = ObjectNew(typ)
	case LinkType:
		ret = &Link{}
		o := ret.(*Link)
		o.Type = typ
	case ActivityType:
		ret = &Activity{}
		o := ret.(*Activity)
		o.Type = typ
	case IntransitiveActivityType:
		ret = &IntransitiveActivity{}
		o := ret.(*IntransitiveActivity)
		o.Type = typ
	case ActorType:
		ret = &Actor{}
		o := ret.(*Actor)
		o.Type = typ
	case CollectionType:
		ret = &Collection{}
		o := ret.(*Collection)
		o.Type = typ
	case OrderedCollectionType:
		ret = &Link{}
		o := ret.(*Link)
		o.Type = typ
	case ArticleType:
		ret = ObjectNew(typ)
	case AudioType:
		ret = ObjectNew(typ)
	case DocumentType:
		ret = ObjectNew(typ)
	case EventType:
		o := Object{}
		o.Type = typ
	case ImageType:
		ret = ObjectNew(typ)
		o := ret.(*Object)
		o.Type = typ
	case NoteType:
		ret = ObjectNew(typ)
	case PageType:
		ret = ObjectNew(typ)
	case PlaceType:
		ret = ObjectNew(typ)
	case ProfileType:
		ret = ObjectNew(typ)
	case RelationshipType:
		ret = ObjectNew(typ)
	case TombstoneType:
		ret = ObjectNew(typ)
	case VideoType:
		ret = ObjectNew(typ)
	case MentionType:
		ret = &Mention{}
		o := ret.(*Mention)
		o.Type = typ
	case ApplicationType:
		ret = &Application{}
		o := ret.(*Application)
		o.Type = typ
	case GroupType:
		ret = &Group{}
		o := ret.(*Group)
		o.Type = typ
	case OrganizationType:
		ret = &Organization{}
		o := ret.(*Organization)
		o.Type = typ
	case PersonType:
		ret = &Person{}
		o := ret.(*Person)
		o.Type = typ
	case ServiceType:
		ret = &Service{}
		o := ret.(*Service)
		o.Type = typ
	case AcceptType:
		ret = &Accept{}
		o := ret.(*Accept)
		o.Type = typ
	case AddType:
		ret = &Add{}
		o := ret.(*Add)
		o.Type = typ
	case AnnounceType:
		ret = &Announce{}
		o := ret.(*Announce)
		o.Type = typ
	case ArriveType:
		ret = &Arrive{}
		o := ret.(*Arrive)
		o.Type = typ
	case BlockType:
		ret = &Block{}
		o := ret.(*Block)
		o.Type = typ
	case CreateType:
		ret = &Create{}
		o := ret.(*Create)
		o.Type = typ
	case DeleteType:
		ret = &Delete{}
		o := ret.(*Delete)
		o.Type = typ
	case DislikeType:
		ret = &Dislike{}
		o := ret.(*Dislike)
		o.Type = typ
	case FlagType:
		ret = &Flag{}
		o := ret.(*Flag)
		o.Type = typ
	case FollowType:
		ret = &Follow{}
		o := ret.(*Follow)
		o.Type = typ
	case IgnoreType:
		ret = &Ignore{}
		o := ret.(*Ignore)
		o.Type = typ
	case InviteType:
		ret = &Invite{}
		o := ret.(*Invite)
		o.Type = typ
	case JoinType:
		ret = &Join{}
		o := ret.(*Join)
		o.Type = typ
	case LeaveType:
		ret = &Leave{}
		o := ret.(*Leave)
		o.Type = typ
	case LikeType:
		ret = &Like{}
		o := ret.(*Like)
		o.Type = typ
	case ListenType:
		ret = &Listen{}
		o := ret.(*Listen)
		o.Type = typ
	case MoveType:
		ret = &Move{}
		o := ret.(*Move)
		o.Type = typ
	case OfferType:
		ret = &Offer{}
		o := ret.(*Offer)
		o.Type = typ
	case QuestionType:
		ret = &Question{}
		o := ret.(*Question)
		o.Type = typ
	case RejectType:
		ret = &Reject{}
		o := ret.(*Reject)
		o.Type = typ
	case ReadType:
		ret = &Read{}
		o := ret.(*Read)
		o.Type = typ
	case RemoveType:
		ret = &Remove{}
		o := ret.(*Remove)
		o.Type = typ
	case TentativeRejectType:
		ret = &TentativeReject{}
		o := ret.(*TentativeReject)
		o.Type = typ
	case TentativeAcceptType:
		ret = &TentativeAccept{}
		o := ret.(*TentativeAccept)
		o.Type = typ
	case TravelType:
		ret = &Travel{}
		o := ret.(*Travel)
		o.Type = typ
	case UndoType:
		ret = &Undo{}
		o := ret.(*Undo)
		o.Type = typ
	case UpdateType:
		ret = &Update{}
		o := ret.(*Update)
		o.Type = typ
	case ViewType:
		ret = &View{}
		o := ret.(*View)
		o.Type = typ
	case "":
		// when no type is available use a plain Object
		ret = &Object{}
	default:
		ret = nil
		err = fmt.Errorf("unrecognized ActivityPub type %q", typ)
	}
	return ret, err
}
