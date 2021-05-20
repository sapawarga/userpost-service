package helper

const (
	DELETED          int64 = -1
	INACTIVED        int64 = 0
	ACTIVED          int64 = 10
	DELETED_STRING         = "deleted"
	INACTIVED_STRING       = "inactived"
	ACTIVED_STRING         = "actived"
	SELECT_QUERY           = "select"
	UPDATE_QUERY           = "update"
	TYPE_USERPOST          = "user_post"
	HTTP_GET               = "GET"
	HTTP_PUT               = "PUT"
	HTTP_POST              = "POST"
	HTTP_DELETE            = "DELETE"
	STATUS_OK              = "status_ok"
	STATUS_CREATED         = "status_created"
	STATUS_UPDATED         = "status_updated"
	STATUS_DELETED         = "status_deleted"
)

type keyType string

var ACTORKEY keyType = "Actor"
