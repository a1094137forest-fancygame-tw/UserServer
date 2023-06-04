package constant

type Service string

const (
	SERVICE_USER  Service = "user"
	SERVICE_CACHE Service = "cache"
	SERVICE_LOBBY Service = "lobby"
)

// mongo
const (
	MONGO_COLLECTION = "player"
)
