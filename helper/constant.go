package helper

const (
	DELETED          int64 = -1
	INACTIVED        int64 = 0
	ACTIVED          int64 = 10
	DELETED_STRING         = "deleted"
	INACTIVED_STRING       = "inactived"
	ACTIVED_STRING         = "actived"
	STATUSCREATED          = "status_created"
	STATUSUPDATED          = "status_updated"
	STATUSDELETED          = "status_deleted"
	SELECT_QUERY           = "select"
	UPDATE_QUERY           = "update"
	TYPE_USERPOST          = "user_post"
)

type keyType string

var ACTORKEY keyType = "Actor"
