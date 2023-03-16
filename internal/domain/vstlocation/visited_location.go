package vstlocation

import "time"

type VisitedLocation struct {
	Id                           int64
	DateTimeOfVisitLocationPoint time.Time
	LocationPointId              int64
}
