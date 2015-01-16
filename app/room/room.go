package room

import "time"

type Event struct {
	Type      string
	User      string
	Timestamp int
	Text      string
}

func NewEvent(typ, user, msg string) Event {
	return Event{typ, user, int(time.Now().Unix()), msg}
}
