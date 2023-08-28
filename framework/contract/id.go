package contract

const IDKey = "webgo:id"

type IDService interface {
	NewID() string
}
